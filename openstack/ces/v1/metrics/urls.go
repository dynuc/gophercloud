package metrics

import "github.com/dynuc/gophercloud"

func getMetricsURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("metrics")
}
