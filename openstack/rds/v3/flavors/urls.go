package flavors

import "github.com/dynuc/gophercloud"

func listURL(sc *gophercloud.ServiceClient, databasename string) string {
	return sc.ServiceURL("flavors", databasename)
}
