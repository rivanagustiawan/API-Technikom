package klien

type GetIdKlien struct {
	ID int `uri:"id" binding:"required"`
}

type KlienInput struct {
	NamaKlien string `json:"namaKlien" binding:"required"`
	PicLogo   string `json:"picLogo" binding:"required"`
}
