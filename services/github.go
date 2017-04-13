package services

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
)

// Github service name checker
type Github struct {
	token string

	logger logrus.FieldLogger
}

// NewGithub returns a new Github checker
func NewGithub(token string, logger logrus.FieldLogger) *Github {
	return &Github{token, logger}
}

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
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.github.com/%s/%s", entity, name), nil)
	if err != nil {
		g.logger.Error(err)

		return false, err
	}

	if g.token != "" {
		req.Header.Add("Authorization", fmt.Sprintf("token %s", g.token))
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		g.logger.Error(err)

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
