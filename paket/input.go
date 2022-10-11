package paket

type GetIdPaket struct {
	ID int `uri:"id" binding:"required"`
}

type PaketInput struct {
	NamaPaket  string `json:"nama_paket" binding:"required"`
	Harga      string `json:"harga" binding:"required"`
	JenisPaket string `json:"jenis_paket" binding:"required"`
	WarnaLabel string `json:"warna_label" binding:"required"`
}

type Deskripsi_PaketInput struct {
	ID        int    `json:"id"`
	Id_Paket  int    `json:"id_paket"`
	Deskripsi string `json:"deskripsi"`
}
