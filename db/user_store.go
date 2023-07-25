package db

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"

	"github.com/sjxiang/ziroom-reservation/types"
)

const userColl = "users"

type Map map[string]any

type Dropper interface {
	Drop(context.Context) error
}

type UserStore interface {
	Dropper

	GetUserByEmail(context.Context, string) (*types.User, error)
	GetUserByID(context.Context, string) (*types.User, error)
	GetUsers(context.Context) ([]*types.User, error)
	InsertUser(context.Context, *types.User) (*types.User, error)
	DeleteUser(context.Context, string) error
	UpdateUser(ctx context.Context, filter Map, params types.UpdateUserParams) error
}

type MongoUserStoreImpl struct {
	coll   *mongo.Collection
	logger *zap.SugaredLogger
}

func NewMongoUserStoreImpl(logger *zap.SugaredLogger, client *mongo.Client) *MongoUserStoreImpl {
	dbname := os.Getenv(MongoDBNameEnvName)
	return &MongoUserStoreImpl{
		logger: logger,
		coll:   client.Database(dbname).Collection(userColl),
	}
}

func (impl *MongoUserStoreImpl) Drop(ctx context.Context) error {
	impl.logger.Infow("--- dropping user collection")
	return impl.coll.Drop(ctx)
}

func (impl *MongoUserStoreImpl) UpdateUser(ctx context.Context, filter Map, params types.UpdateUserParams) error {
	oid, err := primitive.ObjectIDFromHex(filter["_id"].(string))
	if err != nil {
		return err
	}
	filter["_id"] = oid
	update := bson.M{"$set": params.ToBSON()}
	_, err = impl.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (impl *MongoUserStoreImpl) DeleteUser(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	// TODO: Maybe its a good idea to handle if we did not delete any user.
	// maybe log it or something??
	_, err = impl.coll.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}
	return nil
}

func (impl *MongoUserStoreImpl) InsertUser(ctx context.Context, user *types.User) (*types.User, error) {
	res, err := impl.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = res.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (impl *MongoUserStoreImpl) GetUsers(ctx context.Context) ([]*types.User, error) {
	cur, err := impl.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var users []*types.User
	if err := cur.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (impl *MongoUserStoreImpl) GetUserByEmail(ctx context.Context, email string) (*types.User, error) {
	var user types.User
	if err := impl.coll.FindOne(ctx, bson.M{"email": email}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (impl *MongoUserStoreImpl) GetUserByID(ctx context.Context, id string) (*types.User, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var user types.User
	if err := impl.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
