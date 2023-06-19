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
