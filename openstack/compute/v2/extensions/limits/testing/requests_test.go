package testing

import (
	"testing"

	"github.com/dynuc/gophercloud/openstack/compute/v2/extensions/limits"
	th "github.com/dynuc/gophercloud/testhelper"
	"github.com/dynuc/gophercloud/testhelper/client"
)

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t)

	getOpts := limits.GetOpts{
		TenantID: TenantID,
	}

	actual, err := limits.Get(client.ServiceClient(), getOpts).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &LimitsResult, actual)
}
