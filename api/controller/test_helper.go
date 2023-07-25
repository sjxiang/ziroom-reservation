package controller

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	
	"github.com/sjxiang/ziroom-reservation/internal/db"
	"github.com/sjxiang/ziroom-reservation/pkg/logger"
)

type testdb struct {
	client *mongo.Client
	logger *zap.SugaredLogger

	*db.Store
}

func (tdb *testdb) teardown(t *testing.T) {
	dbname := os.Getenv(db.MongoDBNameEnvName)
	if err := tdb.client.Database(dbname).Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}
}

func setup(t *testing.T) *testdb {
	if err := godotenv.Load("../.env"); err != nil {
		t.Error(err)
	}
	dburi := os.Getenv("MONGO_DB_URL")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}
	sugaredLogger := logger.NewSugardLogger()
	communityStore := db.NewMongoCommunityStoreImpl(sugaredLogger, client)
	
	return &testdb{
		client: client,
		logger: sugaredLogger,
		Store:  &db.Store{
			Community:   communityStore,
			User:        db.NewMongoUserStoreImpl(sugaredLogger, client),
			Room:        db.NewMongoRoomStoreImpl(sugaredLogger, client, communityStore),
			Booking:     db.NewMongoBookingStoreImpl(sugaredLogger, client),
		},
	}
}
