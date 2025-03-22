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

type Bic struct {
	Id           int `gorm:"primaryKey"`
	CountryId    int `gorm:"not null"`
	Country      Country
	Bic          string `gorm:"size:11;not null;unique"`
	CodeType     string `gorm:"size:12;not null"`
	BankId       int    `gorm:"not null"`
	Bank         Bank
	Address      string `gorm:"size:255"`
	Town         string `gorm:"size:64;not null"`
	TimeZone     string `gorm:"size:32;not null"`
	LocationCode string `gorm:"size:2;not null"`
	Branch       string `gorm:"size:3;default:'XXX'"`
}
