package client

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mcuadros/go-defaults"

	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	Config         *Config
	Authentication *Authentication
	client         *resty.Client
	HttpMethod     string
	IPath          string
	Error          error
}

type Config struct {
	Addr     string
	Username string
	Password string
}

type Authentication struct {
	AccessToken string `json:"accessToken"`
	TokenTtl    int    `json:"tokenTtl"`
	GlobalAdmin bool   `json:"globalAdmin"`
}

func New(config *Config) *Client {
	return &Client{
		Config:         config,
		Authentication: &Authentication{},
		client:         resty.New().SetTimeout(5 * time.Second).SetDisableWarn(true).SetRetryCount(3),
		HttpMethod:     "",
		IPath:          "",
		Error:          nil}
}

// Health 检查服务是否健康
func (c *Client) Health() error {
	_, err := c.client.R().Get(c.Config.Addr)
	if err != nil {
		return err
	}
	return nil
}

// Login 登录
func (c *Client) Login() error {
	response := &LoginResponse{}
	resp, _ := c.client.R().
		SetFormData(map[string]string{
			"username": c.Config.Username,
			"password": c.Config.Password,
		}).
		SetResult(response).
		Post(c.Config.Addr + IPathAuth)
	if resp.StatusCode() != http.StatusOK || response == nil {
		return fmt.Errorf("nacos auth failed: #%s", resp.Body())
	}
	c.Authentication = &Authentication{
		response.AccessToken,
		response.TokenTtl,
		response.GlobalAdmin,
	}

	go c.resetToken()

	return nil
}

func (c *Client) resetToken() {
	if c.Authentication.TokenTtl <= 10 {
		return
	}
	duration := time.Second * time.Duration(c.Authentication.TokenTtl-10)
	timer := time.AfterFunc(duration, func() {
		c.Authentication.AccessToken = ""
	})

	// Wait for the timer to expire
	<-timer.C
}

func (c *Client) checkAuth() error {
	if c.Authentication.AccessToken == "" {
		err := c.Login()
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) checkReq(req interface{}) error {
	err := c.checkAuth()
	if err != nil {
		return err
	}
	defaults.SetDefaults(req)
	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) check(method string, path string, req interface{}) *Client {
	c.IPath = path
	if !validMethod(method) {
		c.Error = fmt.Errorf("request %s%s invalid http request method: %s", c.Config.Addr, c.IPath, method)
	}
	c.HttpMethod = method
	err := c.checkReq(req)
	if err != nil {
		c.Error = fmt.Errorf("request %s%s invalid params: %s", c.Config.Addr, c.IPath, err)
	}
	return c
}

func (c *Client) clear() *Client {
	c.HttpMethod = ""
	c.IPath = ""
	c.Error = nil
	return c
}

func (c *Client) do(result interface{}, queryParams map[string]string) error {
	defer c.clear()
	if c.Error != nil {
		return c.Error
	}
	resp, err := c.client.R().
		SetQueryParam(AccessToken, c.Authentication.AccessToken).
		SetQueryParams(queryParams).SetResult(result).Execute(c.HttpMethod, c.Config.Addr+c.IPath)
	if err != nil || resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("error!  %s %s %s", c.HttpMethod, c.IPath, resp)
	}
	return nil
}
