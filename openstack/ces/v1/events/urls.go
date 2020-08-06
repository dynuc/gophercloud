package events

import (
	"github.com/dynuc/gophercloud"
)

func createURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("events")
}
