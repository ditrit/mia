// Package engine :
// Resolving queries
package engine

import (
	"errors"
	"iam/core/constant"
	"iam/core/database"
)

//Enforce :
// the enforce function
// TODO description
func Enforce(
	idb database.IAMDatabase,
	subjectName string,
	domainName string,
	objectName string,
	action constant.Action,
) (constant.Effect, error) {
	//TODO
	return constant.EFFECT_DENY, errors.New("not implemented")
}
