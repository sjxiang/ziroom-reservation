package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"context"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/sjxiang/ziroom-reservation/internal/types"
	"github.com/sjxiang/ziroom-reservation/internal/db"
	"github.com/sjxiang/ziroom-reservation/internal/db/fixtures"
	"github.com/sjxiang/ziroom-reservation/pkg/logger"
	"github.com/sjxiang/ziroom-reservation/pkg/mws"
)


func SetUp(t *testing.T) *db.Store {
	if err := godotenv.Load(); err != nil {
		t.Fatal(err)
	}
	
	var (
		ctx           = context.Background()
		mongoEndpoint = os.Getenv("MONGO_DB_URL")
		mongoDBName   = os.Getenv("MONGO_DB_NAME")
	)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoEndpoint))
	if err != nil {
		t.Fatal(err)
	}
	if err := client.Database(mongoDBName).Drop(ctx); err != nil {
		t.Fatal(err)
	}
	
	sugaredLogger := logger.NewSugardLogger()
	
	userStore      := db.NewMongoUserStoreImpl(sugaredLogger, client)
	communityStore := db.NewMongoCommunityStoreImpl(sugaredLogger, client)
	roomStore      := db.NewMongoRoomStoreImpl(sugaredLogger, client, communityStore)
	bookingStore   := db.NewMongoBookingStoreImpl(sugaredLogger, client)
	
	return db.NewStore(userStore, roomStore, communityStore, bookingStore)
}

func TestUserGetBooking(t *testing.T) {
	store := SetUp(t)

	client := &http.Client{}
	
	var (
		nonAuthUser    = fixtures.AddUser(store, "Jimmy", "watercooler", false)
		user           = fixtures.AddUser(store, "james", "foo", false)
		community      = fixtures.AddCommunity(store, "bar hotel", "a", 4, nil)
		room           = fixtures.AddRoom(store, "small", true, 4.4, community.ID)
		from           = time.Now()
		till           = from.AddDate(0, 0, 5)
		booking        = fixtures.AddBooking(store, user.ID, room.ID, from, till)
	)
	req := httptest.NewRequest("GET", fmt.Sprintf("http://0.0.0.0:8001/api/v1/admin/bookings/%s", booking.ID.Hex()), nil)
	req.Header.Add("X-Api-Token", mws.CreateTokenFromUser(user))
	
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("non 200 code got %d", resp.StatusCode)
	}

	var bookingResp *types.Booking
	if err := json.NewDecoder(resp.Body).Decode(&bookingResp); err != nil {
		t.Fatal(err)
	}
	if bookingResp.ID != booking.ID {
		t.Fatalf("expected %s got %s", booking.ID, bookingResp.ID)
	}
	if bookingResp.UserID != booking.UserID {
		t.Fatalf("expected %s got %s", booking.UserID, bookingResp.UserID)
	}
	
	req = httptest.NewRequest("GET", fmt.Sprintf("/%s", booking.ID.Hex()), nil)
	req.Header.Add("X-Api-Token", mws.CreateTokenFromUser(nonAuthUser))
	
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode == http.StatusOK {
		t.Fatalf("expected a non 200 status code got %d", resp.StatusCode)
	}
}
