package client

import (
	"net/http"
	"strconv"
)

// nacos usr management

// GetUsers 查询用户列表 //todo
func (c *Client) GetUsers(req *Page) (*GetUsersResponse, error) {
	result := &GetUsersResponse{}
	return result, c.Execute(
		http.MethodGet,
		req,
		IPathUser,
		result,
		map[string]string{
			Search:   SearchType,
			PageNo:   strconv.Itoa(req.PageNo),
			PageSize: strconv.Itoa(req.PageSize),
		})
}

// CreateUser 创建用户
func (c *Client) CreateUser(req *User) error {
	return c.Execute(
		http.MethodPost,
		req,
		IPathUser,
		&struct{}{},
		map[string]string{
			Username: req.Username,
			Password: req.Password,
		})
}

// PutUser 修改用户
func (c *Client) PutUser(req *User) error {
	return c.Execute(
		http.MethodPut,
		req,
		IPathUser,
		&struct{}{},
		map[string]string{
			Username: req.Username,
			Password: req.Password,
		})
}

// DeleteUser 删除用户
func (c *Client) DeleteUser(req *DeleteUserRequest) error {
	return c.Execute(
		http.MethodDelete,
		req,
		IPathUser,
		&struct{}{},
		map[string]string{
			Username: req.Username,
		})
}

func (c *Client) GetRoles(req *Page) (*GetRolesResponse, error) {
	result := &GetRolesResponse{}
	return result, c.Execute(
		http.MethodGet,
		req,
		IPathRoles,
		result,
		map[string]string{
			Search:   SearchType,
			PageNo:   strconv.Itoa(req.PageNo),
			PageSize: strconv.Itoa(req.PageSize),
		})
}

func (c *Client) CreateRoles(req *CreateRoleRequest) error {
	return c.Execute(
		http.MethodPost,
		req,
		IPathRoles,
		&struct{}{},
		map[string]string{
			Username: req.Username,
			Role:     req.Role,
		})
}

func (c *Client) DeleteRoles(req *DeleteRoleRequest) error {
	return c.Execute(
		http.MethodDelete,
		req,
		IPathRoles,
		&struct{}{},
		map[string]string{
			Username: req.Username,
			Role:     req.Role,
		})
}

// permissions management

// GetPermissions 查询权限列表
func (c *Client) GetPermissions(req *Page) (*GetPermissions, error) {
	result := &GetPermissions{}
	return result, c.Execute(
		http.MethodGet,
		req,
		IPathPermission,
		result,
		map[string]string{
			Search:   SearchType,
			PageNo:   strconv.Itoa(req.PageNo),
			PageSize: strconv.Itoa(req.PageSize),
		})
}

// CreatePermission 创建权限
func (c *Client) CreatePermission(req *CreatePermissionRequest) error {
	return c.Execute(
		http.MethodPost,
		req,
		IPathPermission,
		&struct{}{},
		map[string]string{
			Role:             req.Role,
			Resource:         req.NamespaceId + PermissionSuffix,
			PermissionAction: req.Action,
		})
}

func (c *Client) DeletePermission(req *DeletePermissionRequest) error {
	return c.Execute(
		http.MethodDelete,
		req,
		IPathPermission,
		&struct{}{},
		map[string]string{
			Role:             req.Role,
			Resource:         req.Resource,
			PermissionAction: req.Action,
		})
}
