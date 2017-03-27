package services

import (
	"errors"
	"fmt"
	"net/http"
)

// Github service name checker
// TODO: use Github API token
type Github struct{}

// Check implements the NameChecker interface
func (g *Github) Check(name string) (bool, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s", name))
	if err != nil {
		// TODO: wrap error
		return false, errors.New("Cannot determine name availability")
	}
	resp.Body.Close()

	if resp.StatusCode != http.StatusNotFound {
		return false, nil
	}

	resp, err = http.Get(fmt.Sprintf("https://api.github.com/orgs/%s", name))
	if err != nil {
		// TODO: wrap error
		return false, errors.New("Cannot determine name availability")
	}
	resp.Body.Close()

	if resp.StatusCode != http.StatusNotFound {
		return false, nil
	}

	return true, nil
}
