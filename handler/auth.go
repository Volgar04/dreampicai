package handler

import (
	"log/slog"
	"net/http"

	"github.com/Volgar04/dreampicai/pkg/sb"
	"github.com/Volgar04/dreampicai/pkg/util"
	"github.com/Volgar04/dreampicai/view/auth"
	"github.com/nedpals/supabase-go"
)

func HandleLoginIndex(w http.ResponseWriter, r *http.Request) error {
	return render(r, w, auth.Login())
}

func HandleLoginCreate(w http.ResponseWriter, r *http.Request) error {
	credentials := supabase.UserCredentials{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	if !util.IsValidateEmail(credentials.Email) {
		return render(r, w, auth.LoginForm(credentials, auth.LoginErrors{
			Email: "Please enter a valid email",
		}))
	}
	if reason, ok := util.ValidatePassword(credentials.Password); !ok {
		return render(r, w, auth.LoginForm(credentials, auth.LoginErrors{
			Password: reason,
		}))
	}
	resp, err := sb.Client.Auth.SignIn(r.Context(), credentials)
	if err != nil {
		slog.Error("login error", "err", err)
		return render(r, w, auth.LoginForm(credentials, auth.LoginErrors{
			InvalidCredentials: "The credentials you have entered are invalid",
		}))
	}
	cookie := &http.Cookie{
		Value:    resp.AccessToken,
		Name:     "at",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	}
	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/", http.StatusSeeOther)
	return nil
}