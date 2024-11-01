package dto

import "time"

type FileInfo struct {
	ID       int       `db:"id"`
	FilePath string    `db:"file_path"`
	UploadedAt time.Time `db:"uploaded_at"`
}