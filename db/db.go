package db


const (
	MongoDBNameEnvName = "MONGO_DB_NAME"
)

type Pagination struct {
	Limit int64
	Page  int64
}

type Map map[string]any


type Store struct {
	User      UserStore
	Room      RoomStore
	Community CommunityStore
	Booking   BookingStore
}

func NewStore(user UserStore, room RoomStore, community CommunityStore, booking BookingStore) *Store {
	return &Store{
		User:      user,
		Room:      room,
		Community: community,
		Booking:   booking,
	}
}