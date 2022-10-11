package portofolio

type GetIdPorto struct {
	ID int `uri:"id" binding:"required"`
}

type PortofolioInput struct {
	Id_Klien      int    `json:"id_klien" binding:"required"`
	NamaApps      string `json:"namaApps" binding:"required"`
	JenisPaket    string `json:"jenisPaket" binding:"required"`
	LinkApps      string `json:"linkApps" binding:"required"`
	DeskripsiApps string `json:"deskripsiApps" binding:"required"`
	PhotoApps1    string `json:"photoApps1" binding:"required"`
	PhotoApps2    string `json:"photoApps2" binding:"required"`
	PhotoApps3    string `json:"photoApps3" binding:"required"`
	PhotoApps4    string `json:"photoApps4" binding:"required"`
}
