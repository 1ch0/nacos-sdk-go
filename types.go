package client

type ConfigBase struct {
	DataId string `json:"dataId" validate:"required"`
	Group  string `json:"group" validate:"required" default:"DEFAULT_GROUP"`
	Tenant string `json:"tenant" validate:"omitempty" default:"public"`
}

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
	TokenTtl    int    `json:"tokenTtl"`
	GlobalAdmin bool   `json:"globalAdmin"`
	Username    string `json:"username"`
}

type PublishConfigRequest struct {
	ConfigBase
	Content     string `json:"content" validate:"required"`
	ContentType string `json:"contentType" validate:"omitempty"`
}

type ListeningConfigs struct {
	ConfigBase
	ContentMD5 string `json:"contentMD5"  validate:"required"`
}

type GetConfigHistoryRequest struct {
	ConfigBase
	PageNo   int `json:"pageNo" validate:"omitempty" default:"1"`
	PageSize int `json:"pageSize" validate:"omitempty,max=500" default:"100"`
}

type GetConfigHistoryResponse struct {
	TotalCount     int `json:"totalCount"`
	PageNumber     int `json:"pageNumber"`
	PagesAvailable int `json:"pagesAvailable"`
	PageItems      []struct {
		Id               string      `json:"id"`
		LastId           int         `json:"lastId"`
		DataId           string      `json:"dataId"`
		Group            string      `json:"group"`
		Tenant           string      `json:"tenant"`
		AppName          string      `json:"appName"`
		Md5              interface{} `json:"md5"` //
		Content          interface{} `json:"content"`
		SrcIp            string      `json:"srcIp"`
		SrcUser          string      `json:"srcUser"`
		OpType           string      `json:"opType"`
		CreatedTime      string      `json:"createdTime"`
		LastModifiedTime string      `json:"lastModifiedTime"`
	} `json:"pageItems"`
}

type GetConfigHistoryDetailRequest struct {
	Nid string `json:"nid" validate:"required"`
	ConfigBase
}

type GetConfigHistoryDetailResponse struct {
	Id               string      `json:"id"`
	LastId           int         `json:"lastId"`
	DataId           string      `json:"dataId"`
	Group            string      `json:"group"`
	Tenant           string      `json:"tenant"`
	AppName          string      `json:"appName"`
	Md5              string      `json:"md5"`
	Content          string      `json:"content"`
	SrcIp            string      `json:"srcIp"`
	SrcUser          interface{} `json:"srcUser"`
	OpType           string      `json:"opType"`
	CreatedTime      string      `json:"createdTime"`
	LastModifiedTime string      `json:"lastModifiedTime"`
}

type GetConfigHistoryPreviousRequest struct {
	Id int `json:"id" validate:"required"` // 配置 ID
	ConfigBase
}

type GetNamespacesResponse struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
	Data    []struct {
		Namespace         string `json:"namespace"`
		NamespaceShowName string `json:"namespaceShowName"`
		Quota             int    `json:"quota"`
		ConfigCount       int    `json:"configCount"`
		Type              int    `json:"type"`
	} `json:"data"`
}

type CreateNamespaceRequest struct {
	//customNamespaceId	字符串	是	命名空间ID
	//namespaceName	字符串	是	命名空间名
	//namespaceDesc	字符串	否	命名空间描述
	CustomNamespaceId string `json:"customNamespaceId" validate:"required"`
	NamespaceName     string `json:"namespaceName" validate:"required"`
	NamespaceDesc     string `json:"namespaceDesc" validate:"omitempty"`
}

type PutNamespaceRequest struct {
	Namespace         string `json:"namespace" validate:"required"`         // 命名空间 ID
	NamespaceShowName string `json:"namespaceShowName" validate:"required"` // 命名空间展示名
	NamespaceDesc     string `json:"namespaceDesc" validate:"required"`
}

type DeleteNamespaceRequest struct {
	NamespaceId string `json:"namespaceId" validate:"required"` // 命名空间 ID
}
