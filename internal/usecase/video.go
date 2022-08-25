package usecase

import (
	"context"
	"fmt"
	"io"
	"media-service/internal/config"
	"media-service/internal/helper"
	"media-service/internal/model"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"

	mime "github.com/cubewise-code/go-mime"
	"github.com/kumparan/go-utils"
	"github.com/sirupsen/logrus"
)

type videoUsecase struct {
	videoRepo         model.VideoRepository
	allowedMimeTypes  []string
	allowedExtentions []string
	maxSizeBytes      int64
}

func NewVideoUsecase(videoRepo model.VideoRepository) model.VideoUsecase {
	return &videoUsecase{
		videoRepo:         videoRepo,
		allowedMimeTypes:  config.AllowedVideoMimeType(),
		allowedExtentions: config.AllowedVideoExt(),
		maxSizeBytes:      config.MaxVideoSizeBytes(),
	}
}

func (u *videoUsecase) Upload(ctx context.Context, file *multipart.FileHeader, title string) (*model.Video, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":   utils.DumpIncomingContext(ctx),
		"file":  utils.Dump(file),
		"title": title,
	})

	mimetype, err := u.filterVideoMimeType(file)
	if err != nil {
		logger.Info(err)
		return nil, ErrMimeTypeIsForbidden
	}

	ext, err := u.filterVideoExtentionAgainstMimeType(file, mimetype)
	if err != nil {
		logger.Info(err)
		return nil, ErrMimeTypeIsForbidden
	}

	if err := u.filterVideoSize(file); err != nil {
		logger.Info(err)
		return nil, ErrFileSizeToLarge
	}

	now := time.Now()
	video := &model.Video{
		ID:        utils.GenerateID(),
		Title:     title,
		Filename:  fmt.Sprintf("%s.%s", helper.GenerateRandString(config.VideoNameLength(), time.Now().Unix()), ext),
		Extention: ext,
		SizeBytes: file.Size,

		// TODO check video length
		LengthSeconds: 0,
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	if err := u.saveVideo(file, video.Filename); err != nil {
		logger.Error(err)
		return nil, ErrInternal
	}

	if err := u.videoRepo.Store(ctx, video); err != nil {
		logger.Error(err)
		return nil, ErrInternal
	}

	return video, nil
}

func (u *videoUsecase) filterVideoMimeType(file *multipart.FileHeader) (string, error) {
	mimetype := file.Header["Content-Type"]

	if len(mimetype) == 0 {
		return "", ErrMimeTypeNotFound
	}

	for _, actual := range mimetype {
		for _, allowed := range u.allowedMimeTypes {
			if actual == allowed {
				return actual, nil
			}
		}
	}

	return "", ErrMimeTypeIsForbidden
}

func (u *videoUsecase) filterVideoSize(file *multipart.FileHeader) error {
	if file.Size > u.maxSizeBytes {
		return ErrFileSizeToLarge
	}

	return nil
}

func (u *videoUsecase) filterVideoExtentionAgainstMimeType(file *multipart.FileHeader, actualMimetype string) (string, error) {
	filename := file.Filename
	t := strings.Split(filename, ".")
	ext := t[len(t)-1]

	contentType := mime.TypeByExtension(ext)

	if contentType != actualMimetype {
		return "", ErrMimeTypeIsForbidden
	}

	return ext, nil
}

func (u *videoUsecase) saveVideo(file *multipart.FileHeader, filename string) error {
	video, err := file.Open()
	if err != nil {
		return err
	}

	defer video.Close()

	dst, err := os.Create(path.Join(config.VideoStoragePath(), filename))
	if err != nil {
		logrus.Error(err)
		return err
	}

	defer dst.Close()

	if _, err := io.Copy(dst, video); err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}
