package services

import (
	"errors"
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/Sirupsen/logrus"
)

type domainCheckResponse struct {
	DomainInfo domainInfo `json:"DomainInfo"`
}

type domainInfo struct {
	DomainAvailability string `json:"domainAvailability"`
	DomainName         string `json:"domainName"`
}

// Whoisxml service name checker
type Whoisxml struct {
	tld      string
	username string
	password string

	logger logrus.FieldLogger
}

// NewWhoisxml returns a new Whoisxml checker
func NewWhoisxml(tld string, username string, password string, logger logrus.FieldLogger) *Whoisxml {
	return &Whoisxml{tld, username, password, logger}
}

// Check implements the NameChecker interface
func (w *Whoisxml) Check(name string) (bool, error) {
	resp, err := http.Post(fmt.Sprintf("https://www.whoisxmlapi.com/whoisserver/WhoisService?domainName=%s.%s&outputFormat=json&cmd=GET_DN_AVAILABILITY&getMode=DNS_AND_WHOIS&username=%s&password=%s", name, w.tld, w.username, w.password), "", nil)
	if err != nil {
		w.logger.Error(err)

		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, errors.New("Cannot determine name availability")
	}

	result := &domainCheckResponse{}

	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		w.logger.Error(err)

		return false, err
	}

	if result.DomainInfo.DomainAvailability == "AVAILABLE" {
		return true, nil
	} else if result.DomainInfo.DomainAvailability == "UNAVAILABLE" {
		return false, nil
	}

	return false, errors.New("Cannot determine name availability")
}
