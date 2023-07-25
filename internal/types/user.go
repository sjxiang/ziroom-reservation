package types

import (
	"fmt"
	"regexp"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type UpdateUserParams struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (p UpdateUserParams) ToBSON() bson.M {
	m := bson.M{}
	if len(p.FirstName) > 0 {
		m["first_name"] = p.FirstName
	}
	if len(p.LastName) > 0 {
		m["last_name"] = p.LastName
	}
	return m
}

type CreateUserParams struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name"  binding:"required"`
	Email     string `json:"email"      binding:"required"`
	Password  string `json:"password"   binding:"required,gte=6"` // 必填，大于等于 6
}

// Todo - validator 参数校验
func (params CreateUserParams) Validate() error {

	if !isEmailValid(params.Email) {
		return fmt.Errorf("email %s is invalid", params.Email)
	}
	return nil
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"      json:"id,omitempty"`
	FirstName         string             `bson:"first_name"         json:"first_name"`
	LastName          string             `bson:"last_ame"           json:"last_name"`
	Email             string             `bson:"email"              json:"email"`
	EncryptedPassword string             `bson:"encrypted_password" json:"-"`
	IsAdmin           bool               `bson:"is_admin"           json:"is_admin"`
	CreatedAt         int64              `bson:"created_at"         json:"created_at"`
	UpdatedAt         int64              `bson:"updated_at"         json:"updated_at"`
}

func NewUserFromParams(params CreateUserParams) (*User, error) {

	now := time.Now().Unix()
	encpw, _ := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)

	return &User{
		FirstName:         params.FirstName,
		LastName:          params.LastName,
		Email:             params.Email,
		EncryptedPassword: string(encpw),
		CreatedAt:         now,
		UpdatedAt:         now,
	}, nil
}
