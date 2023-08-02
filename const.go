package client

const (
	// NACOS

	Scheme     = "client"
	WebContext = "/"

	// NacosDefaultAuthTokenExpireSeconds = 18000

	HttpClientTimeout  = 10
	AccessToken        = "accessToken"
	TokenTtl           = "tokenTtl"
	TokenRefreshWindow = "tokenRefreshWindow"

	// Request Params

	PageNo   = "pageNo"
	PageSize = "pageSize"

	// Nacos services management request params

	ServiceIP               = "ip"               // 服务IP, string
	ServicePort             = "port"             // 服务端口, int
	ServiceNamespaceId      = "namespaceId"      // 命名空间ID, string
	ServiceWeight           = "weight"           // 权重, double
	ServiceEnable           = "enable"           // 是否上线, bool
	ServiceHealthy          = "healthy"          // 是否只健康, bool
	ServiceMetadata         = "metadata"         // 元数据, string
	ServiceClusterName      = "clusterName"      // 集群名称, string
	ServiceName             = "serviceName"      // 服务名称, string
	ServiceGroupName        = "groupName"        // 分组名称, string
	ServiceEphemeral        = "ephemeral"        // 是否临时实例, bool
	ServiceClusters         = "clusters"         // 集群名称, string,多个集群用逗号分隔
	ServiceProtectThreshold = "protectThreshold" // 保护阈值, float64
	ServiceSelector         = "selector"         // 选择器, 服务访问策略

	OperatorEntry = "entry" // 开关名
	OperatorValue = "value" // 开关值
	OperatorDebug = "debug" // 开关值

	// NacosNaming

	PermissionSuffix = ":*:*"
	Search           = "search"
	SearchType       = "accurate"
	Username         = "username"
	Password         = "password"
	NewPassword      = "newPassword"
	Role             = "role"
	NameSpace        = "namespace"
	Tenant           = "tenant"   // tenant	租户信息，对应 Nacos 的命名空间ID字段,命名空间ID字段与命名空间名保持一致
	Resource         = "resource" // 添加权限时资源字段，该值与命名空间名称保持一致

	ConfigDataId      = "dataId"  // Nacos 配置 ID
	ConfigGroup       = "group"   // Nacos 配置分组
	ConfigContent     = "content" // Nacos 配置内置
	ConfigContentType = "type"    // Nacos 配置类型
	ConfigNid         = "nid"     // Nacos 配置历史版本ID
	ConfigId          = "id"      // Nacos 配置历史版本ID

	PermissionAction            = "action"            // 添加权限，动作字段[r, w, rw]
	PermissionCustomNamespaceId = "customNamespaceId" // 命名空间ID,与命名空间名保持一致
	PermissionNamespaceName     = "namespaceName"     // 命名空间名
	PermissionNamespaceShowName = "namespaceShowName" // 命名空间名
	PermissionNamespaceDesc     = "namespaceDesc"     // 命名空间描述，与命名空间名保持一致
	PermissionNamespaceId       = "namespaceId"       //命名空间ID,删除命名空间时所需参数
	Env                         = "env"               //环境
	ServerAddr                  = "nacosUrl"          // nacos服务器地址，example: http://xxx.com:80

	// Nacos 默认值

	DefaultGroup  = "DEFAULT_GROUP" // Nacos 默认 ConfigGroup
	DefaultTenant = "public"        // Nacos 默认 Tenant

	// NACOS 接口请求地址

	IPathAuth                  = "/nacos/v1/auth/login"                 // 鉴权
	IPathNamespaces            = "/nacos/v1/console/namespaces"         // 命名空间
	IPathUser                  = "/nacos/v1/auth/users"                 // 用户
	IPathRoles                 = "/nacos/v1/auth/roles"                 // 角色
	IPathPermission            = "/nacos/v1/auth/permissions"           // 权限
	IPathConfig                = "/nacos/v1/cs/configs"                 //获取配置 ,发布配置 ,删除配置
	IPathConfigListener        = "/nacos/v1/cs/configs/listener"        //监听配置
	IPathConfigHistory         = "/nacos/v1/cs/history?search=accurate" //查询历史配置
	IPathConfigHistoryDetail   = "/nacos/v1/cs/history"                 //查询历史版本详情
	IPathConfigHistoryPrevious = "/nacos/v1/cs/history/previous"        //查询配置上一版本信息

	// NACOS 服务发现

	IPathInstance        = "/nacos/v1/ns/instance"          //注册实例,查询实例详情,修改实例,注销实例
	IPathInstanceList    = "/nacos/v1/ns/instance/list"     //查询实例列表
	IPathInstanceBeat    = "/nacos/v1/ns/instance/beat"     //发送实例心跳
	IPathService         = "/nacos/v1/ns/service"           //创建服务,查询服务,修改服务,删除服务
	IPathServiceList     = "/nacos/v1/ns/service/list"      //查询服务列表
	IPathOperatorSwitch  = "/nacos/v1/ns/operator/switches" //查询系统开关,修改系统开关
	IPathOperatorMetrics = "/nacos/v1/ns/operator/metrics"  //查看系统当前数据指标
	IPathOperatorServers = "/nacos/v1/ns/operator/servers"  //查看集群server列表
	IPathOperatorLeader  = "/nacos/v1/ns/raft/leader"       //查看集群leader
	IPathInstanceHealth  = "/nacos/v1/ns/health/instance"   //更新实例健康状态
)
