package klien

import "time"

type Klien struct {
	ID        uint      `json:"id"`
	NamaKlien string    `json:"namaKlien"`
	PicLogo   string    `json:"picLogo"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
