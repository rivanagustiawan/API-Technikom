package handler

import (
	"net/http"
	"technikom/paket"

	"github.com/gin-gonic/gin"
)

type paketHandler struct {
	paketService paket.Service
}

func NewPaketHandler(paketService paket.Service) *paketHandler {
	return &paketHandler{paketService}
}

func (h *paketHandler) GetPakets(c *gin.Context) {

	pakets, err := h.paketService.GetPakets()

	if err != nil {

		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, pakets)
}
func (h *paketHandler) GetPaket(c *gin.Context) {
	var input paket.GetIdPaket

	err := c.ShouldBindUri(&input)
	if err != nil {

		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	paket, err := h.paketService.GetPaket(input)

	if err != nil {

		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, paket)
}

func (h *paketHandler) CreatePaket(c *gin.Context) {
	var input paket.PaketInput

	err := c.ShouldBindJSON(&input)
	if err != nil {

		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	newPaket, err := h.paketService.CreatePaket(input)
	if err != nil {

		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, newPaket)
}
func (h *paketHandler) CreateDeskPaket(c *gin.Context) {
	var input paket.Deskripsi_PaketInput

	err := c.ShouldBindJSON(&input)
	if err != nil {

		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	newDeskPaket, err := h.paketService.CreateDeskPaket(input)
	if err != nil {

		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, newDeskPaket)
}

func (h *paketHandler) UpdatePaket(c *gin.Context) {
	var inputID paket.GetIdPaket

	err := c.ShouldBindUri(&inputID)
	if err != nil {

		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var input paket.PaketInput

	err = c.ShouldBindJSON(&input)
	if err != nil {

		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	updatePaket, err := h.paketService.Update(inputID, input)
	if err != nil {

		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, updatePaket)
}

func (h *paketHandler) DeletePaket(c *gin.Context) {
	var input paket.GetIdPaket

	err := c.ShouldBindUri(&input)
	if err != nil {

		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	h.paketService.DeletePaket(input)

	c.JSON(http.StatusOK, "Paket Has Been Deleted !")
}
func (h *paketHandler) DeleteDesk(c *gin.Context) {
	var input paket.GetIdPaket

	err := c.ShouldBindUri(&input)
	if err != nil {

		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	h.paketService.DeleteDeskripsi(input)

	c.JSON(http.StatusOK, "Deskripsi Paket Has Been Deleted !")
}
