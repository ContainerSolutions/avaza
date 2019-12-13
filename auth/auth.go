package auth

import "os"

var AvazaAppId string = "107"

// TODO: extract from repo.
var AvazaAppSecret string = "" // --> links this app to CS workspace
var AvazaRedirectUrl string = "http://localhost:8083"
var AvazaRedirectHost string = "localhost:8083"

func init() {
  AvazaAppSecret = os.Getenv("AVAZA_APP_SECRET")
}

// Start a simple HTTP server to receive the callback.
// Print user instructions, and await a request to the HTTP server.
func Register() (*BearerToken, error) {
	tempCode := getTempCode()
	return getBearerFromTempCode(tempCode)
}
