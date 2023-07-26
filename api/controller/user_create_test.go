package controller

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
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
	req, err := http.NewRequest("POST", "http://0.0.0.0:8001/api/v1/user", bytes.NewReader(buf))
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

	bodyText, err := ioutil.ReadAll(req.Body)
	if err != nil {
		t.Fatal(err)
	}
	
	var user types.User
	// bug - 塞了提示信息，所以 json 序列化不行
	err = json.Unmarshal(bodyText, &user)

	// 方法 - 2
	// if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
	//     t.Fatal(err)
	// }
	
	if err != nil {
		t.Fatal(err)
	}

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
		t.Errorf("expected lastname %s but got %s", params.LastName, user.LastName)
	}
	if user.Email != params.Email {
		t.Errorf("expected email %s but got %s", params.Email, user.Email)
	}
}
