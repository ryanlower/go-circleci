package circleci

import (
	"fmt"
	"log"
)

type Build struct {
	Number int    `json:"build_num"`
	URL    string `json:"build_url"`
	Branch string `json:"branch"`
	Status string `json:"status"`
}

// RecentBuilds fetches recent builds for the specified project
// (by username & repo name)
// See https://circleci.com/docs/api#recent-builds-project
// TODO, allow customization of params (e.g. limit, offset, filter)
func (c *Client) RecentBuilds(username, project string) []Build {
	builds := new([]Build)

	path := fmt.Sprintf("/project/%s/%s", username, project)
	err := c.get(path, builds)
	if err != nil {
		log.Print("ERROR ", err.Error())
	}

	return *builds
}
