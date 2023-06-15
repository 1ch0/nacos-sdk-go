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
			NacosAccessToken: c.Authentication.AccessToken,
			NacosSearch:      NacosSearchType,
			PageNo:           strconv.Itoa(req.PageNo),
			PageSize:         strconv.Itoa(req.PageSize),
		}).
		SetResult(result).
		Get(c.Config.Addr + NacosUser)

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
				NacosAccessToken: c.Authentication.AccessToken,
				NacosUsername:    req.Username,
				NacosPassword:    req.Password,
			},
		).
		Post(c.Config.Addr + NacosUser)

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
				NacosAccessToken: c.Authentication.AccessToken,
				NacosUsername:    req.Username,
				NacosNewPassword: req.Password,
			},
		).
		Put(c.Config.Addr + NacosUser)

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
				NacosAccessToken: c.Authentication.AccessToken,
				NacosUsername:    req.Username,
			},
		).
		Delete(c.Config.Addr + NacosUser)

	if err != nil || resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("nacos client delete user failed: %s", resp)
	}
	fmt.Printf("resp: %v\n", resp)
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
			NacosAccessToken: c.Authentication.AccessToken,
			NacosSearch:      NacosSearchType,
			PageNo:           strconv.Itoa(req.PageNo),
			PageSize:         strconv.Itoa(req.PageSize),
		}).
		SetResult(result).
		Get(c.Config.Addr + NacosRoles)

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
				NacosAccessToken: c.Authentication.AccessToken,
				NacosUsername:    req.Username,
				NacosRole:        req.Role,
			},
		).
		Post(c.Config.Addr + NacosRoles)

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
				NacosAccessToken: c.Authentication.AccessToken,
				NacosUsername:    req.Username,
				NacosRole:        req.Role,
			},
		).
		Delete(c.Config.Addr + NacosRoles)

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
			NacosAccessToken: c.Authentication.AccessToken,
			NacosSearch:      NacosSearchType,
			PageNo:           strconv.Itoa(req.PageNo),
			PageSize:         strconv.Itoa(req.PageSize),
		}).
		SetResult(result).
		Get(c.Config.Addr + NacosPermission)

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
				NacosAccessToken: c.Authentication.AccessToken,
				NacosRole:        req.Role,
				NacosResource:    req.NamespaceId + NacosCreatePermissionResource,
				NacosAction:      req.Action,
			},
		).
		Post(c.Config.Addr + NacosPermission)

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
				NacosAccessToken: c.Authentication.AccessToken,
				NacosRole:        req.Role,
				NacosResource:    req.Resource,
				NacosAction:      req.Action,
			},
		).
		Delete(c.Config.Addr + NacosPermission)

	if err != nil || resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("nacos client delete role failed: %s", resp)
	}
	return nil
}
