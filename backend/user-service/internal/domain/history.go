package domain

import "time"


type UploadHistory struct {
	Filename string    `json:"filename" bson:"filename"`
	Size     int64     `json:"size" bson:"size"`
	Uploaded time.Time `json:"uploaded" bson:"uploaded"`
}


type UploadRepository interface {
	SaveUploadHistory(history *UploadHistory) error
	GetUploadHistory() ([]UploadHistory, error)
}
