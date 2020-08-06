package cloudservers

import "github.com/dynuc/gophercloud"

func createURL(sc *gophercloud.ServiceClient) string {
	return sc.ServiceURL("cloudservers")
}

func jobURL(sc *gophercloud.ServiceClient, jobId string) string {
	return sc.ServiceURL("jobs", jobId)
}
