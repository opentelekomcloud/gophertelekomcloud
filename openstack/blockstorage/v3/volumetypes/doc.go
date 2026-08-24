/*
Package volumetypes provides information and interaction with volume types in the
OpenStack Block Storage service. A volume type is a collection of specs used to
define the volume capabilities.

Example to list Volume Types

	volumeTypes, err := volumetypes.List(client, volumetypes.ListOpts{})
	if err != nil{
		panic(err)
	}
	for _,vt := range volumeTypes{
		fmt.Println(vt)
	}

Example to show a Volume Type

	typeID := "7ffaca22-f646-41d4-b79d-d7e4452ef8cc"
	volumeType, err := volumetypes.Get(client, typeID)
	if err != nil{
		panic(err)
	}
	fmt.Println(volumeType)

Example to create a Volume Type

	isPublic := true
	volumeType, err := volumetypes.Create(client, volumetypes.CreateOpts{
		Name:"volume_type_001",
		IsPublic:&isPublic,
		Description:"description_001",
	})
	if err != nil{
		panic(err)
	}
	fmt.Println(volumeType)

Example to delete a Volume Type

	typeID := "7ffaca22-f646-41d4-b79d-d7e4452ef8cc"
	err := volumetypes.Delete(client, typeID)
	if err != nil{
		panic(err)
	}

Example to update a Volume Type

	typeID := "7ffaca22-f646-41d4-b79d-d7e4452ef8cc"
	isPublic := false
	volumeType, err := volumetypes.Update(client, typeID, volumetypes.UpdateOpts{
		Name: "volume_type_002",
		Description:"description_002",
		IsPublic:&isPublic,
	})
	if err != nil{
		panic(err)
	}
	fmt.Println(volumeType)
*/

package volumetypes
