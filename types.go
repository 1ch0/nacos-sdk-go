package client

type ConfigBase struct {
	DataId string `json:"dataId" validate:"required"`
	Group  string `json:"group" validate:"required" default:"DEFAULT_GROUP"`
	Tenant string `json:"tenant" validate:"omitempty" default:"public"`
}

var StringResult string

var BoolResult bool

type DoParams struct {
	Method      string
	Req         interface{}
	Path        string
	Result      interface{}
	QueryParams map[string]string
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

type InstanceBase struct {
	IP          string `json:"ip" validate:"required"`           // 实例 IP
	Port        int    `json:"port" validate:"required"`         // 实例端口
	NamespaceId string `json:"namespaceId" validate:"omitempty"` // 命名空间 ID
	ClusterName string `json:"clusterName" validate:"omitempty"` // 集群名称
	ServiceName string `json:"serviceName" validate:"required"`  // 服务名称
	GroupName   string `json:"groupName" validate:"omitempty"`   // 分组名称
	Ephemeral   bool   `json:"ephemeral" validate:"omitempty"`   // 是否临时实例
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

type GetInstanceRequest struct {
	InstanceBase
	Healthy bool `json:"healthy" validate:"omitempty"`
}

type GetInstanceResponse struct {
	Metadata struct {
	} `json:"metadata"`
	InstanceId  string  `json:"instanceId"`
	Port        int     `json:"port"`
	Service     string  `json:"service"`
	Healthy     bool    `json:"healthy"`
	Ip          string  `json:"ip"`
	ClusterName string  `json:"clusterName"`
	Weight      float64 `json:"weight"`
}

type SendHeartbeatRequest struct {
	InstanceBase
	HeartBeat HeartBeat `json:"beat" validate:"required"`
}

type HeartBeat struct {
	Cluster  string `json:"cluster"`
	Ip       string `json:"ip"`
	Metadata struct {
	} `json:"metadata" validate:"omitempty"`
	Port        int    `json:"port"`
	Scheduled   bool   `json:"scheduled"`
	ServiceName string `json:"serviceName"`
	Weight      int    `json:"weight"`
}

type ServiceBase struct {
	ServiceName string `json:"serviceName" validate:"required"`  // 服务名称
	GroupName   string `json:"groupName" validate:"omitempty"`   // 分组名称
	NamespaceId string `json:"namespaceId" validate:"omitempty"` // 命名空间 ID
}

type ServiceBaseResponse struct {
	Metadata struct {
	} `json:"metadata"`
	GroupName   string `json:"groupName"`
	NamespaceId string `json:"namespaceId"`
	Name        string `json:"name"`
	Selector    struct {
		Type string `json:"type"`
	} `json:"selector"`
	ProtectThreshold int `json:"protectThreshold"`
	Clusters         []struct {
		HealthChecker struct {
			Type string `json:"type"`
		} `json:"healthChecker"`
		Metadata struct {
		} `json:"metadata"`
		Name string `json:"name"`
	} `json:"clusters"`
}

type CreateServiceRequest struct {
	ServiceBase
	ProtectThreshold float64               `json:"protectThreshold" validate:"omitempty,gte=0,lte=1" default:"0"` // 保护阈值,取值0到1,默认0
	Metadata         string                `json:"metadata" validate:"omitempty"`                                 // 元数据
	Selector         CreateServiceSelector `json:"selector" validate:"omitempty"`                                 // 访问策略
}

// CreateServiceSelector 服务访问策略
type CreateServiceSelector struct {
	Default bool `json:"default"`
}

type DeleteServiceRequest struct {
	ServiceBase
}

type ModifyServiceRequest struct {
	ServiceBase
	ProtectThreshold float64               `json:"protectThreshold" validate:"omitempty,gte=0,lte=1" default:"0"` // 保护阈值,取值0到1,默认0
	Metadata         string                `json:"metadata" validate:"omitempty"`                                 // 元数据
	Selector         CreateServiceSelector `json:"selector" validate:"omitempty"`
}

type GetServiceListRequest struct {
	Page
	GroupName   string `json:"groupName" validate:"omitempty"`   // 分组名称
	NamespaceId string `json:"namespaceId" validate:"omitempty"` // 命名空间 ID
}

type GetServiceListResponse struct {
	Count int      `json:"count"`
	Doms  []string `json:"doms"`
}

type GetOperatorSwitchResponse struct {
	Name        string      `json:"name"`
	Masters     interface{} `json:"masters"`
	AdWeightMap struct {
	} `json:"adWeightMap"`
	DefaultPushCacheMillis int     `json:"defaultPushCacheMillis"`
	ClientBeatInterval     int     `json:"clientBeatInterval"`
	DefaultCacheMillis     int     `json:"defaultCacheMillis"`
	DistroThreshold        float64 `json:"distroThreshold"`
	HealthCheckEnabled     bool    `json:"healthCheckEnabled"`
	DistroEnabled          bool    `json:"distroEnabled"`
	EnableStandalone       bool    `json:"enableStandalone"`
	PushEnabled            bool    `json:"pushEnabled"`
	CheckTimes             int     `json:"checkTimes"`
	HttpHealthParams       struct {
		Max    int     `json:"max"`
		Min    int     `json:"min"`
		Factor float64 `json:"factor"`
	} `json:"httpHealthParams"`
	TcpHealthParams struct {
		Max    int     `json:"max"`
		Min    int     `json:"min"`
		Factor float64 `json:"factor"`
	} `json:"tcpHealthParams"`
	MysqlHealthParams struct {
		Max    int     `json:"max"`
		Min    int     `json:"min"`
		Factor float64 `json:"factor"`
	} `json:"mysqlHealthParams"`
	IncrementalList                          []string `json:"incrementalList"`
	ServerStatusSynchronizationPeriodMillis  int      `json:"serverStatusSynchronizationPeriodMillis"`
	ServiceStatusSynchronizationPeriodMillis int      `json:"serviceStatusSynchronizationPeriodMillis"`
	DisableAddIP                             bool     `json:"disableAddIP"`
	SendBeatOnly                             bool     `json:"sendBeatOnly"`
	LimitedUrlMap                            struct {
	} `json:"limitedUrlMap"`
	DistroServerExpiredMillis int      `json:"distroServerExpiredMillis"`
	PushGoVersion             string   `json:"pushGoVersion"`
	PushJavaVersion           string   `json:"pushJavaVersion"`
	PushPythonVersion         string   `json:"pushPythonVersion"`
	PushCVersion              string   `json:"pushCVersion"`
	EnableAuthentication      bool     `json:"enableAuthentication"`
	OverriddenServerStatus    string   `json:"overriddenServerStatus"`
	DefaultInstanceEphemeral  bool     `json:"defaultInstanceEphemeral"`
	HealthCheckWhiteList      []string `json:"healthCheckWhiteList"`
	Checksum                  string   `json:"checksum"`
}

type ModifyOperatorSwitchRequest struct {
	Entry string `json:"entry" validate:"required"`  // 开关名
	Value string `json:"value" validate:"required"`  // 开关值
	Debug bool   `json:"debug" validate:"omitempty"` // 是否只在本机生效,true表示本机生效,false表示集群生效
}

type GetMetricsResponse struct {
	ServiceCount             int     `json:"serviceCount"`
	Load                     float64 `json:"load"`
	Mem                      float64 `json:"mem"`
	ResponsibleServiceCount  int     `json:"responsibleServiceCount"`
	InstanceCount            int     `json:"instanceCount"`
	Cpu                      float64 `json:"cpu"`
	Status                   string  `json:"status"`
	ResponsibleInstanceCount int     `json:"responsibleInstanceCount"`
}

type GetServerListRequest struct {
	Healthy bool `json:"healthy" validate:"omitempty"` // 是否只返回健康Server节点
}

type GetServerListResponse struct {
	Servers []struct {
		Ip             string `json:"ip"`
		ServePort      int    `json:"servePort"`
		Site           string `json:"site"`
		Weight         int    `json:"weight"`
		AdWeight       int    `json:"adWeight"`
		Alive          bool   `json:"alive"`
		LastRefTime    int    `json:"lastRefTime"`
		LastRefTimeStr string `json:"lastRefTimeStr"`
		Key            string `json:"key"`
	} `json:"servers"`
}

type GetLeaderResponse struct {
	Leader struct {
		HeartbeatDueMs int    `json:"heartbeatDueMs"`
		Ip             string `json:"ip"`
		LeaderDueMs    int    `json:"leaderDueMs"`
		State          string `json:"state"`
		Term           int    `json:"term"`
		VoteFor        string `json:"voteFor"`
	} `json:"leader"`
}

type UpdateInstanceHealthStatusRequest struct {
	IP          string `json:"ip" validate:"required"`           // 实例 IP
	Port        int    `json:"port" validate:"required"`         // 实例端口
	NamespaceId string `json:"namespaceId" validate:"omitempty"` // 命名空间 ID
	ClusterName string `json:"clusterName" validate:"omitempty"` // 集群名称
	ServiceName string `json:"serviceName" validate:"required"`  // 服务名称
	GroupName   string `json:"groupName" validate:"omitempty"`   // 分组名称
	Healthy     bool   `json:"healthy" validate:"required"`      // 是否健康
}
