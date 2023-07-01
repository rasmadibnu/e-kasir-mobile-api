package entity

import (
	"time"
)

type Transaksi struct {
	ID              int               `gorm:"column:id;type:int(11);primary_key" json:"id"`
	Diskon          int               `gorm:"column:diskon;type:int(11)" json:"diskon"`
	Ppn             int               `gorm:"column:ppn;type:int(11)" json:"ppn"`
	TotalBelanja    int               `gorm:"column:total_belanja;type:int(11)" json:"total_belanja"`
	KasirID         int               `gorm:"column:kasir_id;type:int(11)" json:"kasir_id"`
	DetailTransaksi []DetailTransaksi `json:"detail"`
	Tanggal         time.Time         `gorm:"column:tanggal;type:timestamp" json:"tanggal"`
}

func (m *Transaksi) TableName() string {
	return "transaksi"
}

type DetailTransaksi struct {
	ID          int `gorm:"column:id;type:int(11);primary_key" json:"id"`
	TransaksiID int `gorm:"column:transaksi_id;type:int(11)" json:"transaksi_id"`
	ProdukID    int `gorm:"column:produk_id;type:int(11)" json:"produk_id"`
	JumlahBeli  int `gorm:"column:jumlah_beli;type:int(11)" json:"jumlah_beli"`
}

func (m *DetailTransaksi) TableName() string {
	return "detail_transaksi"
}