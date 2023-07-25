package db


import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"

	"github.com/sjxiang/ziroom-reservation/internal/types"
)

type BookingStore interface {
	InsertBooking(context.Context, *types.Booking) (*types.Booking, error)
	GetBookings(context.Context, bson.M) ([]*types.Booking, error)
	GetBookingByID(context.Context, string) (*types.Booking, error)
	UpdateBooking(context.Context, string, bson.M) error
}

type MongoBookingStoreImpl struct {
	logger *zap.SugaredLogger
	coll   *mongo.Collection
}

func NewMongoBookingStoreImpl(logger *zap.SugaredLogger, client *mongo.Client) *MongoBookingStoreImpl {
	dbname := os.Getenv(MongoDBNameEnvName)
	return &MongoBookingStoreImpl{
		logger: logger,
		coll:   client.Database(dbname).Collection("bookings"),
	}
}

func (impl *MongoBookingStoreImpl) UpdateBooking(ctx context.Context, id string, update bson.M) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	m := bson.M{"$set": update}
	_, err = impl.coll.UpdateByID(ctx, oid, m)
	return err
}

func (impl *MongoBookingStoreImpl) GetBookings(ctx context.Context, filter bson.M) ([]*types.Booking, error) {
	curr, err := impl.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var bookings []*types.Booking
	if err := curr.All(ctx, &bookings); err != nil {
		return nil, err
	}
	return bookings, nil
}

func (impl *MongoBookingStoreImpl) GetBookingByID(ctx context.Context, id string) (*types.Booking, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var booking types.Booking
	if err := impl.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&booking); err != nil {
		return nil, err
	}
	return &booking, nil
}

func (impl *MongoBookingStoreImpl) InsertBooking(ctx context.Context, booking *types.Booking) (*types.Booking, error) {
	resp, err := impl.coll.InsertOne(ctx, booking)
	if err != nil {
		return nil, err
	}
	booking.ID = resp.InsertedID.(primitive.ObjectID)
	return booking, nil
}
