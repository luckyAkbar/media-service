package repository

import (
	"image-service/internal/db"
	"image-service/internal/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ImageRepo struct {
	db        *gorm.DB
	Name      string
	ImageKey  string
	DeleteKey string
	UpdateKey string
	Size      int64
}

func NewImageRepo() *ImageRepo {
	return &ImageRepo{
		db: db.DB,
	}
}

func (r *ImageRepo) Save(name, imageKey, deleteKey, updateKey string, size int64) error {
	image := &model.Image{
		Name:           name,
		ImageKey:       imageKey,
		DeleteKey:      deleteKey,
		UpdateKey:      updateKey,
		ImageSizeBytes: size,
	}

	if err := r.db.Save(image).Error; err != nil {
		return err
	}

	return nil
}

func (r *ImageRepo) Find(imageKey string) (model.Image, error) {
	image := model.Image{}
	err := r.db.Model(&image).Where("image_key = ?", imageKey).First(&image).Error
	if err != nil {
		logrus.Error(err)
		return image, err
	}

	return image, nil
}
