package stacks

import (
	"fmt"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

type ErrInvalidEnvironment struct {
	golangsdk.BaseError
	Section string
}

func (e ErrInvalidEnvironment) Error() string {
	return fmt.Sprintf("Environment has wrong section: %s", e.Section)
}

type ErrInvalidDataFormat struct {
	golangsdk.BaseError
}

func (e ErrInvalidDataFormat) Error() string {
	return "Data in neither json nor yaml format."
}

type ErrInvalidTemplateFormatVersion struct {
	golangsdk.BaseError
	Version string
}

func (e ErrInvalidTemplateFormatVersion) Error() string {
	return "Template format version not found."
}
