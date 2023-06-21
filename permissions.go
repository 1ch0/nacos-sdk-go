package client

import (
	"net/http"
	"strconv"
)

// nacos usr management

// GetUsers 查询用户列表 //todo
func (c *Client) GetUsers(req *Page) (*GetUsersResponse, error) {
	result := &GetUsersResponse{}
	return result,
		c.set(http.MethodGet, IPathUser, req).
			do(result,
				map[string]string{
					Search:   SearchType,
					PageNo:   strconv.Itoa(req.PageNo),
					PageSize: strconv.Itoa(req.PageSize),
				})
}

// CreateUser 创建用户
func (c *Client) CreateUser(req *User) error {
	return c.set(http.MethodPost, IPathUser, req).
		do(&struct{}{},
			map[string]string{
				Username: req.Username,
				Password: req.Password,
			})
}

// PutUser 修改用户
func (c *Client) PutUser(req *User) error {
	return c.set(http.MethodPut, IPathUser, req).
		do(&struct{}{},
			map[string]string{
				Username:    req.Username,
				NewPassword: req.Password,
			})
}

// DeleteUser 删除用户
func (c *Client) DeleteUser(req *DeleteUserRequest) error {
	return c.set(http.MethodDelete, IPathUser, req).
		do(&struct{}{},
			map[string]string{
				Username: req.Username,
			})
}

func (c *Client) GetRoles(req *Page) (*GetRolesResponse, error) {
	result := &GetRolesResponse{}
	return result,
		c.set(http.MethodGet, IPathRoles, req).
			do(&struct{}{},
				map[string]string{
					Search:   SearchType,
					PageNo:   strconv.Itoa(req.PageNo),
					PageSize: strconv.Itoa(req.PageSize),
				})
}

func (c *Client) CreateRoles(req *CreateRoleRequest) error {
	return c.set(http.MethodPost, IPathRoles, req).
		do(&struct{}{},
			map[string]string{
				Username: req.Username,
				Role:     req.Role,
			})
}

func (c *Client) DeleteRoles(req *DeleteRoleRequest) error {
	return c.set(http.MethodDelete, IPathRoles, req).
		do(&struct{}{},
			map[string]string{
				Username: req.Username,
				Role:     req.Role,
			})
}

// permissions management

// GetPermissions 查询权限列表
func (c *Client) GetPermissions(req *Page) (*GetPermissions, error) {
	result := &GetPermissions{}
	return result,
		c.set(http.MethodGet, IPathPermission, req).
			do(result,
				map[string]string{
					Search:   SearchType,
					PageNo:   strconv.Itoa(req.PageNo),
					PageSize: strconv.Itoa(req.PageSize),
				})
}

// CreatePermission 创建权限
func (c *Client) CreatePermission(req *CreatePermissionRequest) error {
	return c.set(http.MethodPost, IPathPermission, req).
		do(&struct{}{},
			map[string]string{
				Role:             req.Role,
				Resource:         req.NamespaceId + PermissionSuffix,
				PermissionAction: req.Action,
			})
}

func (c *Client) DeletePermission(req *DeletePermissionRequest) error {
	return c.set(http.MethodDelete, IPathPermission, req).
		do(&struct{}{},
			map[string]string{
				Role:             req.Role,
				Resource:         req.Resource,
				PermissionAction: req.Action,
			})
}
