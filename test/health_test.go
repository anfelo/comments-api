//go:build e2e
// +build e2e

package test

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestHealthEndpoint(t *testing.T) {
	client := resty.New()
	resp, err := client.R().Get(BASE_URL + "/api/health")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 200, resp.StatusCode())
}
