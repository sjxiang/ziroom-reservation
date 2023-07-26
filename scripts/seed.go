package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/sjxiang/ziroom-reservation/internal/db"
	"github.com/sjxiang/ziroom-reservation/internal/db/fixtures"
	"github.com/sjxiang/ziroom-reservation/pkg/logger"
	"github.com/sjxiang/ziroom-reservation/pkg/mws"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	
	var (
		ctx           = context.Background()
		mongoEndpoint = os.Getenv("MONGO_DB_URL")
		mongoDBName   = os.Getenv("MONGO_DB_NAME")
	)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoEndpoint))
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Database(mongoDBName).Drop(ctx); err != nil {
		log.Fatal(err)
	}
	sugaredLogger := logger.NewSugardLogger()
	
	userStore      := db.NewMongoUserStoreImpl(sugaredLogger, client)
	communityStore := db.NewMongoCommunityStoreImpl(sugaredLogger, client)
	roomStore      := db.NewMongoRoomStoreImpl(sugaredLogger, client, communityStore)
	bookingStore   := db.NewMongoBookingStoreImpl(sugaredLogger, client)
	store := db.NewStore(userStore, roomStore, communityStore, bookingStore)

	// 注入数据（绕过 JWT）
	user := fixtures.AddUser(store, "xiang", "qq", false)
	fmt.Println("xiang ->", mws.CreateTokenFromUser(user))
	
	admin := fixtures.AddUser(store, "admin", "admin", true)
	fmt.Println("admin ->", mws.CreateTokenFromUser(admin))
	
	community := fixtures.AddCommunity(store, "laimeng", "yuhuatai", 5, nil)
	room := fixtures.AddRoom(store, "large", true, 88.44, community.ID)
	booking := fixtures.AddBooking(store, user.ID, room.ID, time.Now(), time.Now().AddDate(0, 0, 5))  // 这个时间设置可以
	fmt.Println("booking ->", booking.ID)

	for i := 0; i < 100; i++ {
		name := fmt.Sprintf("random community name %d", i)
		location := fmt.Sprintf("location %d", i)
		fixtures.AddCommunity(store, name, location, rand.Intn(5)+1, nil)
	}
}
