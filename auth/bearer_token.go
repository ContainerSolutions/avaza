package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
	//  _ "github.com/motemen/go-loghttp/global"
)

type BearerToken struct {
	AccessToken  string    `json:"access_token"`
	ExpiresAt    time.Time `json:"expires_at"`
	RefreshToken string    `json:"refresh_token"`
}

type getBearerResponse struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int64  `json:"expires_in"`
	RefreshToken     string `json:"refresh_token"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func getBearerFromTempCode(code tempCode) (*BearerToken, error) {
	form := url.Values{}
	form.Add("code", string(code))
	form.Add("grant_type", "authorization_code")
	form.Add("redirect_uri", AvazaRedirectUrl)
	form.Add("client_id", AvazaAppId)
	form.Add("client_secret", AvazaAppSecret)

	req, err := http.NewRequest("POST", fmt.Sprintf("https://any.avaza.com/oauth2/token?%s", form.Encode()), nil)

	hc := http.Client{}
	resp, err := hc.Do(req)

	if err != nil {
		return nil, err
	}

	var body getBearerResponse
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	return &BearerToken{
		AccessToken:  body.AccessToken,
		ExpiresAt:    time.Unix(time.Now().Unix()+body.ExpiresIn, 0),
		RefreshToken: body.RefreshToken,
	}, nil
}
