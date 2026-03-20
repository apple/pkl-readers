// Code generated from Pkl module `helm.helm`. DO NOT EDIT.
package hookpolicy

import (
	"encoding"
	"fmt"
)

type HookPolicy string

const (
	Enabled   HookPolicy = "enabled"
	SkipTests HookPolicy = "skip-tests"
	Disabled  HookPolicy = "disabled"
)

// String returns the string representation of HookPolicy
func (rcv HookPolicy) String() string {
	return string(rcv)
}

var _ encoding.BinaryUnmarshaler = new(HookPolicy)

// UnmarshalBinary implements encoding.BinaryUnmarshaler for HookPolicy.
func (rcv *HookPolicy) UnmarshalBinary(data []byte) error {
	switch str := string(data); str {
	case "enabled":
		*rcv = Enabled
	case "skip-tests":
		*rcv = SkipTests
	case "disabled":
		*rcv = Disabled
	default:
		return fmt.Errorf(`illegal: "%s" is not a valid HookPolicy`, str)
	}
	return nil
}
