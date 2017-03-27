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
	result, err := g.check("users", name)

	if !result {
		return result, err
	}

	result, err = g.check("orgs", name)

	return result, err
}

func (g *Github) check(entity, name string) (bool, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/%s/%s", entity, name))
	if err != nil {
		// TODO: wrap error
		return false, errors.New("Cannot determine name availability")
	}
	resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return true, nil
	} else if resp.StatusCode == http.StatusOK {
		return false, nil
	}

	return false, errors.New("Cannot determine name availability")
}
