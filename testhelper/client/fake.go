package client

import (
	"net/http"

	"github.com/magicmemories/go-jerakia"
	"github.com/magicmemories/go-jerakia/testhelper"
)

// Fake token to use.
const Token = "myapp:abcd"

// FakeClient returns a generic client to use for testing.
func FakeClient() *jerakia.Client {
	config := jerakia.ClientConfig{
		Token: Token,
		URL:   testhelper.URL(),
	}

	client := jerakia.NewClient(http.DefaultClient, config)
	return &client
}
