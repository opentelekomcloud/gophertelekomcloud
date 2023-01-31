/*
Package limits shows rate and limit information for a tenant/project.

Example to Retrieve Limits for a Tenant

	getOpts := limits.GetOpts{
		TenantId: "tenant-id",
	}

	limits, err := limits.Get(computeClient, getOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", limits)
*/
package limits
