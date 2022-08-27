package model

import (
	"mime/multipart"

	"github.com/jinzhu/gorm"
)

type Image struct {
	gorm.Model

	ImageKey          string `gorm:"unique"`
	DeleteKey         string
	UpdateKey         string
	Name              string
	UploaderIP        string
	UploaderUserAgent string
	ImageSizeBytes    int64
}

// TODO use model to fix image repository and usecase

type ImageKey struct {
	ImageKey  string `json:"image_key"`
	UpdateKey string `json:"update_key"`
	DeleteKey string `json:"delete_key"`
}

type ImageUsecase interface {
	HandleUpload(file *multipart.FileHeader) (ImageKey, error)
	HandleGet(imageKey string) (Image, error)
	HandleUpdate(file *multipart.FileHeader, imageKey, updateKey string) (ImageKey, error)
	HandleDelete(imageKey, deleteKey string) error
}

type ImageRepository interface {
	Save()
}
