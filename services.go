package client

import (
	"fmt"
	"net/http"
	"strconv"
)

// service management

// RegisterInstance 注册实例
func (c *Client) RegisterInstance(req *RegisterInstanceRequest) error {
	return c.Execute(
		http.MethodPost,
		req,
		IPathInstance,
		&struct{}{},
		map[string]string{
			ServiceIP:          req.IP,
			ServicePort:        strconv.Itoa(req.Port),
			ServiceNamespaceId: req.NamespaceId,
			ServiceWeight:      fmt.Sprintf("%f", req.Weight),
			ServiceEnable:      strconv.FormatBool(req.Enable),
			ServiceHealthy:     strconv.FormatBool(req.Healthy),
			ServiceMetadata:    req.Metadata,
			ServiceClusterName: req.ClusterName,
			ServiceName:        req.ServiceName,
			ServiceGroupName:   req.GroupName,
			ServiceEphemeral:   strconv.FormatBool(req.Ephemeral),
		})
}

// DeregisterInstance 注销实例
func (c *Client) DeregisterInstance(req *DeregisterInstanceRequest) error {
	return c.Execute(
		http.MethodDelete,
		req,
		IPathInstance,
		&struct{}{},
		map[string]string{
			ServiceIP:          req.IP,
			ServicePort:        strconv.Itoa(req.Port),
			ServiceNamespaceId: req.NamespaceId,
			ServiceClusterName: req.ClusterName,
			ServiceName:        req.ServiceName,
			ServiceGroupName:   req.GroupName,
			ServiceEphemeral:   strconv.FormatBool(req.Ephemeral),
		})
}

// ModifyInstance 修改实例
func (c *Client) ModifyInstance(req *ModifyInstanceRequest) error {
	return c.Execute(
		http.MethodPut,
		req,
		IPathInstance,
		&struct{}{},
		map[string]string{
			ServiceIP:          req.IP,
			ServicePort:        strconv.Itoa(req.Port),
			ServiceNamespaceId: req.NamespaceId,
			ServiceWeight:      fmt.Sprintf("%f", req.Weight),
			ServiceEnable:      strconv.FormatBool(req.Enable),
			ServiceHealthy:     strconv.FormatBool(req.Healthy),
			ServiceMetadata:    req.Metadata,
			ServiceClusterName: req.ClusterName,
			ServiceName:        req.ServiceName,
			ServiceGroupName:   req.GroupName,
			ServiceEphemeral:   strconv.FormatBool(req.Ephemeral),
		})
}

// GetInstances 查询实例列表
func (c *Client) GetInstances(req *GetInstancesRequest) (*GetInstancesResponse, error) {
	result := &GetInstancesResponse{}
	return result, c.Execute(
		http.MethodGet,
		req,
		IPathInstanceList,
		result,
		map[string]string{
			ServiceNamespaceId: req.NamespaceId,
			ServiceName:        req.ServiceName,
			ServiceGroupName:   req.GroupName,
			ServiceClusters:    req.Clusters,
		})
}

// GetInstance 查询实例详情
func (c *Client) GetInstance(req *GetInstanceRequest) (*GetInstanceResponse, error) {
	result := &GetInstanceResponse{}
	return result, c.Execute(
		http.MethodGet,
		req,
		IPathInstance,
		result,
		map[string]string{
			ServiceIP:          req.IP,
			ServicePort:        strconv.Itoa(req.Port),
			ServiceNamespaceId: req.NamespaceId,
			ServiceClusterName: req.ClusterName,
			ServiceName:        req.ServiceName,
			ServiceGroupName:   req.GroupName,
			ServiceHealthy:     strconv.FormatBool(req.Healthy),
			ServiceEphemeral:   strconv.FormatBool(req.Ephemeral),
		})
}

// SendHeartbeat 发送心跳
func (c *Client) SendHeartbeat(req *SendHeartbeatRequest) (bool, error) {
	result := &BoolResult
	return *result, c.Execute(
		http.MethodPut,
		req,
		IPathInstanceBeat,
		result,
		map[string]string{
			ServiceName:        req.ServiceName,
			ServiceIP:          req.IP,
			ServicePort:        strconv.Itoa(req.Port),
			ServiceNamespaceId: req.NamespaceId,
			ServiceClusterName: req.ClusterName,
			ServiceGroupName:   req.GroupName,
			ServiceEphemeral:   strconv.FormatBool(req.Ephemeral),
		})
}

// CreateService 创建服务 //todo
func (c *Client) CreateService(req *CreateServiceRequest) (bool, error) {
	result := &BoolResult
	return *result, c.Execute(
		http.MethodPost,
		req,
		IPathService,
		result,
		map[string]string{
			ServiceName:             req.ServiceName,
			ServiceGroupName:        req.GroupName,
			ServiceNamespaceId:      req.NamespaceId,
			ServiceProtectThreshold: fmt.Sprintf("%f", req.ProtectThreshold),
		})
}

