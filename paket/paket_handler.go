package paket

import (
	"gocatering/helper"
	"gocatering/model"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/labstack/echo"
)

type PaketHandler struct {
	PaketService Service
}

func NewPaketHandler(s Service) *PaketHandler {
	return &PaketHandler{PaketService: s}
}

func (h *PaketHandler) CreatePaket(e echo.Context) error {
	var paket model.Paket
	if err := e.Bind(&paket); err != nil {
		return e.JSON(http.StatusUnprocessableEntity,
			helper.Apiresponse(err.Error(), http.StatusUnprocessableEntity, "failed", nil))
	}

	err := h.PaketService.CreatePaket(&paket)
	if err != nil {
		return e.JSON(http.StatusInternalServerError,
			helper.Apiresponse(err.Error(), http.StatusInternalServerError, "failed", nil))
	}
	return e.JSON(http.StatusCreated,
		helper.Apiresponse("Paket Created", http.StatusOK, "success", paket))

}

func (h *PaketHandler) GetPaketByID(e echo.Context) error {
	id, _ := strconv.Atoi(e.Param("id"))
	res, err := h.PaketService.GetPaketByID(id)

	if err != nil {
		return e.JSON(http.StatusInternalServerError,
			helper.Apiresponse(err.Error(), http.StatusInternalServerError, "failed", nil))
	}
	if res == nil {
		return e.JSON(http.StatusNotFound,
			helper.Apiresponse(err.Error(), http.StatusNotFound, "Paket Not Found", nil))
	}

	return e.JSON(http.StatusOK,
		helper.Apiresponse("Paket Found", http.StatusOK, "success", res))
}

func (h *PaketHandler) GetAllPaket(e echo.Context) error {

	res, err := h.PaketService.GetAllPaket()

	if err != nil {
		return e.JSON(http.StatusInternalServerError,
			helper.Apiresponse(err.Error(), http.StatusInternalServerError, "failed", nil))
	}

	return e.JSON(http.StatusOK,
		helper.Apiresponse("List of Paket", http.StatusOK, "success", res))

}

func (h *PaketHandler) UpdatePaket(e echo.Context) error {
	var paket model.Paket
	id, _ := strconv.Atoi(e.Param("id"))

	if err := e.Bind(&paket); err != nil {
		return e.JSON(http.StatusUnprocessableEntity,
			helper.Apiresponse(err.Error(), http.StatusUnprocessableEntity, "failed", nil))
	}

	errors := h.PaketService.UpdatePaket(id, &paket)

	if errors != nil {
		return e.JSON(http.StatusInternalServerError,
			helper.Apiresponse(errors.Error(), http.StatusInternalServerError, "failed", nil))
	}

	return e.JSON(http.StatusCreated,
		helper.Apiresponse("Paket Updated", http.StatusOK, "success", paket))
}

func (h *PaketHandler) DeletePaket(e echo.Context) error {
	id, _ := strconv.Atoi(e.Param("id"))
	deletedReg, errs := h.PaketService.DeletePaket(id)
	if errs != nil {
		return e.JSON(http.StatusInternalServerError,
			helper.Apiresponse(errs.Error(), http.StatusInternalServerError, "failed", nil))
	}

	return e.JSON(http.StatusOK,
		helper.Apiresponse("Paket Deleted", http.StatusOK, "success", deletedReg))
}

func (h *PaketHandler) CreateImage(e echo.Context) error {
	var image model.Image

	if err := e.Bind(&image); err != nil {
		return e.JSON(http.StatusUnprocessableEntity,
			helper.Apiresponse(err.Error(), http.StatusUnprocessableEntity, "failed", nil))
	}

	var mainImage bool
	if image.IsMain == true {
		mainImage = true
	} else {
		mainImage = false
	}

	findPaket, err := h.PaketService.GetPaketByID(image.PaketID)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "error sql",
			"error":   err.Error(),
		})
	}
	//paketName := paket.Name
	paketName := strings.Replace(findPaket.Name, " ", "_", -1)
	file, err := e.FormFile("file")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	filePath := "./uploads/" + paketName + "_" + file.Filename
	fileSrc := "uploads/" + paketName + "_" + file.Filename

	dst, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return err
	}

	newImage := model.Image{
		PaketID:  findPaket.ID,
		FileName: fileSrc,
		IsMain:   mainImage,
	}

	errors := h.PaketService.CreateImage(&newImage)

	if errors != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "ooopss.... there is something error",
			"status":  err.Error,
		})
	}

	return e.JSON(http.StatusCreated,
		helper.Apiresponse("Image Uploaded", http.StatusOK, "success", newImage))

}

func (h *PaketHandler) DeleteImage(e echo.Context) error {
	id, _ := strconv.Atoi(e.Param("id"))
	deletedImg, errs := h.PaketService.DeleteImage(id)
	if errs != nil {
		return e.JSON(http.StatusInternalServerError,
			helper.Apiresponse(errs.Error(), http.StatusInternalServerError, "failed", nil))
	}

	return e.JSON(http.StatusOK,
		helper.Apiresponse("Image Deleted", http.StatusOK, "success", deletedImg))
}
