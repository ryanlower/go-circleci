package circleci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Shared client setup for tests
func setup() *Client {
	// Client with test API token (for go-circleci user)
	client := NewClient("a6fee742b2ba5ac9885131476e77b00ee26066a7")

	return client
}

func TestNewClient(t *testing.T) {
	client := NewClient("TOKEN")

	assert.NotNil(t, client)
	assert.Equal(t, client.apiToken, "TOKEN")
}
