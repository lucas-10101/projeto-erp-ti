package models

import (
	"time"
)

// Default LogModel structure application logs
type LogModel struct {
	ApplicationName string            `bson:"application_name"`
	GroupName       string            `bson:"group_name"`
	Level           string            `bson:"level"`
	Attrs           map[string]string `bson:"attrs"`
	Time            time.Time         `bson:"timestamp"`
	Message         string            `bson:"message"`
}
