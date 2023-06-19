package client

import (
	"fmt"
	"net/http"
	"strconv"
)

// GetConfig 获取配置
func (c *Client) GetConfig(req *ConfigBase) (string, error) {
	err := c.Check(req)
	if err != nil {
		return "", err
	}

	resp, err := c.client.R().
		SetQueryParams(
			map[string]string{
				AccessToken:  c.Authentication.AccessToken,
				ConfigDataId: req.DataId,
				Tenant:       req.Tenant,
				ConfigGroup:  req.Group,
			},
		).
		Get(c.Config.Addr + IPathConfig)

	if err != nil || resp.StatusCode() != http.StatusOK {
		return "", fmt.Errorf("nacos client get config failed: %s", resp)
	}
	return string(resp.Body()), nil
}

// ListenConfig 监听配置 TODO: 未完成
func (c *Client) ListenConfig(req *ListeningConfigs) (string, error) {
	err := c.Check(req)
	if err != nil {
		return "", err
	}

	// 发送长连接请求
	resp, err := c.client.R().
		SetDoNotParseResponse(true).
		SetHeader("Long-Pulling-Timeout", "30000").
		SetFormData(
			map[string]string{
				AccessToken:  c.Authentication.AccessToken,
				ConfigDataId: req.DataId,
				Tenant:       req.Tenant,
				ConfigGroup:  req.Group,
			},
		).
		Post(c.Config.Addr + IPathConfigListener)
	if err != nil || resp.StatusCode() != http.StatusOK {
		return "", fmt.Errorf("nacos client listen config failed: %s", resp)
	}
	return string(resp.Body()), nil
}

// PublishConfig 发布配置
func (c *Client) PublishConfig(req *PublishConfigRequest) error {
	err := c.Check(req)
	if err != nil {
		return err
	}

	resp, err := c.client.R().
		SetFormData(
			map[string]string{
				AccessToken:       c.Authentication.AccessToken,
				ConfigDataId:      req.DataId,
				Tenant:            req.Tenant,
				ConfigGroup:       req.Group,
				ConfigContent:     req.Content,
				ConfigContentType: req.ContentType,
			},
		).
		Post(c.Config.Addr + IPathConfig)

	if err != nil || resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("nacos client publish config failed: %s", resp)
	}
	return nil
}

// DeleteConfig 删除配置
func (c *Client) DeleteConfig(req *ConfigBase) error {
	err := c.Check(req)
	if err != nil {
		return err
	}
	if req.Group == "" {
		req.Group = DefaultGroup
	}

	resp, err := c.client.R().
		SetQueryParams(
			map[string]string{
				AccessToken:  c.Authentication.AccessToken,
				ConfigDataId: req.DataId,
				Tenant:       req.Tenant,
				ConfigGroup:  req.Group,
			},
		).
		Delete(c.Config.Addr + IPathConfig)

	if err != nil || resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("nacos client delete config failed: %s", resp)
	}
	return nil
}

// GetConfigHistory 配置历史
func (c *Client) GetConfigHistory(req *GetConfigHistoryRequest) (*GetConfigHistoryResponse, error) {
	err := c.Check(req)
	if err != nil {
		return nil, err
	}
	result := &GetConfigHistoryResponse{}
	resp, err := c.client.R().
		SetQueryParams(
			map[string]string{
				AccessToken:  c.Authentication.AccessToken,
				ConfigDataId: req.DataId,
				Tenant:       req.Tenant,
				ConfigGroup:  req.Group,
				PageNo:       strconv.Itoa(req.PageNo),
				PageSize:     strconv.Itoa(req.PageSize),
			},
		).
		SetResult(result).
		Get(c.Config.Addr + IPathConfigHistory)

	if err != nil || resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("nacos client get config history failed: %s", resp)
	}
	return result, nil
}

// GetConfigHistoryDetail 查询历史版本详情
func (c *Client) GetConfigHistoryDetail(req *GetConfigHistoryDetailRequest) (*GetConfigHistoryDetailResponse, error) {
	err := c.Check(req)
	if err != nil {
		return nil, err
	}
	result := &GetConfigHistoryDetailResponse{}
	resp, err := c.client.R().
		SetQueryParams(
			map[string]string{
				AccessToken:  c.Authentication.AccessToken,
				ConfigDataId: req.DataId,
				Tenant:       req.Tenant,
				ConfigGroup:  req.Group,
				ConfigNid:    req.Nid,
			},
		).
		SetResult(result).
		Get(c.Config.Addr + IPathConfigHistoryDetail)

	if err != nil || resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("nacos client get config history detail failed: %s", resp)
	}
	return result, nil
}

// GetConfigHistoryPrevious 查询配置上一版本信息
func (c *Client) GetConfigHistoryPrevious(req *GetConfigHistoryPreviousRequest) (*GetConfigHistoryDetailResponse, error) {
	err := c.Check(req)
	if err != nil {
		return nil, err
	}
	result := &GetConfigHistoryDetailResponse{}
	resp, err := c.client.R().
		SetQueryParams(
			map[string]string{
				AccessToken:  c.Authentication.AccessToken,
				ConfigDataId: req.DataId,
				Tenant:       req.Tenant,
				ConfigGroup:  req.Group,
				ConfigId:     strconv.Itoa(req.Id),
			},
		).
		SetResult(result).
		Get(c.Config.Addr + IPathConfigHistoryPrevious)

	if err != nil || resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("nacos client get config history previous failed: %s", resp)
	}
	return result, nil
}
