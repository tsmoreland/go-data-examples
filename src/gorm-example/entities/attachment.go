package entities

import "github.com/jinzhu/gorm"

type Attachment struct {
	gorm.Model
	Data          []byte
	AppointmentID uint
}

func CreateAttachmentTable(db *gorm.DB) {
	db.
		CreateTable(&Attachment{})
}

func (e *Attachment) TableName() string {
	return "attachments"
}
