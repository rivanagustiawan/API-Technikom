package handler

import (
	"net/http"
	"technikom/klien"

	"github.com/gin-gonic/gin"
)

type klienHandler struct {
	klienService klien.Service
}

func NewKlienHandler(klienService klien.Service) *klienHandler {
	return &klienHandler{klienService}
}

func (h *klienHandler) GetKliens(c *gin.Context) {

	kliens, err := h.klienService.GetKliens()

	if err != nil {

		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, kliens)
}
func (h *klienHandler) GetKlien(c *gin.Context) {

	var inputID klien.GetIdKlien

	err := c.ShouldBindUri(&inputID)
	if err != nil {

		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	kliens, err := h.klienService.GetKlien(inputID)

	if err != nil {

		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, kliens)
}

func (h *klienHandler) CreateKlien(c *gin.Context) {

	var input klien.KlienInput

	err := c.ShouldBindJSON(&input)
	if err != nil {

		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	newKlien, err := h.klienService.CreateKlien(input)
	if err != nil {

		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, newKlien)
}

func (h *klienHandler) UpdateKlien(c *gin.Context) {
	var inputID klien.GetIdKlien

	err := c.ShouldBindUri(&inputID)
	if err != nil {

		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var input klien.KlienInput

	err = c.ShouldBindJSON(&input)
	if err != nil {

		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	updateKlien, err := h.klienService.UpdateKlien(inputID, input)
	if err != nil {

		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, updateKlien)
}

func (h *klienHandler) DeleteKlien(c *gin.Context) {
	var input klien.GetIdKlien

	err := c.ShouldBindUri(&input)
	if err != nil {

		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	h.klienService.DeleteKlien(input)

	c.JSON(http.StatusOK, "Klien Has Been Deleted !")
}
