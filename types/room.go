package types

import "go.mongodb.org/mongo-driver/bson/primitive"


// 房间
type Room struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"   json:"id,omitempty"`
	Size           string             `bson:"size"            json:"size"`
	SingleBathroom bool               `bson:"single_bathroom" json:"single_bathroom"`  // 独卫
	Price          float64            `bson:"price"           json:"price"`
	CommunityID    primitive.ObjectID `bson:"community_id"    json:"community_id"`
}
