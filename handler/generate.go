package handler

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/Volgar04/dreampicai/db"
	"github.com/Volgar04/dreampicai/types"
	"github.com/Volgar04/dreampicai/view/generate"
	"github.com/go-chi/chi/v5"
)

func HandleGenerateIndex(w http.ResponseWriter, r *http.Request) error {
	user := getAuthenticatedUser(r)
	images, err := db.GetImagesByUserID(user.ID)
	if err != nil {
		return err
	}
	data := generate.ViewData{
		Images: images,
	}
	return render(r, w, generate.Index(data))
}

func HandleGenerateCreate(w http.ResponseWriter, r *http.Request) error {
	user := getAuthenticatedUser(r)
	prompt := "red sport car in a garden"
	image := types.Image{
		Prompt: prompt,
		Status: types.ImageStatusPending,
		UserID: user.ID,
	}
	if err := db.CreateImage(&image); err != nil {
		return err
	}
	return render(r, w, generate.GalleryImage(image))
}

func HandleGenerateImageStatus(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return err
	}
	image, err := db.GetImageByID(id)
	if err != nil {
		return err
	}
	slog.Info("checking image status", "id", id)
	return render(r, w, generate.GalleryImage(image))
}
