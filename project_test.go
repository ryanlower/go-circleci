package circleci

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecentBuilds(t *testing.T) {
	c := setup()

	builds := c.RecentBuilds("ryanlower", "go-circleci")
	log.Print(builds)

	// TODO, fix these tests. These are completely horrible.
	// They actively rely on the state of the build that's testing them!
	lastBuild := builds[0]
	assert.Equal(t, lastBuild.Branch, "master")
	assert.Equal(t, lastBuild.Status, "running")
}
