package client

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// nacos usr management

// GetUsers 查询用户列表
func (c *Client) GetUsers(req *Page) (*GetUsersResponse, error) {
	err := c.Check(req)
	if err != nil {
		return nil, err
	}
	result := &GetUsersResponse{}
	resp, err := c.Resty.R().
		SetQueryParams(map[string]string{
			AccessToken: c.Authentication.AccessToken,
			Search:      SearchType,
			PageNo:      strconv.Itoa(req.PageNo),
			PageSize:    strconv.Itoa(req.PageSize),
		}).
		SetResult(result).
		Get(c.Config.Addr + IPathUser)

	if err != nil || resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("nacos client get users failed: %s", resp)
	}
	return result, nil
}

// CreateUser 创建用户
func (c *Client) CreateUser(req *User) error {
	err := c.Check(req)
	if err != nil {
		return err
	}

	resp, err := c.Resty.R().
		SetQueryParams(
			map[string]string{
				AccessToken: c.Authentication.AccessToken,
				Username:    req.Username,
				Password:    req.Password,
			},
		).
		Post(c.Config.Addr + IPathUser)

	if err != nil || resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("nacos client create user failed: %s", resp)
	}
	return nil
}

// PutUser 修改用户
func (c *Client) PutUser(req *User) error {
	err := c.Check(req)
	if err != nil {
		return err
	}

	resp, err := c.Resty.R().
		SetQueryParams(
			map[string]string{
				AccessToken: c.Authentication.AccessToken,
				Username:    req.Username,
				NewPassword: req.Password,
			},
		).
		Put(c.Config.Addr + IPathUser)

	if err != nil || resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("nacos client put user failed: %s", resp)
	}
	return nil
}

// DeleteUser 删除用户
func (c *Client) DeleteUser(req *DeleteUserRequest) error {
	err := c.Check(req)
	if err != nil {
		return err
	}

	resp, err := c.Resty.R().
		SetQueryParams(
			map[string]string{
				AccessToken: c.Authentication.AccessToken,
				Username:    req.Username,
			},
		).
		Delete(c.Config.Addr + IPathUser)

	if err != nil || resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("nacos client delete user failed: %s", resp)
	}
	return nil
}

func (c *Client) GetRoles(req *Page) (*GetRolesResponse, error) {
	err := c.Check(req)
	if err != nil {
		return nil, err
	}
	result := &GetRolesResponse{}
	resp, err := c.Resty.R().
		SetQueryParams(map[string]string{
			AccessToken: c.Authentication.AccessToken,
			Search:      SearchType,
			PageNo:      strconv.Itoa(req.PageNo),
			PageSize:    strconv.Itoa(req.PageSize),
		}).
		SetResult(result).
		Get(c.Config.Addr + IPathRoles)

	if err != nil || resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("nacos client get role list failed: %s", resp)
	}
	return result, nil
}

func (c *Client) CreateRoles(req *CreateRoleRequest) error {
	err := c.Check(req)
	if err != nil {
		return err
	}

	resp, err := c.Resty.R().
		SetQueryParams(
			map[string]string{
				AccessToken: c.Authentication.AccessToken,
				Username:    req.Username,
				Role:        req.Role,
			},
		).
		Post(c.Config.Addr + IPathRoles)

	if err != nil || resp.StatusCode() != http.StatusOK {
		if strings.Contains(resp.String(), "Duplicate entry") {
			return nil
		}
		return fmt.Errorf("nacos client create role failed: %s", resp)
	}
	return nil
}

func (c *Client) DeleteRoles(req *DeleteRoleRequest) error {
	err := c.Check(req)
	if err != nil {
		return err
	}

	resp, err := c.Resty.R().
		SetQueryParams(
			map[string]string{
				AccessToken: c.Authentication.AccessToken,
				Username:    req.Username,
				Role:        req.Role,
			},
		).
		Delete(c.Config.Addr + IPathRoles)

	if err != nil || resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("nacos client delete role failed: %s", resp)
	}
	return nil
}

// permissions management

// GetPermissions 查询权限列表
func (c *Client) GetPermissions(req *Page) (*GetPermissions, error) {
	err := c.Check(req)
	if err != nil {
		return nil, err
	}
	result := &GetPermissions{}
	resp, err := c.Resty.R().
		SetQueryParams(map[string]string{
			AccessToken: c.Authentication.AccessToken,
			Search:      SearchType,
			PageNo:      strconv.Itoa(req.PageNo),
			PageSize:    strconv.Itoa(req.PageSize),
		}).
		SetResult(result).
		Get(c.Config.Addr + IPathPermission)

	if err != nil || resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("nacos client get users failed: %s", resp)
	}
	return result, nil
}

// CreatePermission 创建权限
func (c *Client) CreatePermission(req *CreatePermissionRequest) error {
	err := c.Check(req)
	if err != nil {
		return err
	}

	resp, err := c.Resty.R().
		SetQueryParams(
			map[string]string{
				AccessToken:      c.Authentication.AccessToken,
				Role:             req.Role,
				Resource:         req.NamespaceId + PermissionSuffix,
				PermissionAction: req.Action,
			},
		).
		Post(c.Config.Addr + IPathPermission)

	if err != nil || resp.StatusCode() != http.StatusOK {
		if strings.Contains(resp.String(), "Duplicate entry") {
			return nil
		}
		return fmt.Errorf("nacos client create role failed: %s", resp)
	}
	return nil
}

func (c *Client) DeletePermission(req *DeletePermissionRequest) error {
	err := c.Check(req)
	if err != nil {
		return err
	}

	resp, err := c.Resty.R().
		SetQueryParams(
			map[string]string{
				AccessToken:      c.Authentication.AccessToken,
				Role:             req.Role,
				Resource:         req.Resource,
				PermissionAction: req.Action,
			},
		).
		Delete(c.Config.Addr + IPathPermission)

	if err != nil || resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("nacos client delete role failed: %s", resp)
	}
	return nil
}
