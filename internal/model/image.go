package model

import (
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
