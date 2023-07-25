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

type RoomStore interface {
	InsertRoom(context.Context, *types.Room) (*types.Room, error)
	GetRooms(context.Context, bson.M) ([]*types.Room, error)
}

type MongoRoomStoreImpl struct {
	logger *zap.SugaredLogger
	coll   *mongo.Collection

	// 依赖项
	CommunityStore
}

func NewMongoRoomStoreImpl(logger *zap.SugaredLogger, client *mongo.Client, communityStore CommunityStore) *MongoRoomStoreImpl {
	dbname := os.Getenv(MongoDBNameEnvName)
	return &MongoRoomStoreImpl{
		logger:         logger,
		coll:           client.Database(dbname).Collection("rooms"),
		CommunityStore: communityStore,
	}
}

func (impl *MongoRoomStoreImpl) GetRooms(ctx context.Context, filter bson.M) ([]*types.Room, error) {
	resp, err := impl.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var rooms []*types.Room
	if err := resp.All(ctx, &rooms); err != nil {
		return nil, err
	}
	return rooms, nil
}

func (impl *MongoRoomStoreImpl) InsertRoom(ctx context.Context, room *types.Room) (*types.Room, error) {
	resp, err := impl.coll.InsertOne(ctx, room)
	if err != nil {
		return nil, err
	}
	room.ID = resp.InsertedID.(primitive.ObjectID)

	// update the community with this room id
	filter := Map{"_id": room.CommunityID}
	update := Map{"$push": bson.M{"rooms": room.ID}}
	if err := impl.CommunityStore.Update(ctx, filter, update); err != nil {
		return nil, err
	}

	return room, nil
}
