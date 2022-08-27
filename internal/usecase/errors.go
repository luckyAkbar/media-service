package usecase

import (
	"errors"
	"fmt"
)

var (
	ErrServerFailedToSaveImage = errors.New("server failed to save the uploaded image")
	ErrMimeTypeNotFound        = errors.New("mimetype is not found")
	ErrMimeTypeIsForbidden     = errors.New("mimetype is forbidden")
	ErrFileNotFound            = errors.New("image not found")
	ErrFileSizeToLarge         = errors.New("file too large")
	ErrUpdateKeyMismatch       = errors.New("update key mismatch")
	ErrDeleteKeyMismatch       = errors.New("delete key mismatch")
	ErrServerFailedToDelete    = errors.New("server failed to delete that image")
	ErrInternal                = errors.New("internal error")
	ErrNotFound                = errors.New("not found")
)

func ErrMimeTypeForbidden(mimetype string) error {
	return fmt.Errorf("mimetype: %s is forbidden", mimetype)
}

func ErrFileToLarge(actual, max int64) error {
	return fmt.Errorf("file too large with size: %d, max only: %d", actual, max)
}
