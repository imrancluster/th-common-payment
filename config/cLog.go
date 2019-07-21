package config

import (
	log "github.com/sirupsen/logrus"
)

func CLog(actionCode string, msisdn string, context map[string]interface{}) *log.Entry {
	log.SetFormatter(&log.JSONFormatter{})

	data := log.Fields{
		"action":  actionCode,
		"msisdn":  msisdn,
		"context": context,
	}

	return log.WithFields(data)
}
