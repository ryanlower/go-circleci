package circleci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMe(t *testing.T) {
	c := setup()

	user := c.Me()

	assert.Equal(t, user.Name, "Go Circle CI")
}
