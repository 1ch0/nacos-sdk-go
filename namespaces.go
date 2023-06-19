package client

import (
	"net/http"
)

// GetNamespaces 查询命名空间列表
func (c *Client) GetNamespaces() (*GetNamespacesResponse, error) {
	result := &GetNamespacesResponse{}
	return result, c.Execute(
		http.MethodGet,
		&struct{}{},
		IPathNamespaces,
		result,
		map[string]string{})
}

// CreateNamespace 创建命名空间
func (c *Client) CreateNamespace(req *CreateNamespaceRequest) (bool, error) {
	success := &BoolResult
	return *success, c.Execute(
		http.MethodPost,
		req,
		IPathNamespaces,
		success,
		map[string]string{
			PermissionCustomNamespaceId: req.CustomNamespaceId,
			PermissionNamespaceName:     req.NamespaceName,
			PermissionNamespaceDesc:     req.NamespaceDesc,
		})
}

// PutNamespace 修改命名空间
func (c *Client) PutNamespace(req *PutNamespaceRequest) (bool, error) {
	success := &BoolResult
	return *success, c.Execute(
		http.MethodPut,
		req,
		IPathNamespaces,
		&struct{}{},
		map[string]string{
			NameSpace:                   req.Namespace,
			PermissionNamespaceShowName: req.NamespaceShowName,
			PermissionNamespaceDesc:     req.NamespaceDesc,
		})
}

// DeleteNamespace 删除命名空间
func (c *Client) DeleteNamespace(req *DeleteNamespaceRequest) (bool, error) {
	success := &BoolResult
	return *success, c.Execute(
		http.MethodDelete,
		req,
		IPathNamespaces,
		&struct{}{},
		map[string]string{
			PermissionNamespaceId: req.NamespaceId,
		})
}