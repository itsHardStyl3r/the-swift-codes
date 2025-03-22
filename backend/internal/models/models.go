package models

type Country struct {
	Id   int    `gorm:"primaryKey"`
	Name string `gorm:"size:60;not null"`
	Iso2 string `gorm:"size:2;unique;not null"`
}

type Bank struct {
	Id       int    `gorm:"primaryKey"`
	Name     string `gorm:"size:255;not null"`
	BankCode string `gorm:"size:4;not null"`
}
