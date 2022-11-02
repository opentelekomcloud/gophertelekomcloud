/*
Package quotasets enables retrieving and managing Block Storage quotas.

Example to Get a Quota Set

	quotaset, err := quotasets.Get(blockStorageClient, "project-id").Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", quotaset)

Example to Get Quota Set Usage

	quotaset, err := quotasets.GetUsage(blockStorageClient, "project-id").Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", quotaset)

Example to Update a Quota Set

	updateOpts := quotasets.UpdateOpts{
		Volumes: golangsdk.IntToPointer(100),
	}

	quotaset, err := quotasets.Update(blockStorageClient, "project-id", updateOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", quotaset)

Example to Update a Quota set with volume_type quotas

	updateOpts := quotasets.UpdateOpts{
		Volumes: golangsdk.IntToPointer(100),
		Extra: map[string]interface{}{
			"gigabytes_foo": golangsdk.IntToPointer(100),
			"snapshots_foo": golangsdk.IntToPointer(10),
			"volumes_foo":   golangsdk.IntToPointer(10),
		},
	}

	quotaset, err := quotasets.Update(blockStorageClient, "project-id", updateOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", quotaset)

Example to Delete a Quota Set

	err := quotasets.Delete(blockStorageClient, "project-id").ExtractErr()
	if err != nil {
		panic(err)
	}
*/
package quotasets
