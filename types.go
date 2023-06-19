package client

type ConfigBase struct {
	DataId string `json:"dataId" validate:"required"`
	Group  string `json:"group" validate:"required" default:"DEFAULT_GROUP"`
	Tenant string `json:"tenant" validate:"omitempty" default:"public"`
}

var StringResult string

var BoolResult bool

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
	Page
}

type Page struct {
	PageNo   int `json:"pageNo" validate:"omitempty" default:"1"`
	PageSize int `json:"pageSize" validate:"omitempty,max=500" default:"500"`
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

type User struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type GetUsersResponse struct {
	TotalCount     int `json:"totalCount"`
	PageNumber     int `json:"pageNumber"`
	PagesAvailable int `json:"pagesAvailable"`
	PageItems      []struct {
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"pageItems"`
}

type DeleteUserRequest struct {
	Username string `json:"username" validate:"required"`
}

type GetRolesResponse struct {
	TotalCount     int `json:"totalCount"`
	PageNumber     int `json:"pageNumber"`
	PagesAvailable int `json:"pagesAvailable"`
	PageItems      []struct {
		Role     string `json:"role"`
		Username string `json:"username"`
	} `json:"pageItems"`
}

type CreateRoleRequest struct {
	Role     string `json:"role" validate:"required"`
	Username string `json:"username" validate:"required"`
}

type DeleteRoleRequest struct {
	Role     string `json:"role" validate:"required"`
	Username string `json:"username" validate:"required"`
}

type GetPermissions struct {
	TotalCount     int `json:"totalCount"`
	PageNumber     int `json:"pageNumber"`
	PagesAvailable int `json:"pagesAvailable"`
	PageItems      []struct {
		Role     string `json:"role"`
		Resource string `json:"resource"`
		Action   string `json:"action"`
	} `json:"pageItems"`
}

type CreatePermissionRequest struct {
	Role        string `json:"role" validate:"required"`
	NamespaceId string `json:"namespaceId" validate:"required"`
	Action      string `json:"action" validate:"required,oneof=r w rw"`
}

type DeletePermissionRequest struct {
	Role     string `json:"role" validate:"required"`
	Resource string `json:"resource" validate:"required"`
	Action   string `json:"action" validate:"required,oneof=r w rw"`
}

type RegisterInstanceRequest struct {
	IP          string  `json:"ip" validate:"required"`
	Port        int     `json:"port" validate:"required"`
	NamespaceId string  `json:"namespaceId" validate:"omitempty"`
	Weight      float64 `json:"weight" validate:"omitempty"`
	Enable      bool    `json:"enable" validate:"omitempty"`
	Healthy     bool    `json:"healthy" validate:"omitempty"`
	Metadata    string  `json:"metadata" validate:"omitempty"`
	ClusterName string  `json:"clusterName" validate:"omitempty"`
	ServiceName string  `json:"serviceName" validate:"required"`
	GroupName   string  `json:"groupName" validate:"omitempty"`
	Ephemeral   bool    `json:"ephemeral" validate:"omitempty"`
}

type DeregisterInstanceRequest struct {
	IP          string `json:"ip" validate:"required"`
	Port        int    `json:"port" validate:"required"`
	NamespaceId string `json:"namespaceId" validate:"omitempty"`
	ClusterName string `json:"clusterName" validate:"omitempty"`
	ServiceName string `json:"serviceName" validate:"required"`
	GroupName   string `json:"groupName" validate:"omitempty"`
	Ephemeral   bool   `json:"ephemeral" validate:"omitempty"`
}

type ModifyInstanceRequest struct {
	IP          string  `json:"ip" validate:"required"`
	Port        int     `json:"port" validate:"required"`
	NamespaceId string  `json:"namespaceId" validate:"omitempty"`
	Weight      float64 `json:"weight" validate:"omitempty"`
	Enable      bool    `json:"enable" validate:"omitempty"`
	Healthy     bool    `json:"healthy" validate:"omitempty"`
	Metadata    string  `json:"metadata" validate:"omitempty"`
	ClusterName string  `json:"clusterName" validate:"omitempty"`
	ServiceName string  `json:"serviceName" validate:"required"`
	GroupName   string  `json:"groupName" validate:"omitempty"`
	Ephemeral   bool    `json:"ephemeral" validate:"omitempty"`
}

type GetInstancesRequest struct {
	ServiceName string `json:"serviceName" validate:"required"`
	GroupName   string `json:"groupName" validate:"omitempty"`
	NamespaceId string `json:"namespaceId" validate:"omitempty"`
	Clusters    string `json:"clusters" validate:"omitempty"`
	HealthyOnly bool   `json:"healthyOnly" validate:"omitempty"`
}

type GetInstancesResponse struct {
	Name        string `json:"name"`
	GroupName   string `json:"groupName"`
	Clusters    string `json:"clusters"`
	CacheMillis int    `json:"cacheMillis"`
	Hosts       []struct {
		InstanceId  string  `json:"instanceId"`
		Ip          string  `json:"ip"`
		Port        int     `json:"port"`
		Weight      float64 `json:"weight"`
		Healthy     bool    `json:"healthy"`
		Enabled     bool    `json:"enabled"`
		Ephemeral   bool    `json:"ephemeral"`
		ClusterName string  `json:"clusterName"`
		ServiceName string  `json:"serviceName"`
		Metadata    struct {
		} `json:"metadata"`
		InstanceHeartBeatInterval int    `json:"instanceHeartBeatInterval"`
		InstanceIdGenerator       string `json:"instanceIdGenerator"`
		InstanceHeartBeatTimeOut  int    `json:"instanceHeartBeatTimeOut"`
		IpDeleteTimeout           int    `json:"ipDeleteTimeout"`
	} `json:"hosts"`
	LastRefTime              int64  `json:"lastRefTime"`
	Checksum                 string `json:"checksum"`
	AllIPs                   bool   `json:"allIPs"`
	ReachProtectionThreshold bool   `json:"reachProtectionThreshold"`
	Valid                    bool   `json:"valid"`
}
