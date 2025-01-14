package auth

import (
	"os"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
)

func auth() {

	// define google oauth flow vars
	google_key := os.Getenv("GOOGLE_KEY")
	google_secret := os.Getenv("GOOGLE_SECRET")
	google_url := "http://localhost:3000/auth/google/callback"

	goth.UseProviders(google.New(google_key, google_secret, google_url))
}
