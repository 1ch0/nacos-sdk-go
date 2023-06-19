package client

import "testing"

func TestClient_GetNamespaces(t *testing.T) {
	resp, err := client.GetNamespaces()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", resp)
}

func TestClient_CreateNamespace(t *testing.T) {
	resp, err := client.CreateNamespace(&CreateNamespaceRequest{
		CustomNamespaceId: "dev1",
		NamespaceName:     "test",
		NamespaceDesc:     "test",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", resp)
}

func TestClient_PutNamespace(t *testing.T) {
	resp, err := client.PutNamespace(&PutNamespaceRequest{
		Namespace:         "test",
		NamespaceShowName: "测试",
		NamespaceDesc:     "测试",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", resp)
}

func TestClient_DeleteNamespace(t *testing.T) {
	resp, err := client.DeleteNamespace(&DeleteNamespaceRequest{
		NamespaceId: "test1",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", resp)
}
