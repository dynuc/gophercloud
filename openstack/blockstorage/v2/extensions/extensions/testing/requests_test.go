package testing

import (
	"testing"

	"github.com/dynuc/gophercloud/openstack/blockstorage/v2/extensions/extensions"
	"github.com/dynuc/gophercloud/pagination"
	th "github.com/dynuc/gophercloud/testhelper"
	"github.com/dynuc/gophercloud/testhelper/client"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListExtensionsSuccessfully(t)

	count := 0

	extensions.List(client.ServiceClient()).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := extensions.ExtractExtensions(page)
		th.AssertNoErr(t, err)
		th.AssertDeepEquals(t, ExpectedExtensions, actual)

		return true, nil
	})

	th.CheckEquals(t, 1, count)
}
