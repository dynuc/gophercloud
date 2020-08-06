package publicips

import "github.com/dynuc/gophercloud"

func CreateURL(c *gophercloud.ServiceClient)string{
	return c.ServiceURL("publicips")
}