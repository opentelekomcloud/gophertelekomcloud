package gpfs

import (
	"errors"
)

// ListFS lists file systems.
//
// You can use this API to obtain the file system list. In the list, file system names are displayed in lexicographical order.
func (client Client) ListFS(input *ListFSInput) (output *ListFSOutput, err error) {
	if input == nil {
		return nil, errors.New("ListFSInput is nil")
	}
	output = &ListFSOutput{}
	err = client.doActionWithoutFS(httpGet, input, output)
	if err != nil {
		output = nil
	}
	return
}

// CreateFS creates a file system.
//
// The file system name must be unique in the region.
func (client Client) CreateFS(input *CreateFSInput) (output *BaseModel, err error) {
	if input == nil {
		return nil, errors.New("CreateFSInput is nil")
	}
	output = &BaseModel{}
	err = client.doActionWithFS(httpPut, input.FSName, input, output)
	if err != nil {
		output = nil
	}
	return
}

// DeleteFS deletes a GPFS.
//
// You can use this API to delete a storage file system.
func (client Client) DeleteFS(fsName string) (output *BaseModel, err error) {
	output = &BaseModel{}
	err = client.doActionWithFS(httpDelete, fsName, emptyInput, output)
	if err != nil {
		output = nil
	}
	return
}

// CreateFSAccessRules configures access rules for a general-purpose file system.
//
// Each rule grants a VPC read/write or read-only access. Configuring rules replaces
// the current file system access-rule configuration.
func (client Client) CreateFSAccessRules(input *CreateFSAccessRulesInput) (output *BaseModel, err error) {
	if input == nil {
		return nil, errors.New("CreateFSAccessRulesInput is nil")
	}
	output = &BaseModel{}
	err = client.doActionWithFSResult(httpPut, input.FSName, input, output, false)
	if err != nil {
		output = nil
	}
	return
}

// GetFSAccessRules queries the access rules configured for a general-purpose file system.
func (client Client) GetFSAccessRules(fsName string) (output *GetFSAccessRulesOutput, err error) {
	output = &GetFSAccessRulesOutput{Statement: []AccessRule{}}
	err = client.doActionWithFSResult(httpGet, fsName, newSubResourceInput(subResourceSFSACL), output, false)
	if err != nil {
		output = nil
	}
	return
}

// DeleteFSAccessRules deletes all access rules configured for a general-purpose file system.
func (client Client) DeleteFSAccessRules(fsName string) (output *BaseModel, err error) {
	output = &BaseModel{}
	err = client.doActionWithFSResult(httpDelete, fsName, newSubResourceInput(subResourceSFSACL), output, false)
	if err != nil {
		output = nil
	}
	return
}
