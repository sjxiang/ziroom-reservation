package types

import "go.mongodb.org/mongo-driver/bson/primitive"


// 社区
type Community struct {
	ID       primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string               `bson:"name"          json:"name"`
	Location string               `bson:"location"      json:"location"`
	Rooms    []primitive.ObjectID `bson:"rooms"         json:"rooms"`
	Rating   int                  `bson:"rating"        json:"rating"`  // 社区评级
}
