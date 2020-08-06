package quotas

import (
	"github.com/dynuc/gophercloud"
)

func getURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("quotas")
}
