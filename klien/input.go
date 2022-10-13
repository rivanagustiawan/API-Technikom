package klien

import "mime/multipart"

type GetIdKlien struct {
	ID int `uri:"id" binding:"required"`
}

type KlienInput struct {
	NamaKlien string                `form:"namaKlien" binding:"required"`
	PicLogo   *multipart.FileHeader `form:"picLogo" binding:"required"`
}
type KlienInputUpdate struct {
	NamaKlien string                `form:"namaKlien"`
	PicLogo   *multipart.FileHeader `form:"picLogo"`
}
