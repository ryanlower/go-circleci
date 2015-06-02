package circleci

import (
	"log"
)

// User represents a CircleCI user
type User struct {
	Name string `json:"name"`
}

// Me fetches the currently signed in user (as authenticated via
// the API token)
// See https://circleci.com/docs/api#user
func (c *Client) Me() User {
	u := &User{}
	err := c.get("/me", u)
	if err != nil {
		log.Print("ERROR ", err.Error())
	}

	return *u
}
