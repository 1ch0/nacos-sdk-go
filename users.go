package client

import (
	"fmt"
	"net/http"
	"strconv"
)

// nacos usr management

// GetUsers 查询用户列表
func (c *Client) GetUsers(req *GetUsersRequest) (*GetUsersResponse, error) {
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
		return nil, fmt.Errorf("nacos client get namespaces failed: %s", resp)
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
		return fmt.Errorf("nacos client create namespace failed: %s", resp)
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
		return fmt.Errorf("nacos client create namespace failed: %s", resp)
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
		return fmt.Errorf("nacos client create namespace failed: %s", resp)
	}
	fmt.Printf("resp: %v\n", resp)
	return nil
}
