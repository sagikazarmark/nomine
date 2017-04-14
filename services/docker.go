package services

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Sirupsen/logrus"
)

// Docker service name checker
type Docker struct {
	logger logrus.FieldLogger
}

// NewDocker returns a new Docker checker
func NewDocker(logger logrus.FieldLogger) *Docker {
	return &Docker{logger}
}

// Check implements the NameChecker interface
func (g *Docker) Check(name string) (bool, error) {
	resp, err := http.Get(fmt.Sprintf("https://hub.docker.com/u/%s/", name))
	if err != nil {
		g.logger.Error(err)

		return false, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		g.logger.Error(err)

		return false, err
	}

	result := strings.Index(string(body), "Page Not Found")

	return result > -1, nil
}
