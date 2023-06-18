package client

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/mcuadros/go-defaults"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	Config         *Config
	Resty          *resty.Client
	Authentication *Authentication
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
		Config: config,
		Resty:  resty.New().SetTimeout(5 * time.Second).SetDisableWarn(true).SetRetryCount(3),
	}

}

// Health 检查服务是否健康
func (c *Client) Health() error {
	_, err := c.Resty.R().SetBasicAuth(c.Config.Username, c.Config.Password).Get(c.Config.Addr)
	if err != nil {
		return err
	}
	return nil
}

// Login 登录
func (c *Client) Login() error {
	response := &LoginResponse{}
	resp, _ := c.Resty.R().
		SetFormData(map[string]string{
			"username": c.Config.Username,
			"password": c.Config.Password,
		}).
		SetResult(response).
		Post(c.Config.Addr + IPathAuth)
	if resp.StatusCode() != http.StatusOK || response == nil {
		return fmt.Errorf("nacos auth failed: #%s", resp.Body())
	}
	c.Authentication = &Authentication{}
	c.Authentication.AccessToken = response.AccessToken
	c.Authentication.TokenTtl = response.TokenTtl
	c.Authentication.GlobalAdmin = response.GlobalAdmin
	return nil
}

func (c *Client) checkAuth() error {
	if c.Authentication == nil {
		err := c.Login()
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) Check(req interface{}) error {
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
