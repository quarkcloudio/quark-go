package model

// 字段
type AttachmentCategory struct {
	Id          int    `json:"id" gorm:"autoIncrement"`
	Source      string `json:"source" gorm:"size:100"`
	Uid         int    `json:"uid" gorm:"size:11;default:0"`
	Title       string `json:"title" gorm:"size:255;not null"`
	Sort        int    `json:"sort" gorm:"size:11;default:0"`
	Description string `json:"description" gorm:"size:255"`
}
