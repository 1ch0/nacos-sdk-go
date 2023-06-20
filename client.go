package client

import (
	"encoding/json"
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
	body           []byte
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
		body:           []byte{},
	}
}

// Health 检查服务是否健康
func (c *Client) Health() error {
	_, err := c.client.R().SetBasicAuth(c.Config.Username, c.Config.Password).Get(c.Config.Addr)
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

	return nil
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
	data, err := json.Marshal(req)
	if err != nil {
		return err
	}
	c.body = data

	return nil
}

func (c *Client) Execute(method string, req interface{}, path string, result interface{}, queryParams map[string]string) error {
	if !validMethod(method) {
		return fmt.Errorf("invalid method: %s", method)
	}
	err := c.Check(req)
	if err != nil {
		return err
	}

	resp, err := c.client.R().
		SetBody(c.body).
		SetQueryParam(AccessToken, c.Authentication.AccessToken).
		SetQueryParams(queryParams).SetResult(result).Execute(method, c.Config.Addr+path)

	if err != nil || resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("error!  %s %s %s", method, path, resp)
	}

	return nil
}

func (c *Client) Do(params *DoParams) error {
	if !validMethod(params.Method) {
		return fmt.Errorf("invalid method: %s", params.Method)
	}
	err := c.Check(params.Req)
	if err != nil {
		return err
	}

	resp, err := c.client.R().
		SetQueryParam(AccessToken, c.Authentication.AccessToken).
		SetQueryParams(params.QueryParams).SetResult(params.Result).Execute(params.Method, c.Config.Addr+params.Path)

	if err != nil || resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("error!  %s %s %s", params.Method, params.Path, resp)
	}

	return nil
}

func (c *Client) CheckDoParams(params *DoParams) error {
	if !validMethod(params.Method) {
		return fmt.Errorf("invalid method: %s", params.Method)
	}
	if err := c.checkAuth(); err != nil {
		return err
	}
	defaults.SetDefaults(params.Req)
	validate := validator.New()
	if err := validate.Struct(params.Req); err != nil {
		return err
	}

	// TODO check params.Result
	for k, v := range params.QueryParams {
		if v == "" {
			return fmt.Errorf("invalid query param: %s", k)
		}
	}
	return nil
}
