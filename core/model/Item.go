// Package model :
//Item is an struct that can be a Subject, a Domain, or an object
package model

import (
	"errors"
	"iam/core/constant"
)

//ItemType :
// identifier type to differ a subject, a domain and an object
type ItemType uint16

// identifiers
// nolint: golint, stylecheck
const (
	ITEM_TYPE_SUBJ ItemType = iota
	ITEM_TYPE_DOMAIN
	ITEM_TYPE_OBJ
)

//Item :
// THE item
type Item struct {
	Type ItemType
	ID   uint64 `gorm:"unique;not null"`
	Name string
}

//NewItem :
// Constructor
func NewItem(iType ItemType, name string) (*Item, error) {
	if err := IsNameValidForItem(name); err != nil {
		return nil, err
	}

	if err := IsTypeValid(iType); err != nil {
		return nil, err
	}

	res := new(Item)
	res.Type = iType
	res.Name = name

	return res, nil
}

//IsNameValidForItem :
// Is name valid for item
func IsNameValidForItem(name string) error {
	switch {
	case len(name) > constant.NAME_MAX_LEN:
		return errors.New("the name cannot be longer that 255 characters")
	case len(name) == 0:
		return errors.New("the name cannot be empty")
	}

	return nil
}

//IsTypeValid :
// Is type in the list
func IsTypeValid(iType ItemType) error {
	switch iType {
	case ITEM_TYPE_DOMAIN:
		return nil
	case ITEM_TYPE_OBJ:
		return nil
	case ITEM_TYPE_SUBJ:
		return nil
	}

	return errors.New("the type is not correct")
}

//NewSubject :
// Constructor
func NewSubject(name string) (*Item, error) {
	return NewItem(ITEM_TYPE_SUBJ, name)
}

//NewObject :
// Constructor
func NewObject(name string) (*Item, error) {
	return NewItem(ITEM_TYPE_OBJ, name)
}

//NewDomain :
// Constructor
func NewDomain(name string) (*Item, error) {
	return NewItem(ITEM_TYPE_DOMAIN, name)
}

//GetRoots :
// Used in database initialization
// Used to check if a name is correct
func GetRoots() []*Item {
	res := []*Item{nil, nil, nil}

	res[0], _ = NewDomain(constant.ROOT_DOMAINS)
	res[1], _ = NewSubject(constant.ROOT_SUBJECTS)
	res[2], _ = NewObject(constant.ROOT_OBJECTS)

	return res
}

//GetRootNameWithType :
// Used to check if a name is correct
func GetRootNameWithType(iType ItemType) (string, error) {
	roots := GetRoots()

	if err := IsTypeValid(iType); err != nil {
		return "", errors.New("the type is unknown")
	}

	for _, elem := range roots {
		if elem.Type == iType {
			return elem.Name, nil
		}
	}

	panic("should never happened")
}
