package auth

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type tempCode string

func getTempCode() tempCode {
	callback := callbackHandler{
		done: make(chan bool, 1),
	}

	srv := &http.Server{
		Addr:    AvazaRedirectHost,
		Handler: &callback,
	}

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("Could not start callback listening sever: %s", err)
		}
	}()

	fmt.Printf("Please visit the addres below to authorize this CLI app to Avaza\n")
	fmt.Printf("  %s\n", registrationURL())

	// Await the one callback we should get
	<-callback.done
	// Then shutdown the server.
	srv.Shutdown(nil)

	return tempCode(callback.tempCode)
}

type callbackHandler struct {
	tempCode string
	done     chan bool
}

func (s *callbackHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	s.tempCode = q.Get("code")
	s.done <- true
}

func registrationURL() string {
	scope := "read,write"
	return fmt.Sprintf("https://any.avaza.com/oauth2/authorize?response_type=code&client_id=%s&redirect_uri=%s&scope=%s", AvazaAppId, url.PathEscape(AvazaRedirectUrl), url.PathEscape(scope))
}
