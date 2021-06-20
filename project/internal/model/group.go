package model

type Group struct {
	Name string  `gorm:"type:varchar(191);unique;not null;default:''"`
}
