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
	authentication *authentication
	client         *resty.Client
	httpMethod     string
	iPath          string
	error          error
}

type Config struct {
	Version  string `yaml:"version"`
	Addr     string `yaml:"addr"`
	Scheme   string `yaml:"scheme"`
	IP       string `yaml:"ip"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type authentication struct {
	AccessToken string `json:"accessToken"`
	TokenTtl    int    `json:"tokenTtl"`
	GlobalAdmin bool   `json:"globalAdmin"`
}

func New(config *Config) *Client {
	return &Client{
		Config:         config.set(),
		authentication: &authentication{},
		client:         resty.New().SetTimeout(5 * time.Second).SetDisableWarn(true).SetRetryCount(3),
		httpMethod:     "",
		iPath:          "",
		error:          nil}
}

func (c *Config) set() *Config {
	if c.Addr == "" {
		if c.Scheme != "" {
			c.Addr = fmt.Sprintf("%s://%s:%s", c.Scheme, c.IP, c.Port)
		}
		c.Addr = fmt.Sprintf("http://%s:%s", c.IP, c.Port)
	}
	return c
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
	c.authentication = &authentication{
		response.AccessToken,
		response.TokenTtl,
		response.GlobalAdmin,
	}

	go c.resetToken()

	return nil
}

func (c *Client) resetToken() {
	if c.authentication.TokenTtl <= 10 {
		return
	}
	duration := time.Second * time.Duration(c.authentication.TokenTtl-10)
	timer := time.AfterFunc(duration, func() {
		c.authentication.AccessToken = ""
	})

	// Wait for the timer to expire
	<-timer.C
}

func (c *Client) checkAuth() error {
	if c.authentication.AccessToken == "" {
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

func (c *Client) set(method string, path string, req interface{}) *Client {
	c.iPath = path
	if !validMethod(method) {
		c.error = fmt.Errorf("request %s%s invalid http request method: %s", c.Config.Addr, c.iPath, method)
	}
	c.httpMethod = method
	err := c.checkReq(req)
	if err != nil {
		c.error = fmt.Errorf("request %s%s invalid params: %s", c.Config.Addr, c.iPath, err)
	}
	return c
}

func (c *Client) clear() *Client {
	c.httpMethod = ""
	c.iPath = ""
	c.error = nil
	return c
}

func (c *Client) do(result interface{}, queryParams map[string]string) error {
	defer c.clear()
	if c.error != nil {
		return c.error
	}
	resp, err := c.client.R().
		SetQueryParam(AccessToken, c.authentication.AccessToken).
		SetQueryParams(queryParams).SetResult(result).Execute(c.httpMethod, c.Config.Addr+c.iPath)
	if err != nil || resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("error!  %s %s %s", c.httpMethod, c.iPath, resp)
	}
	return nil
}
