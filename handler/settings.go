package handler

import (
	"net/http"

	"github.com/Volgar04/dreampicai/view/settings"
)

func HandleSettingsIndex(w http.ResponseWriter, r *http.Request) error {
	user := getAtuhenticatedUser(r)
	return render(r, w, settings.Index(user))
}
