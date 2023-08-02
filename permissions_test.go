package client

import "testing"

func TestClient_GetUsers(t *testing.T) {
	resp, err := client.GetUsers(
		&Page{})
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

func TestClient_GetRoles(t *testing.T) {
	resp, err := client.GetRoles(&Page{})
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range resp.PageItems {
		t.Logf("role:%s username:%s\n", v.Role, v.Username)
	}
}

func TestClient_CreateRole(t *testing.T) {
	err := client.CreateRoles(&CreateRoleRequest{
		Role:     "dev1",
		Username: "test",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_DeleteRole(t *testing.T) {
	err := client.DeleteRoles(&DeleteRoleRequest{
		Role:     "dev1",
		Username: "test",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_GetPermissions(t *testing.T) {
	resp, err := client.GetPermissions(&Page{})
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range resp.PageItems {
		t.Logf("permission:%s", v)
	}
}

func TestClient_CreatePermission(t *testing.T) {
	err := client.CreatePermission(&CreatePermissionRequest{
		Role:        "dev",
		NamespaceId: "test",
		Action:      "rw",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_DeletePermission(t *testing.T) {
	err := client.DeletePermission(&DeletePermissionRequest{
		Role:     "dev",
		Resource: "test:*:*",
		Action:   "rw",
	})
	if err != nil {
		t.Fatal(err)
	}
}
