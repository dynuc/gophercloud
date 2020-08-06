package datastores

import "github.com/dynuc/gophercloud"

func listURL(sc *gophercloud.ServiceClient, databasename string) string {
	return sc.ServiceURL("datastores", databasename)
}
