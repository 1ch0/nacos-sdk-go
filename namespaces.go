package client

import (
	"fmt"
	"net/http"
)

// GetNamespaces 查询命名空间列表
func (c *Client) GetNamespaces() (*GetNamespacesResponse, error) {
	err := c.Check(&struct{}{})
	if err != nil {
		return nil, err
	}
	result := &GetNamespacesResponse{}
	resp, err := c.Resty.R().
		SetQueryParam(NacosAccessToken, c.Authentication.AccessToken).
		SetResult(result).
		Get(c.Config.Addr + NacosNamespaces)

	if err != nil || resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("nacos client get namespaces failed: %s", resp)
	}
	return result, nil
}

// CreateNamespace 创建命名空间
func (c *Client) CreateNamespace(req *CreateNamespaceRequest) (bool, error) {
	err := c.Check(req)
	if err != nil {
		return false, err
	}

	resp, err := c.Resty.R().
		SetQueryParams(
			map[string]string{
				NacosAccessToken:       c.Authentication.AccessToken,
				NacosCustomNamespaceId: req.CustomNamespaceId,
				NacosNamespaceName:     req.NamespaceName,
				NacosNamespaceDesc:     req.NamespaceDesc,
			},
		).
		Post(c.Config.Addr + NacosNamespaces)

	if err != nil || resp.StatusCode() != http.StatusOK {
		return false, fmt.Errorf("nacos client create namespace failed: %s", resp)
	}
	return true, nil
}

// PutNamespace 修改命名空间
func (c *Client) PutNamespace(req *PutNamespaceRequest) (bool, error) {
	err := c.Check(req)
	if err != nil {
		return false, err
	}

	resp, err := c.Resty.R().
		SetQueryParams(
			map[string]string{
				NacosAccessToken:       c.Authentication.AccessToken,
				NacosNameSpace:         req.Namespace,
				NacosNamespaceShowName: req.NamespaceShowName,
				NacosNamespaceDesc:     req.NamespaceDesc,
			},
		).
		Put(c.Config.Addr + NacosNamespaces)

	if err != nil || resp.StatusCode() != http.StatusOK {
		return false, fmt.Errorf("nacos client put namespace failed: %s", resp)
	}
	return true, nil
}

// DeleteNamespace 删除命名空间
func (c *Client) DeleteNamespace(req *DeleteNamespaceRequest) (bool, error) {
	err := c.Check(req)
	if err != nil {
		return false, err
	}

	resp, err := c.Resty.R().
		SetQueryParams(
			map[string]string{
				NacosAccessToken: c.Authentication.AccessToken,
				NacosNamespaceId: req.NamespaceId,
			},
		).
		Delete(c.Config.Addr + NacosNamespaces)

	if err != nil || resp.StatusCode() != http.StatusOK {
		return false, fmt.Errorf("nacos client delete namespace failed: %s", resp)
	}
	return true, nil
}
