package gpfs

import (
	"errors"
)

// ListFS lists file systems.
//
// You can use this API to obtain the file system list. In the list, file system names are displayed in lexicographical order.
func (obsClient ObsClient) ListFS(input *ListFSInput) (output *ListFSOutput, err error) {
	if input == nil {
		input = &ListFSInput{}
	}
	output = &ListFSOutput{}
	err = obsClient.doActionWithoutFS("ListFS", HTTP_GET, input, output)
	if err != nil {
		output = nil
	}
	return
}

// CreateFS creates a file system.
//
// You can use this API to file system a bucket and name it as you specify. The created file system name must be unique in OBS.
func (obsClient ObsClient) CreateFS(input *CreateFSInput) (output *BaseModel, err error) {
	if input == nil {
		return nil, errors.New("CreateFSInput is nil")
	}
	output = &BaseModel{}
	err = obsClient.doActionWithFS("CreateFS", HTTP_PUT, input.FSName, input, output)
	if err != nil {
		output = nil
	}
	return
}

// DeleteFS deletes a GPFS.
//
// You can use this API to delete a storage file system.
func (obsClient ObsClient) DeleteFS(fsName string) (output *BaseModel, err error) {
	output = &BaseModel{}
	err = obsClient.doActionWithFS("DeleteFS", HTTP_DELETE, fsName, defaultSerializable, output)
	if err != nil {
		output = nil
	}
	return
}
