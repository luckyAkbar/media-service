package usecase

import (
	cryptoRand "crypto/rand"
	"fmt"
	"io"
	"math/rand"
	"media-service/internal/config"
	"media-service/internal/helper"
	"media-service/internal/model"
	"media-service/internal/repository"
	"mime"
	"mime/multipart"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type ImageHandler struct {
	AllowedImageExt []string
	AllowedMimeType []string
	File            *multipart.FileHeader
	MaxSizeBytes    int64
	ImageFullName   string
	MimeType        string
	Keys            imageKey
}

type imageKey struct {
	ImageKey  string `json:"image_key"`
	UpdateKey string `json:"update_key"`
	DeleteKey string `json:"delete_key"`
}

func NewImageHandler() *ImageHandler {
	return &ImageHandler{
		AllowedImageExt: config.AllowedImageExt(),
		AllowedMimeType: config.AllowedImageMimeType(),
		MaxSizeBytes:    config.MaxImageSizeBytes(),
	}
}

func (i *ImageHandler) HandleUpload(file *multipart.FileHeader) (imageKey, error) {
	i.File = file
	if err := i.filterMimeType(); err != nil {
		logrus.Error(err)
		return imageKey{}, err
	}

	if err := i.filterSize(); err != nil {
		logrus.Error(err)
		return imageKey{}, err
	}

	i.Keys = i.generateKeys()
	i.ImageFullName = i.generateImageName()

	if err := i.saveImage(); err != nil {
		logrus.Error(err)
		return imageKey{}, err
	}

	imageRepo := repository.NewImageRepo()
	err := imageRepo.Save(
		i.ImageFullName,
		i.Keys.ImageKey,
		i.Keys.DeleteKey,
		i.Keys.UpdateKey,
		i.File.Size,
	)
	if err != nil {
		logrus.Error(err)
		return imageKey{}, ErrServerFailedToSaveImage
	}

	return i.Keys, nil
}

func (i *ImageHandler) HandleGet(imageKey string) (model.Image, error) {
	imageRepo := repository.NewImageRepo()
	data, err := imageRepo.Find(imageKey)

	if err != nil {
		logrus.Error(err)
		return data, err
	}

	return data, nil
}

func (i *ImageHandler) HandleUpdate(file *multipart.FileHeader, key, updateKey string) (imageKey, error) {
	i.File = file

	if err := i.applyCommonFileFilter(); err != nil {
		return imageKey{}, err
	}

	imageRepo := repository.NewImageRepo()
	imageData, err := imageRepo.Find(key)
	if err != nil {
		return imageKey{}, ErrFileNotFound
	}

	if updateKey != imageData.UpdateKey {
		logrus.Error(ErrUpdateKeyMismatch)
		return imageKey{}, ErrUpdateKeyMismatch
	}

	i.ImageFullName = i.generateImageName()
	imageData.Name = i.ImageFullName
	imageData.UpdateKey = i.generateNewUpdateKey()

	if err := i.saveImage(); err != nil {
		logrus.Error(err)
		return imageKey{}, ErrServerFailedToSaveImage
	}

	if err := imageRepo.Update(imageData); err != nil {
		logrus.Error(err)
		return imageKey{}, ErrServerFailedToSaveImage
	}

	return imageKey{
		ImageKey:  imageData.ImageKey,
		UpdateKey: imageData.UpdateKey,
		DeleteKey: imageData.DeleteKey,
	}, nil
}

func (i *ImageHandler) HandleDelete(imageKey, deleteKey string) error {
	imageRepo := repository.NewImageRepo()
	imageData, err := imageRepo.Find(imageKey)
	if err != nil {
		logrus.Error(err)
		return ErrFileNotFound
	}

	if imageData.DeleteKey != deleteKey {
		logrus.Error(ErrDeleteKeyMismatch)
		return ErrDeleteKeyMismatch
	}

	if err := imageRepo.Delete(imageData); err != nil {
		logrus.Error(err)
		return ErrServerFailedToDelete
	}

	return nil
}

func (i *ImageHandler) applyCommonFileFilter() error {
	if err := i.filterMimeType(); err != nil {
		logrus.Error(err)
		return err
	}

	if err := i.filterSize(); err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func (i *ImageHandler) saveImage() error {
	src, err := i.File.Open()
	if err != nil {
		logrus.Error(err)
		return ErrServerFailedToSaveImage
	}
	defer src.Close()

	dst, err := os.Create(fmt.Sprintf("%s%s", config.ImageStoragePath(), i.ImageFullName))
	if err != nil {
		logrus.Error(err)
		return ErrServerFailedToSaveImage
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		logrus.Error(err)
		return ErrServerFailedToSaveImage
	}

	return nil
}

func (i *ImageHandler) filterMimeType() error {
	mimetype := i.File.Header["Content-Type"][0]
	if mimetype == "" {
		return ErrMimeTypeNotFound
	}

	for _, allowedMimetype := range i.AllowedMimeType {
		if allowedMimetype == mimetype {
			i.MimeType = mimetype
			return nil
		}
	}

	return ErrMimeTypeForbidden(mimetype)
}

func (i *ImageHandler) filterSize() error {
	if i.File.Size > i.MaxSizeBytes {
		return ErrFileToLarge(i.File.Size, i.MaxSizeBytes)
	}

	return nil
}

func (i *ImageHandler) generateImageName() string {
	now := time.Now().Unix()
	imageName := helper.GenerateRandString(config.ImageNameLength(), now)

	ext, err := mime.ExtensionsByType(i.MimeType)
	if err != nil {
		logrus.Error(fmt.Sprintf("Failed to get type from mimetype: %s, but it pass the mimetype check", i.MimeType))

		ext = []string{".jpg"}
	}

	return fmt.Sprintf("%s%d%s", imageName, now, ext[0])
}

func (i *ImageHandler) generateKeys() imageKey {
	now := time.Now().Unix()
	rand.Seed(now)
	int64ImageKeySeeder := now
	int64UpdateKeySeeder := rand.Int63n(now)
	int64DeleteKeySeeder, _ := cryptoRand.Prime(cryptoRand.Reader, 64)

	return imageKey{
		ImageKey:  helper.GenerateRandString(config.IMAGE_KEY_LENGTH, int64ImageKeySeeder),
		UpdateKey: helper.GenerateRandString(config.UPDATE_KEY_LENGTH, int64UpdateKeySeeder),
		DeleteKey: helper.GenerateRandString(config.DELETE_KEY_LENGTH, int64DeleteKeySeeder.Int64()),
	}
}

func (i *ImageHandler) generateNewUpdateKey() string {
	now := time.Now().Unix()
	int64UpdateKeySeeder := rand.Int63n(now)

	return helper.GenerateRandString(config.UPDATE_KEY_LENGTH, int64UpdateKeySeeder)
}
