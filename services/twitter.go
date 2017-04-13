package services

import (
	"github.com/ChimeraCoder/anaconda"
	"github.com/Sirupsen/logrus"
)

// Twitter service name checker
type Twitter struct {
	client *anaconda.TwitterApi

	logger logrus.FieldLogger
}

// NewTwitter returns a new Twitter checker
func NewTwitter(client *anaconda.TwitterApi, logger logrus.FieldLogger) *Twitter {
	return &Twitter{client, logger}
}

// Check implements the NameChecker interface
func (t *Twitter) Check(name string) (bool, error) {
	_, err := t.client.GetUsersShow(name, nil)
	if err != nil {
		if err, ok := err.(*anaconda.ApiError); ok {
			if err.StatusCode == 404 {
				return true, nil
			}
		}

		t.logger.Error(err)

		return false, err
	}

	return false, nil
}
