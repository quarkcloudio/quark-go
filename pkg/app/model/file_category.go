package model

// 字段
type FileCategory struct {
	Id          int    `json:"id" gorm:"autoIncrement"`
	ObjType     string `json:"obj_type" gorm:"size:100"`
	ObjId       int    `json:"obj_id" gorm:"size:11;default:0"`
	Title       string `json:"title" gorm:"size:255;not null"`
	Sort        int    `json:"sort" gorm:"size:11;default:0"`
	Description string `json:"description" gorm:"size:255"`
}
