package authorizer

import (
	"fmt"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/moretea/avaza/client/client"
  "github.com/moretea/avaza/client/simple_client"
	"github.com/moretea/avaza/state"
	"log"
	"time"
)

type Authorizer struct {
	config *state.Config
}

func NewAuthorizer(config *state.Config) *Authorizer {
	return &Authorizer{
		config: config,
	}
}

func (a *Authorizer) RefreshIfNecessary() {
	if a.config.BearerToken.ExpiresAt.Before(time.Now()) {
		log.Fatalf("Bearer token expired at %s. Please log in again", a.config.BearerToken.ExpiresAt)
	}
}

func (a *Authorizer) SimpleClient() *simple_client.Client {
  return simple_client.NewClient(a, a.Client())
}

func (a *Authorizer) Client() *client.AvazaAPIDocumentation {
	return client.NewHTTPClient(strfmt.Default)
}


func (a *Authorizer) CreateAuth() runtime.ClientAuthInfoWriterFunc {
	return func(r runtime.ClientRequest, _ strfmt.Registry) error {
		return r.SetHeaderParam("Authorization", fmt.Sprintf("Bearer %s", a.config.BearerToken.AccessToken))
	}
}
