package client

import "testing"

func TestClient_RegisterInstance(t *testing.T) {
	err := client.RegisterInstance(
		&RegisterInstanceRequest{
			IP:          "127.0.0.1",
			Port:        8848,
			ServiceName: "test",
			Weight:      1,
			Enable:      true,
		})
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_DeregisterInstance(t *testing.T) {
	err := client.DeregisterInstance(
		&DeregisterInstanceRequest{
			IP:          "127.0.0.1",
			Port:        8848,
			ServiceName: "test",
		})
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_ModifyInstance(t *testing.T) {
	err := client.ModifyInstance(
		&ModifyInstanceRequest{
			IP:          "127.1.0.1",
			Port:        8848,
			ServiceName: "test",
			Weight:      2,
			Enable:      true,
		})
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_GetInstances(t *testing.T) {
	resp, err := client.GetInstances(
		&GetInstancesRequest{
			ServiceName: "test",
		})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", resp)
}
