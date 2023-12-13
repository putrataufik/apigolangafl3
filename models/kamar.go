package models

type Kamar struct {
	Id           int64  `gorm:"primaryKey" json:"id"`
	NomorKamar   string `gorm:"type:varchar(5)" json:"nomor_kamar"`
	PemilikKamar string `gorm:"type:varchar(40)" json:"pemilik_kamar"`
	StatusKamar  string `gorm:"type:varchar(15)" json:"status_kamar"`
	WaktuMasuk   string `gorm:"type:varchar(10)" json:"waktu_masuk"`  // Perubahan varcahar menjadi varchar di sini
	WaktuKeluar  string `gorm:"type:varchar(10)" json:"waktu_keluar"` // Perubahan varcahar menjadi varchar di sini
}

func (k *Kamar) TableName() string {
	return "kamars"
}

type KamarInfo struct {
	NomorKamar string `json:"nomor_kamar"`
	Status     string `json:"status"`
}

type PesanKamarRequest struct {
	NamaPemesan string `json:"namaPemesan"`
	WaktuMasuk  string `json:"waktuMasuk"`
	WaktuKeluar string `json:"waktuKeluar"`
	StatusKos   string `json:"statusKos"`
}
