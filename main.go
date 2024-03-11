package main

import (
	"embed"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/Volgar04/dreampicai/db"
	"github.com/Volgar04/dreampicai/handler"
	"github.com/Volgar04/dreampicai/pkg/sb"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

//go:embed public
var FS embed.FS

func main() {
	if err := initEverything(); err != nil {
		log.Fatal(err)
	}

	router := chi.NewMux()
	router.Use(handler.WithUser)

	router.Handle("/*", http.StripPrefix("/", http.FileServer(http.FS(FS))))
	router.Get("/", handler.Make(handler.HandleHomeIndex))
	router.Get("/login", handler.Make(handler.HandleLoginIndex))
	router.Get("/login/provider/google", handler.Make(handler.HandleLoginWithGoogleCreate))
	router.Get("/signup", handler.Make(handler.HandleSignupIndex))
	router.Post("/logout", handler.Make(handler.HandleLogoutCreate))
	router.Post("/login", handler.Make(handler.HandleLoginCreate))
	router.Post("/signup", handler.Make(handler.HandleSignupCreate))
	router.Get("/auth/callback", handler.Make(handler.HandleAuthCallback))
	router.Get("/account/setup", handler.Make(handler.HandleAccountSetupIndex))
	router.Post("/account/setup", handler.Make(handler.HandleAccountSetupCreate))

	router.Group(func(auth chi.Router) {
		auth.Use(handler.WithAccountSetup)
		auth.Get("/", handler.Make(handler.HandleHomeIndex))
		auth.Get("/settings", handler.Make(handler.HandleSettingsIndex))
	})

	port := os.Getenv("HTTP_LISTEN_ADDR")
	slog.Info("application is running", "port", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func initEverything() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	if err := db.Init(); err != nil {
		return err
	}
	return sb.Init()
}
