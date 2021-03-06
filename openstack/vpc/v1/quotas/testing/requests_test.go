package testing

import (
	"testing"

	"github.com/dynuc/gophercloud/openstack/vpc/v1/quotas"
	th "github.com/dynuc/gophercloud/testhelper"
	"github.com/dynuc/gophercloud/testhelper/client"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	actual, err := quotas.List(client.ServiceClient(), quotas.ListOpts{
		Type: "vpc",
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &ListResponse, actual)
}