// DeleteService 删除服务
func (c *Client) DeleteService(req *DeleteServiceRequest) (bool, error) {
	result := &BoolResult
	return *result, c.Execute(
		http.MethodDelete,
		req,
		IPathService,
		result,
		map[string]string{
			ServiceName:        req.ServiceName,
			ServiceGroupName:   req.GroupName,
			ServiceNamespaceId: req.NamespaceId,
		})
}

// ModifyService 修改服务 //todo
func (c *Client) ModifyService(req *ModifyServiceRequest) (bool, error) {
	result := &BoolResult
	return *result, c.Execute(
		http.MethodPut,
		req,
		IPathService,
		result,
		map[string]string{
			ServiceName:             req.ServiceName,
			ServiceGroupName:        req.GroupName,
			ServiceNamespaceId:      req.NamespaceId,
			ServiceProtectThreshold: fmt.Sprintf("%f", req.ProtectThreshold),
		})
}

// GetService 查询服务详情
func (c *Client) GetService(req *ServiceBase) (*ServiceBaseResponse, error) {
	result := &ServiceBaseResponse{}
	return result, c.Execute(
		http.MethodGet,
		req,
		IPathService,
		result,
		map[string]string{
			ServiceName:        req.ServiceName,
			ServiceGroupName:   req.GroupName,
			ServiceNamespaceId: req.NamespaceId,
		})
}

// GetServiceList 查询服务列表
func (c *Client) GetServiceList(req *GetServiceListRequest) (*GetServiceListResponse, error) {
	result := &GetServiceListResponse{}
	return result, c.Execute(
		http.MethodGet,
		req,
		IPathServiceList,
		result,
		map[string]string{
			PageNo:             strconv.Itoa(req.PageNo),
			PageSize:           strconv.Itoa(req.PageSize),
			ServiceNamespaceId: req.NamespaceId,
			ServiceGroupName:   req.GroupName,
		})
}

// GetOperatorSwitch 查询系统开关
func (c *Client) GetOperatorSwitch() (*GetOperatorSwitchResponse, error) {
	result := &GetOperatorSwitchResponse{}
	return result, c.Execute(
		http.MethodGet,
		&struct{}{},
		IPathOperatorSwitch,
		result,
		map[string]string{})
}

// ModifyOperatorSwitch 修改系统开关
func (c *Client) ModifyOperatorSwitch(req *ModifyOperatorSwitchRequest) (bool, error) {
	result := &BoolResult
	return *result, c.Execute(
		http.MethodPut,
		req,
		IPathOperatorSwitch,
		result,
		map[string]string{
			OperatorEntry: req.Entry,
			OperatorValue: req.Value,
			OperatorDebug: strconv.FormatBool(req.Debug),
		})
}

// GetOperatorMetrics 查看系统当前数据指标
func (c *Client) GetOperatorMetrics() (*GetMetricsResponse, error) {
	result := &GetMetricsResponse{}
	return result, c.Execute(
		http.MethodGet,
		&struct{}{},
		IPathOperatorMetrics,
		result,
		map[string]string{})
}

// GetOperatorServerList 查看当前集群Server列表
func (c *Client) GetOperatorServerList(req *GetServerListRequest) (*GetServerListResponse, error) {
	result := &GetServerListResponse{}
	return result, c.Execute(
		http.MethodGet,
		&struct{}{},
		IPathOperatorServers,
		result,
		map[string]string{
			ServiceHealthy: strconv.FormatBool(req.Healthy),
		})
}

// GetOperatorLeader 查看当前集群leader
func (c *Client) GetOperatorLeader() (*GetLeaderResponse, error) {
	result := &GetLeaderResponse{}
	return result, c.Execute(
		http.MethodGet,
		&struct{}{},
		IPathOperatorLeader,
		result,
		map[string]string{})
}

// UpdateInstanceHealthStatus 更新实例的健康状态
func (c *Client) UpdateInstanceHealthStatus(req *UpdateInstanceHealthStatusRequest) (bool, error) {
	result := &BoolResult
	return *result, c.Execute(
		http.MethodPut,
		req,
		IPathInstanceHealth,
		result,
		map[string]string{
			ServiceName:        req.ServiceName,
			ServiceIP:          req.IP,
			ServicePort:        strconv.Itoa(req.Port),
			ServiceNamespaceId: req.NamespaceId,
			ServiceClusterName: req.ClusterName,
			ServiceGroupName:   req.GroupName,
			ServiceHealthy:     strconv.FormatBool(req.Healthy),
		})
}
