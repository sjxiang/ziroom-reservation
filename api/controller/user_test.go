package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	
	"github.com/sjxiang/ziroom-reservation/internal/types"
)

func TestPostUser(t *testing.T) {
	client := &http.Client{}

	params := types.CreateUserParams{
		Email:     "some@foo.com",
		FirstName: "James",
		LastName:  "Foo",
		Password:  "lkdfjkdsjfklfdjkedf",
	}
	buf, _ := json.Marshal(params)

	// 创建请求
	req, err := http.NewRequest("POST", "localhost:8001/api/v1/user", bytes.NewReader(buf))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/json")
	
	// 发起请求
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()  

	var user types.User
	json.NewDecoder(resp.Body).Decode(&user)
	
	if len(user.ID) == 0 {
		t.Errorf("expecting a user id to be set")
	}
	if len(user.EncryptedPassword) > 0 {
		t.Errorf("expecting the EncryptedPassword not to be included in the json response")
	}
	if user.FirstName != params.FirstName {
		t.Errorf("expected firstname %s but got %s", params.FirstName, user.FirstName)
	}
	if user.LastName != params.LastName {
		t.Errorf("expected last name %s but got %s", params.LastName, user.LastName)
	}
	if user.Email != params.Email {
		t.Errorf("expected email %s but got %s", params.Email, user.Email)
	}
}
