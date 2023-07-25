package db


import (
	"context"
	"os"
	
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"

	"github.com/sjxiang/ziroom-reservation/internal/types"
)

type CommunityStore interface {
	InsertCommunity(context.Context, *types.Community) (*types.Community, error)
	Update(context.Context, Map, Map) error
	GetCommunitys(context.Context, Map, *Pagination) ([]*types.Community, error)
	GetCommunityByID(context.Context, string) (*types.Community, error)
}

type MongoCommunityStoreImpl struct {
	logger *zap.SugaredLogger
	coll   *mongo.Collection
}

func NewMongoCommunityStoreImpl(logger *zap.SugaredLogger, client *mongo.Client) *MongoCommunityStoreImpl {
	dbname := os.Getenv(MongoDBNameEnvName)
	return &MongoCommunityStoreImpl{
		logger: logger,
		coll:   client.Database(dbname).Collection("community"),
	}
}

func (impl *MongoCommunityStoreImpl) GetCommunityByID(ctx context.Context, id string) (*types.Community, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var community types.Community
	if err := impl.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&community); err != nil {
		return nil, err
	}
	return &community, nil
}

func (impl *MongoCommunityStoreImpl) GetCommunitys(ctx context.Context, filter Map, pag *Pagination) ([]*types.Community, error) {
	opts := options.FindOptions{}
	opts.SetSkip((pag.Page - 1) * pag.Limit)
	opts.SetLimit(pag.Limit)
	
	resp, err := impl.coll.Find(ctx, filter, &opts)
	if err != nil {
		return nil, err
	}
	var communitys []*types.Community
	if err := resp.All(ctx, &communitys); err != nil {
		return nil, err
	}
	return communitys, nil
}

func (impl *MongoCommunityStoreImpl) Update(ctx context.Context, filter Map, update Map) error {
	_, err := impl.coll.UpdateOne(ctx, filter, update)
	return err
}

func (impl *MongoCommunityStoreImpl) InsertCommunity(ctx context.Context, hotel *types.Community) (*types.Community, error) {
	resp, err := impl.coll.InsertOne(ctx, hotel)
	if err != nil {
		return nil, err
	}
	hotel.ID = resp.InsertedID.(primitive.ObjectID)
	return hotel, nil
}
