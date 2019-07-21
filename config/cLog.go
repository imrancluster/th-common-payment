package config

/*
	Logging examples
	----------------
	1. config.CLog("VISIT_INDEX_PAGE", "01799997163", map[string]interface{}{"data1": "test1", "data2": "test2"}).Info("Visit Index Page")
	2. config.CLog("INDEX_WARNING", "01799997163", map[string]interface{}{"data1": "t1", "data2": "t2"}).Warning("Index warning")
	3. config.CLog("INDEX_ERROR", "01799997163", map[string]interface{}{"data1": "t1", "data2": "t2"}).Error("Index Error")
*/

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
