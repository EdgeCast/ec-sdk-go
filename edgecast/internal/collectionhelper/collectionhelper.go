// Copyright 2021 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

// Package collectionhelper provides helper methods for working with
// aggregate/collection types
package collectionhelper

// IsInterfaceArray determines if an interface{} is actually an []interface{}
func IsInterfaceArray(input interface{}) bool {
	switch input.(type) {
	case []interface{}:
		return true
	default:
		return false
	}
}
