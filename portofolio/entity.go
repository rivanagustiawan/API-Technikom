package portofolio

import (
	"time"
)

type Portofolio struct {
	ID            uint      `json:"id"`
	Id_Klien      int       `json:"id_klien"`
	NamaApps      string    `json:"namaApps"`
	JenisPaket    string    `json:"jenisPaket"`
	LinkApps      string    `json:"linkApps"`
	DeskripsiApps string    `json:"deskripsiApps"`
	PhotoApps1    string    `json:"photoApps1"`
	PhotoApps2    string    `json:"photoApps2"`
	PhotoApps3    string    `json:"photoApps3"`
	PhotoApps4    string    `json:"photoApps4"`
	Klien         Klien     `gorm:"Foreignkey:Id_Klien;association_foreignkey:ID;" json:"Klien"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Klien struct {
	ID        int       `json:"id"`
	NamaKlien string    `json:"namaKlien"`
	PicLogo   string    `json:"picLogo"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
