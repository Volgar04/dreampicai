package handler

import (
	"net/http"

	"github.com/Volgar04/dreampicai/view/generate"
)

func HandleGenerateIndex(w http.ResponseWriter, r *http.Request) error {
	return render(r, w, generate.Index())
}
