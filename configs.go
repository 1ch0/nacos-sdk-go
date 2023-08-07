package client

import (
	"net/http"
	"strconv"
)

// GetConfig 获取配置
func (c *Client) GetConfig(req *ConfigBase) (string, error) {
	result := &StringResult
	return *result,
		c.set(http.MethodGet, IPathConfig, req).
			do(result,
				map[string]string{
					ConfigDataId: req.DataId,
					Tenant:       req.Tenant,
					ConfigGroup:  req.Group,
				})
}

// ListenConfig 监听配置 TODO: 未完成
func (c *Client) ListenConfig(req *ListeningConfigs) (string, error) {
	result := &StringResult
	return *result,
		c.set(http.MethodGet, IPathConfigListener, req).
			do(result,
				map[string]string{
					ConfigDataId: req.DataId,
					Tenant:       req.Tenant,
					ConfigGroup:  req.Group,
				})
}

// PublishConfig 发布配置
func (c *Client) PublishConfig(req *PublishConfigRequest) error {
	result := BoolResult
	return c.set(http.MethodPost, IPathConfig, req).
		do(result,
			map[string]string{
				ConfigDataId:      req.DataId,
				Tenant:            req.Tenant,
				ConfigGroup:       req.Group,
				ConfigContent:     req.Content,
				ConfigContentType: req.ContentType,
			})
}

// DeleteConfig 删除配置
func (c *Client) DeleteConfig(req *ConfigBase) error {
	result := BoolResult
	return c.set(http.MethodDelete, IPathConfig, req).
		do(result,
			map[string]string{
				ConfigDataId: req.DataId,
				Tenant:       req.Tenant,
				ConfigGroup:  req.Group,
			})
}

// GetConfigHistory 配置历史
func (c *Client) GetConfigHistory(req *GetConfigHistoryRequest) (*GetConfigHistoryResponse, error) {
	result := &GetConfigHistoryResponse{}
	return result,
		c.set(http.MethodGet, IPathConfigHistory, req).
			do(result,
				map[string]string{
					ConfigDataId: req.DataId,
					Tenant:       req.Tenant,
					ConfigGroup:  req.Group,
					PageNo:       strconv.Itoa(req.PageNo),
					PageSize:     strconv.Itoa(req.PageSize),
				})
}

// GetConfigHistoryDetail 查询历史版本详情
func (c *Client) GetConfigHistoryDetail(req *GetConfigHistoryDetailRequest) (*GetConfigHistoryDetailResponse, error) {
	result := &GetConfigHistoryDetailResponse{}
	return result,
		c.set(http.MethodGet, IPathConfigHistoryDetail, req).
			do(result,
				map[string]string{
					ConfigDataId: req.DataId,
					Tenant:       req.Tenant,
					ConfigGroup:  req.Group,
					ConfigNid:    req.Nid,
				})
}

// GetConfigHistoryPrevious 查询配置上一版本信息
func (c *Client) GetConfigHistoryPrevious(req *GetConfigHistoryPreviousRequest) (*GetConfigHistoryDetailResponse, error) {
	result := &GetConfigHistoryDetailResponse{}
	return result,
		c.set(http.MethodGet, IPathConfigHistoryPrevious, req).
			do(result,
				map[string]string{
					ConfigDataId: req.DataId,
					Tenant:       req.Tenant,
					ConfigGroup:  req.Group,
					ConfigId:     strconv.Itoa(req.Id),
				})
}
