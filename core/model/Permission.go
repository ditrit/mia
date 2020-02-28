//Package model :
// Describe how a permission is stored
package model

import (
	"mia/core/constant"
)

//Permission :
type Permission struct {
	IDRole   uint64
	IDDomain uint64
	IDObject uint64
	Action   constant.Action
	Effect   bool
}
