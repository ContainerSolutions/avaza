package simple_client

import (
	avaza_client "github.com/moretea/avaza/client/client"
	"github.com/go-openapi/runtime"
)

type Authorizer interface {
  RefreshIfNecessary()
  CreateAuth() runtime.ClientAuthInfoWriterFunc
}

type Client struct {
  authorizer Authorizer
  avazaClient *avaza_client.AvazaAPIDocumentation
}

func NewClient(authorizer Authorizer, client *avaza_client.AvazaAPIDocumentation) *Client {
  return &Client {
    authorizer: authorizer,
    avazaClient: client,
  }
}
