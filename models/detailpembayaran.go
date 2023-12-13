package models

type DetailPembayaran struct {
	IdDetailPembayaran int64  `gorm:"primaryKey" json:"id"`
	IdUser             int64  `gorm:"type:int" json:"id_user"`
	WaktuMasuk         string `gorm:"type:varchar(10)" json:"waktu_masuk"` // Perubahan varcahar menjadi varchar di sini
	WaktuKeluar        string `gorm:"type:varchar(10)" json:"waktu_keluar"`
	WaktuBayar         string `gorm:"type:varchar(10)" json:"waktu_bayar"`
	TotalPembayaran    int64  `gorm:"type:int" json:"total_pembayaran"`
}

func (d *DetailPembayaran) TableName() string {
	return "detail_pembayarans"
}
