package model

type Category struct {
	ID          string `gorm:"type:uuid;primaryKey;not null;unique" json:"id" binding:"required"`
	Name        string `gorm:"type:varchar(255);not null;unique" json:"name" binding:"required,alphanum"`
	Description string `gorm:"type:text" json:"description" binding:"omitempty"`
}
