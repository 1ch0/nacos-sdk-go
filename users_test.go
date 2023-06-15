package client

import "testing"

func TestClient_GetUsers(t *testing.T) {
	resp, err := client.GetUsers(
		&GetUsersRequest{})
	if err != nil {
		t.Error(err)
	}
	for _, v := range resp.PageItems {
		t.Logf("username:%s\n", v.Username)
	}
}

func TestClient_CreateUser(t *testing.T) {
	err := client.CreateUser(&User{
		Username: "test",
		Password: "test",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_PutUser(t *testing.T) {
	err := client.PutUser(&User{
		Username: "test",
		Password: "testzxc",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_DeleteUser(t *testing.T) {
	err := client.DeleteUser(&DeleteUserRequest{
		Username: "test1",
	})
	if err != nil {
		t.Fatal(err)
	}
}
