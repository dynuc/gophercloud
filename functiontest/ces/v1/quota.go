package main

import (
	"encoding/json"
	"fmt"

	"github.com/dynuc/gophercloud"
	"github.com/dynuc/gophercloud/functiontest/common"
	"github.com/dynuc/gophercloud/openstack"
	"github.com/dynuc/gophercloud/openstack/ces/v1/quotas"
)

func main() {

	fmt.Println("main start...")

	provider, err := common.AuthAKSK()
	//provider, err := common.AuthToken()
	if err != nil {
		fmt.Println("get provider client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	sc, err := openstack.NewCESV1(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get ces client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	QuotaList(sc)
	fmt.Println("main end...")
}

func QuotaList(sc *gophercloud.ServiceClient) {
	quotaRes, err := quotas.Get(sc).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	bytes, _ := json.MarshalIndent(quotaRes, "", " ")
	fmt.Println(string(bytes))
}
