package portofolio

import "mime/multipart"

type GetIdPorto struct {
	ID int `uri:"id" binding:"required"`
}

type PortofolioInput struct {
	Id_Klien      int                   `form:"id_klien" binding:"required"`
	NamaApps      string                `form:"namaApps" binding:"required"`
	JenisPaket    string                `form:"jenisPaket" binding:"required"`
	LinkApps      string                `form:"linkApps" binding:"required"`
	DeskripsiApps string                `form:"deskripsiApps" binding:"required"`
	PhotoApps1    *multipart.FileHeader `form:"photoApps1" binding:"required"`
	PhotoApps2    *multipart.FileHeader `form:"photoApps2" binding:"required" `
	PhotoApps3    *multipart.FileHeader `form:"photoApps3" binding:"required" `
	PhotoApps4    *multipart.FileHeader `form:"photoApps4" binding:"required" `
}
