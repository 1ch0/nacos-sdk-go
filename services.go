package client

import (
	"fmt"
	"net/http"
	"strconv"
)

// service management

// RegisterInstance 注册实例
func (c *Client) RegisterInstance(req *RegisterInstanceRequest) error {
	err := c.Check(req)
	if err != nil {
		return err
	}

	resp, err := c.Resty.R().
		SetQueryParams(
			map[string]string{
				AccessToken:        c.Authentication.AccessToken,
				ServiceIP:          req.Ip,
				ServicePort:        strconv.Itoa(req.Port),
				ServiceNamespaceId: req.NamespaceId,
				ServiceWeight:      strconv.Itoa(req.Weight),
				ServiceEnable:      strconv.FormatBool(req.Enable),
				ServiceHealthy:     strconv.FormatBool(req.Healthy),
				ServiceMetadata:    req.Metadata,
				ServiceClusterName: req.ClusterName,
				ServiceName:        req.ServiceName,
				ServiceGroupName:   req.GroupName,
				ServiceEphemeral:   strconv.FormatBool(req.Ephemeral),
			},
		).
		Post(c.Config.Addr + IPathInstance)

	if err != nil || resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("nacos client register instance failed: %s", resp)
	}
	return nil
}
