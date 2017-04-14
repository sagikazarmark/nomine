package services

import (
	"github.com/Sirupsen/logrus"
	"github.com/dnsimple/dnsimple-go/dnsimple"
)

// Dnsimple service name checker
type Dnsimple struct {
	client    *dnsimple.Client
	accountID string
	tld       string

	logger logrus.FieldLogger
}

// NewDnsimple returns a new Dnsimple checker
func NewDnsimple(client *dnsimple.Client, accountID string, tld string, logger logrus.FieldLogger) *Dnsimple {
	return &Dnsimple{client, accountID, tld, logger}
}

// Check implements the NameChecker interface
func (d *Dnsimple) Check(name string) (bool, error) {
	result, err := d.client.Registrar.CheckDomain(d.accountID, name)
	if err != nil {
		d.logger.Error(err)

		return false, err
	}

	return result.Data.Available, nil
}
