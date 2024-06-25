package domain

import "gorm.io/gorm"

type DefaultModel struct {
	gorm.Model `gorm:"embedded"`
	CreatedBy  string `gorm:"type:varbinary(40)"`
	UpdatedBy  string `gorm:"type:varbinary(40)"`
}
