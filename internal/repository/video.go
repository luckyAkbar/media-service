package repository

import (
	"context"
	"media-service/internal/model"

	"github.com/kumparan/go-utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type videoRepository struct {
	db *gorm.DB
}

func NewVideoRepository(db *gorm.DB) model.VideoRepository {
	return &videoRepository{
		db: db,
	}
}

func (r *videoRepository) Store(ctx context.Context, v *model.Video) error {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":   utils.DumpIncomingContext(ctx),
		"video": utils.Dump(v),
	})

	if err := r.db.Create(v).Error; err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func (r *videoRepository) GetByID(ctx context.Context, id int64) (*model.Video, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx": utils.DumpIncomingContext(ctx),
		"id":  id,
	})

	video := &model.Video{}
	err := r.db.WithContext(ctx).Model(&model.Video{}).Where("id=?", id).Take(video).Error
	switch err {
	default:
		logger.Error(err)
		return nil, err
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	case nil:
		return video, nil
	}
}
