package model

import (
	"context"
	"mime/multipart"
	"time"
)

type Video struct {
	ID            int64     `gorm:"primaryKey" json:"id"`
	Filename      string    `gorm:"not null" json:"-"`
	Title         string    `gorm:"default: " json:"title"`
	Extention     string    `gorm:"not null" json:"extention"`
	MimeType      string    `gorm:"not null" json:"mime_type"`
	SizeBytes     int64     `gorm:"not null,column:size_bytes" json:"size_bytes"`
	LengthSeconds int64     `gorm:"not null" json:"length_seconds"`
	CreatedAt     time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt     time.Time `gorm:"not null" json:"updated_at"`
}

type VideoUsecase interface {
	Upload(ctx context.Context, file *multipart.FileHeader, title string) (*Video, error)
	Download(ctx context.Context, id int64) (*Video, error)
}

type VideoRepository interface {
	Store(ctx context.Context, video *Video) error
	GetByID(ctx context.Context, id int64) (*Video, error)
}
