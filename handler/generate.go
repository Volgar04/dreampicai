package handler

import (
	"context"
	"database/sql"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/Volgar04/dreampicai/db"
	"github.com/Volgar04/dreampicai/pkg/validate"
	"github.com/Volgar04/dreampicai/types"
	"github.com/Volgar04/dreampicai/view/generate"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
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
	amount, _ := strconv.Atoi(r.FormValue("amount"))
	params := generate.FormParams{
		Prompt: r.FormValue("prompt"),
		Amount: amount,
	}
	var errors generate.FormErrors
	if amount <= 0 || amount > 8 {
		errors.Amount = "Please enter a valid amount"
		return render(r, w, generate.Form(params, errors))
	}
	ok := validate.New(params, validate.Fields{
		"Prompt": validate.Rules(validate.Min(10), validate.Max(200)),
	}).Validate(&errors)
	if !ok {
		return render(r, w, generate.Form(params, errors))
	}
	err := db.Bun.RunInTx(r.Context(), &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		batchID := uuid.New()
		for i := 0; i < params.Amount; i++ {
			image := types.Image{
				Prompt:  params.Prompt,
				Status:  types.ImageStatusPending,
				UserID:  user.ID,
				BatchID: batchID,
			}
			if err := db.CreateImage(&image); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return hxRedirect(w, r, "/generate")
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
