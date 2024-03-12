package sb

import (
	"os"

	"github.com/nedpals/supabase-go"
)

const BaseAuthURL = "https://iytojetxrqffvptgjyts.supabase.co/auth/v1/recover"

const ResetPasswordEndPoint = "auth/v1/recover"

var Client *supabase.Client

func Init() error {
	sbHost := os.Getenv("SUPABASE_URL")
	sbSecret := os.Getenv("SUPABASE_SECRET")
	Client = supabase.CreateClient(sbHost, sbSecret)
	return nil
}
