package handler

import (
	"net/http"
	"technikom/portofolio"

	"github.com/gin-gonic/gin"
)

type portofolioHandler struct {
	portoService portofolio.Service
}

func NewPortoHandler(portoService portofolio.Service) *portofolioHandler {
	return &portofolioHandler{portoService}
}

func (h *portofolioHandler) GetPortofolios(c *gin.Context) {

	porto, err := h.portoService.GetPortofolios()

	if err != nil {

		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, porto)
}
func (h *portofolioHandler) GetPortofolio(c *gin.Context) {
	var inputID portofolio.GetIdPorto

	err := c.ShouldBindUri(&inputID)
	if err != nil {

		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	porto, err := h.portoService.GetPortofolio(inputID)

	if err != nil {

		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, porto)
}

func (h *portofolioHandler) CreatePorto(c *gin.Context) {

	var input portofolio.PortofolioInput

	err := c.ShouldBindJSON(&input)
	if err != nil {

		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	newPorto, err := h.portoService.CreatePortofolio(input)
	if err != nil {

		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, newPorto)
}

func (h *portofolioHandler) UpdatePortofolio(c *gin.Context) {
	var inputID portofolio.GetIdPorto

	err := c.ShouldBindUri(&inputID)
	if err != nil {

		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var input portofolio.PortofolioInput

	err = c.ShouldBindJSON(&input)
	if err != nil {

		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	updatePorto, err := h.portoService.UpdatePortofolio(inputID, input)
	if err != nil {

		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, updatePorto)
}

func (h *portofolioHandler) DetelePortofolio(c *gin.Context) {
	var input portofolio.GetIdPorto

	err := c.ShouldBindUri(&input)
	if err != nil {

		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	h.portoService.DeletePortofolio(input)

	c.JSON(http.StatusOK, "Portofolio Has Been Deleted !")
}
