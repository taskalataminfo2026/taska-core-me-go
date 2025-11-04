package providers

import (
	"taska-core-me-go/cmd/api/clients/rusty"
)

func GetRustyClient() *rusty.RustyClient {
	return &rusty.RustyClient{}
}
