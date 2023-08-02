package client

import "testing"

func TestClient_GetConfig(t *testing.T) {
	resp, err := client.GetConfig(&ConfigBase{DataId: "test"})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v", resp)
}

func TestClient_ListenConfig(t *testing.T) {
	resp, err := client.ListenConfig(&ListeningConfigs{
		ConfigBase: ConfigBase{
			DataId: "test",
		},
		ContentMD5: "67babf3f2726b3b92c1eceed1fe74e39",
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v", resp)
}

func TestClient_PublishConfig(t *testing.T) {
	err := client.PublishConfig(&PublishConfigRequest{
		ConfigBase: ConfigBase{
			DataId: "dev",
		},
		Content:     "text content112",
		ContentType: "yaml",
	})
	if err != nil {
		t.Error(err)
	}
}

func TestClient_DeleteConfig(t *testing.T) {
	err := client.DeleteConfig(&ConfigBase{DataId: "test1"})
	if err != nil {
		t.Error(err)
	}
}

func TestClient_GetConfigHistory(t *testing.T) {
	resp, err := client.GetConfigHistory(&GetConfigHistoryRequest{ConfigBase: ConfigBase{DataId: "test"}})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", resp)
}

func TestClient_GetConfigHistoryDetail(t *testing.T) {
	resp, err := client.GetConfigHistoryDetail(&GetConfigHistoryDetailRequest{
		ConfigBase: ConfigBase{DataId: "test"},
		Nid:        "4",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", resp)
}

func TestClient_GetConfigHistoryPrevious(t *testing.T) {
	resp, err := client.GetConfigHistoryPrevious(&GetConfigHistoryPreviousRequest{
		ConfigBase: ConfigBase{DataId: "test"},
		Id:         1,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", resp)
}
