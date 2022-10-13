package handler

import (
	"fmt"
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

	file, err := c.FormFile("photoApps1")
	if err != nil {

		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	file2, err := c.FormFile("photoApps2")
	if err != nil {

		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	file3, err := c.FormFile("photoApps3")
	if err != nil {

		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	file4, err := c.FormFile("photoApps4")
	if err != nil {

		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	path := fmt.Sprintf("images/portofolio/%s", file.Filename)
	path2 := fmt.Sprintf("images/portofolio/%s", file2.Filename)
	path3 := fmt.Sprintf("images/portofolio/%s", file3.Filename)
	path4 := fmt.Sprintf("images/portofolio/%s", file4.Filename)

	c.SaveUploadedFile(file, path)
	c.SaveUploadedFile(file2, path2)
	c.SaveUploadedFile(file3, path3)
	c.SaveUploadedFile(file4, path4)

	var input portofolio.PortofolioInput

	err = c.ShouldBind(&input)
	if err != nil {

		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	newPorto, err := h.portoService.CreatePortofolio(input, path, path2, path3, path4)
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
