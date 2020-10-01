package utils

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"reflect"
	"strings"
)

func DeleteNotPassParams(params *map[string]interface{}, notPassParams []string) {
	for _, i := range notPassParams {
		delete(*params, i)
	}
}

// merges two interfaces. In cases where a value is defined for both 'overridingInterface' and
// 'inferiorInterface' the value in 'overridingInterface' will take precedence.
func MergeInterfaces(overridingInterface, inferiorInterface interface{}) interface{} {
	switch overriding := overridingInterface.(type) {
	case map[string]interface{}:
		interfaceMap, ok := inferiorInterface.(map[string]interface{})
		if !ok {
			return overriding
		}
		for k, v := range interfaceMap {
			if overridingValue, ok := overriding[k]; ok {
				overriding[k] = MergeInterfaces(overridingValue, v)
			} else {
				overriding[k] = v
			}
		}
	case []interface{}:
		list, ok := inferiorInterface.([]interface{})
		if !ok {
			return overriding
		}
		for i := range list {
			overriding = append(overriding, list[i])
		}
		return overriding
	case nil:
		// mergeClouds(nil, map[string]interface{...}) -> map[string]interface{...}
		v, ok := inferiorInterface.(map[string]interface{})
		if ok {
			return v
		}
	}
	// We don't want to override with empty values
	if reflect.DeepEqual(overridingInterface, nil) || reflect.DeepEqual(reflect.Zero(reflect.TypeOf(overridingInterface)).Interface(), overridingInterface) {
		return inferiorInterface
	} else {
		return overridingInterface
	}
}

func PrependString(item string, slice []string) []string {
	newSize := len(slice) + 1
	result := make([]string, newSize, newSize)
	result[0] = item
	for i, v := range slice {
		result[i+1] = v
	}
	return result
}

func In(item interface{}, slice interface{}) bool {
	for _, it := range slice.([]interface{}) {
		if reflect.DeepEqual(item, it) {
			return true
		}
	}
	return false
}

// GetRegion returns the region that was specified in the auth options. If a
// region was not set, the provider-level region is checked. The provider-level
// region can either be set by the region argument or by OS_REGION_NAME.
func GetRegion(authOpts golangsdk.AuthOptions) string {
	n := authOpts.TenantName
	if n == "" {
		n = authOpts.DelegatedProject
	}
	return strings.Split(n, "_")[0]
}
