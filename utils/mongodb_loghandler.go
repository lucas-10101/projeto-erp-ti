package utils

import (
	"context"
	"erp/organization-api/data/models"
	"fmt"
	"log/slog"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

// MongoDBLogHandler is a custom slog handler that writes logs to MongoDB.

type MongoDBLogHandler struct {
	ApplicationName string
	DatabaseName    string
	CollectionName  string
	Client          *mongo.Client
	Level           slog.Level
	Attrs           []slog.Attr
	GroupName       string
}

func (h *MongoDBLogHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= h.Level
}

func (h *MongoDBLogHandler) Handle(ctx context.Context, record slog.Record) error {

	collection := h.Client.Database(h.DatabaseName).Collection(h.CollectionName)

	attrs := map[string]string{}
	for _, attr := range h.Attrs {
		attrs[attr.Key] = fmt.Sprintf("%v", attr.Value.Any())
	}

	_, err := collection.InsertOne(context.TODO(), &models.LogModel{
		ApplicationName: h.ApplicationName,
		GroupName:       h.GroupName,
		Level:           record.Level.String(),
		Attrs:           attrs,
		Time:            record.Time,
		Message:         record.Message,
	})

	fmt.Println(err)

	return err
}

func (h *MongoDBLogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &MongoDBLogHandler{
		Client:          h.Client,
		Level:           h.Level,
		Attrs:           append(h.Attrs, attrs...),
		GroupName:       h.GroupName,
		ApplicationName: h.ApplicationName,
		DatabaseName:    h.DatabaseName,
		CollectionName:  h.CollectionName,
	}
}

func (h *MongoDBLogHandler) WithGroup(name string) slog.Handler {
	return &MongoDBLogHandler{
		Client:          h.Client,
		Level:           h.Level,
		Attrs:           h.Attrs,
		GroupName:       name,
		ApplicationName: h.ApplicationName,
		DatabaseName:    h.DatabaseName,
		CollectionName:  h.CollectionName,
	}
}
