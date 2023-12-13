package models

import (
	"gorm.io/gorm"
)

type User struct {
	Id                 int64  `gorm:"primaryKey" json:"id_user"`
	NamaUser           string `gorm:"type:varchar(255)" json:"nama_user"`
	IdKamar            int64  `json:"id_kamar"`
	IdDetailPembayaran int64  `json:"id_detail_pembayaran"`
}

func (u *User) TableName() string {
	return "users"
}

// Menambahkan relasi antara User dan Kamar menggunakan foreign key IdKamar
func (u *User) Kamar(db *gorm.DB) (*Kamar, error) {
	var kamar Kamar
	err := db.Model(&Kamar{}).Where("id = ?", u.IdKamar).Find(&kamar).Error
	if err != nil {
		return nil, err
	}
	return &kamar, nil
}

// Menambahkan relasi antara User dan DetailPembayaran menggunakan foreign key IdDetailPembayaran
func (u *User) DetailPembayaran(db *gorm.DB) (*DetailPembayaran, error) {
	var detail DetailPembayaran
	err := db.Model(&DetailPembayaran{}).Where("id_detail_pembayaran = ?", u.IdDetailPembayaran).Find(&detail).Error
	if err != nil {
		return nil, err
	}
	return &detail, nil
}
