package paket

import "time"

type Paket struct {
	ID         uint              `json:"id" gorm:"primary_key"`
	NamaPaket  string            `json:"namaPaket"`
	Harga      string            `json:"Harga"`
	JenisPaket string            `json:"jenisPaket"`
	WarnaLabel string            `json:"warnaLabel"`
	Deskripsi  []Deskripsi_Paket `gorm:"Foreignkey:Id_Paket;association_foreignkey:ID;" json:"deskripsi"`
	CreatedAt  time.Time         `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time         `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Deskripsi_Paket struct {
	ID        int       `json:"id"`
	Id_Paket  int       `json:"id_paket"`
	Deskripsi string    `json:"deskripsi"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
