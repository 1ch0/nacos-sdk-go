package client

const (
	// NACOS

	NacosScheme     = "http"
	NacosWebContext = "/nacos"

	// NacosDefaultAuthTokenExpireSeconds = 18000

	NacosHttpClientTimeout  = 10
	NacosAccessToken        = "accessToken"
	NacosTokenTtl           = "tokenTtl"
	NacosTokenRefreshWindow = "tokenRefreshWindow"

	// 请求参数
	PageNo   = "pageNo"
	PageSize = "pageSize"

	// NacosNaming

	NacosUsername          = "username"
	NacosPassword          = "password"
	NacosNewPassword       = "newPassword"
	NacosRole              = "role"
	NacosNameSpace         = "namespace"
	NacosTenant            = "tenant"            // tenant	租户信息，对应 Nacos 的命名空间ID字段,命名空间ID字段与命名空间名保持一致
	NacosResource          = "resource"          // 添加权限时资源字段，该值与命名空间名称保持一致
	NacosDataId            = "dataId"            // Naocs 配置 ID
	NacosGroup             = "group"             // Nacos 配置分组
	NacosContent           = "content"           // Nacos 配置内置
	NacosContentType       = "type"              // Nacos 配置类型
	NacosAction            = "action"            // 添加权限，动作字段[r, w, rw]
	NacosCustomNamespaceId = "customNamespaceId" // 命名空间ID,与命名空间名保持一致
	NacosNamespaceName     = "namespaceName"     // 命名空间名
	NacosNamespaceShowName = "namespaceShowName" // 命名空间名
	NacosNamespaceDesc     = "namespaceDesc"     // 命名空间描述，与命名空间名保持一致
	NacosNamespaceId       = "namespaceId"       //命名空间ID,删除命名空间时所需参数
	NacosENV               = "env"               //环境
	NacosServerAddr        = "nacosUrl"          // nacos服务器地址，example: http://xxx.com:80

	// Nacos 默认值

	NacosDefaultAction = "rw"
	NacosDefaultGroup  = "DEFAULT_GROUP" // Nacos 默认 Group
	NacosDefaultTenant = "public"        // Nacos 默认 Tenant
	NacosNid           = "nid"           // Nacos 配置历史版本ID
	NacosId            = "id"            // Nacos 配置历史版本ID

	// NACOS 权限认证

	NacosAuth       = "/nacos/v1/auth/login"         // 鉴权
	NacosNamespaces = "/nacos/v1/console/namespaces" // 命名空间
	NacosUser       = "/nacos/v1/auth/users"         // 用户
	NacosRoles      = "/nacos/v1/auth/roles"         // 角色
	NacosPermission = "/nacos/v1/auth/permissions"   // 权限

	// NACOS 配置管理

	NacosConfig                = "/nacos/v1/cs/configs"                 //获取配置 ,发布配置 ,删除配置
	NacosConfigListener        = "/nacos/v1/cs/configs/listener"        //监听配置
	NacosConfigHistory         = "/nacos/v1/cs/history?search=accurate" //查询历史配置
	NacosConfigHistoryDetail   = "/nacos/v1/cs/history"                 //查询历史版本详情
	NacosConfigHistoryPrevious = "/nacos/v1/cs/history/previous"        //查询配置上一版本信息

	// NACOS 服务发现

	NacosInstance        = "/nacos/v1/ns/instance"          //注册实例,查询实例详情,修改实例,注销实例
	NacosInstanceList    = "/nacos/v1/ns/instance/list"     //查询实例列表
	NacosInstanceBeat    = "/nacos/v1/ns/instance/beat"     //发送实例心跳
	NacosService         = "/nacos/v1/ns/service"           //创建服务,查询服务,修改服务,删除服务
	NacosServicelist     = "/nacos/v1/ns/service/list"      //查询服务列表
	NacosSwitchGet       = "/nacos/v1/ns/operator/switches" //查询系统开关,修改系统开关
	NacosMetrics         = "/nacos/v1/ns/operator/metrics"  //查看系统当前数据指标
	NacosOperatorServers = "/nacos/v1/ns/operator/servers"  //查看集群server列表
	NacosLeader          = "/nacos/v1/ns/raft/leader"       //查看集群leader
	NacosHealth          = "/nacos/v1/ns/health/instance"   //更新实例健康状态
)
