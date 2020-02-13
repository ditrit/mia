// Package engine :
// Handling items queries from IAM
package engine

import (
	"errors"
	"iam/core/database"
	"iam/core/model"

	"github.com/jinzhu/gorm"
)

// use this function to have the benefit of abstraction and closures
func askDBForItems(
	idb database.IAMDatabase,
	name string,
	iType model.ItemType,
	haveToOpenConnection bool,
	fIfFound func(*gorm.DB, model.Item) (model.Item, error),
	fIfNotFound func(*gorm.DB, model.Item) (model.Item, error),
) (model.Item, error) {
	var queryItem model.Item

	if len(name) == 0 {
		return queryItem, errors.New("the name cannot be empty")
	}

	if haveToOpenConnection {
		idb.OpenConnection()
		defer idb.CloseConnection() //nolint: errcheck
	}

	res := idb.DB().Where("type = ?", iType).Where("name = ?", name).First(&queryItem)

	if res.Error != nil && !res.RecordNotFound() {
		return queryItem, errors.New("unknown error occurred")
	}

	if res.RecordNotFound() {
		return fIfNotFound(res, queryItem)
	}

	return fIfFound(res, queryItem)
}

func askDBForItemLinks(
	idb database.IAMDatabase,
	parentName string,
	childName string,
	iType model.ItemType,
	haveToOpenConnection bool,
	fIfFound func(*gorm.DB, model.ItemLink) error,
	fIfNotFound func(*gorm.DB, model.ItemLink) error,
) error {
	var (
		sLink    model.ItemLink
		parentDB model.Item
		childDB  model.Item
		err      error
	)

	if haveToOpenConnection {
		idb.OpenConnection()
		defer idb.CloseConnection() //nolint: errcheck
	}

	// Search parent item
	parentDB, err = askDBForItems(idb, parentName, iType, false,
		func(db *gorm.DB, qs model.Item) (model.Item, error) {
			return qs, db.Error
		},
		func(db *gorm.DB, qs model.Item) (model.Item, error) {
			return qs, errors.New("the parent item does not exist in the iam")
		})

	if err != nil {
		return err
	}

	// Search child item
	childDB, err = askDBForItems(idb, childName, iType, false,
		func(db *gorm.DB, qs model.Item) (model.Item, error) {
			return qs, db.Error
		},
		func(db *gorm.DB, qs model.Item) (model.Item, error) {
			return qs, errors.New("the child item does not exist in the iam")
		})

	if err != nil {
		return err
	}

	res := idb.DB().Where("id_parent = ?", parentDB.ID).Where("id_child = ?", childDB.ID).Take(&sLink)
	if res.Error != nil && !res.RecordNotFound() {
		return res.Error
	}

	sLink.IDParent = parentDB.ID
	sLink.IDChild = childDB.ID

	if res.RecordNotFound() {
		return fIfNotFound(res, sLink)
	}

	return fIfFound(res, sLink)
}

func checkIfItemsAreConsistents(iType model.ItemType, items ...model.Item) error {
	if len(items) == 0 {
		return nil
	}

	checkIfKnown := func(item model.Item) bool {
		switch item.Type {
		case model.ITEM_TYPE_SUBJ, model.ITEM_TYPE_DOMAIN, model.ITEM_TYPE_OBJ:
			return true
		}

		return false
	}

	for i := range items {
		if !checkIfKnown(items[i]) {
			return errors.New("unknown type")
		}

		if iType != items[i].Type {
			return errors.New("type must be the same for all items")
		}
	}

	return nil
}

// AddItem :
// add item in the IAM
// returns an error if :
//	- the item already exists in the iam
//	- the item has an empty name
func AddItem(
	idb database.IAMDatabase,
	iType model.ItemType,
	s model.Item,
) error {
	if err := checkIfItemsAreConsistents(iType, s); err != nil {
		return err
	}

	_, err := askDBForItems(idb, s.Name, s.Type, true,
		func(_ *gorm.DB, qs model.Item) (model.Item, error) {
			return qs, errors.New("the item already exists in the iam")
		},
		func(db *gorm.DB, qs model.Item) (model.Item, error) {
			db.Error = nil
			res := db.Create(&s)
			return qs, res.Error
		})

	return err
}

// RemoveItem :
// remove item in the IAM
// returns an error if :
//	- the item does not exist in the iam
//	- the item is a parent or a child in ItemLinks
func RemoveItem(
	idb database.IAMDatabase,
	iType model.ItemType,
	s model.Item,
) error {
	if err := checkIfItemsAreConsistents(iType, s); err != nil {
		return err
	}

	_, err := askDBForItems(idb, s.Name, s.Type, true,
		func(db *gorm.DB, qs model.Item) (model.Item, error) {
			res := db.Delete(&s)
			return qs, res.Error
		},
		func(db *gorm.DB, qs model.Item) (model.Item, error) {
			return qs, errors.New("the item does not exist in the iam")
		})

	return err
}

// RenameItem :
// rename item in the IAM
// returns an error if :
//	- the item does not exist in the iam
//	- the new name given is empty
func RenameItem(
	idb database.IAMDatabase,
	iType model.ItemType,
	s model.Item,
	newName string,
) error {
	if err := checkIfItemsAreConsistents(iType, s); err != nil {
		return err
	}

	if newName == "" {
		return errors.New("the new name cannot be empty")
	}

	_, err := askDBForItems(idb, s.Name, s.Type, true,
		func(db *gorm.DB, qs model.Item) (model.Item, error) {
			res := db.Model(&s).Update("name", newName)
			return qs, res.Error
		},
		func(db *gorm.DB, qs model.Item) (model.Item, error) {
			return qs, errors.New("the item does not exist in the iam")
		})

	return err
}

// GetItem :
// get item in the IAM
// returns an error if :
//	- the item does not exist in the iam
func GetItem(
	idb database.IAMDatabase,
	iType model.ItemType,
	name string,
) (model.Item, error) {
	return askDBForItems(idb, name, iType, true,
		func(db *gorm.DB, qs model.Item) (model.Item, error) {
			return qs, db.Error
		},
		func(db *gorm.DB, qs model.Item) (model.Item, error) {
			return qs, errors.New("the item does not exist in the iam")
		})
}

// AddItemLink :
// add a relation between two items in the IAM
// returns an error if :
//	- one of the items does not exists in the iam
//	- the link already exists
func AddItemLink(
	idb database.IAMDatabase,
	iType model.ItemType,
	sParent model.Item,
	sChild model.Item,
) error {
	if err := checkIfItemsAreConsistents(iType, sParent, sChild); err != nil {
		return err
	}

	return askDBForItemLinks(idb, sParent.Name, sChild.Name, sParent.Type, true,
		func(db *gorm.DB, qs model.ItemLink) error {
			return errors.New("the connection link already exists")
		},
		func(db *gorm.DB, qs model.ItemLink) error {
			db.Error = nil
			res := db.Create(&qs)
			return res.Error
		},
	)
}

// RemoveItemLink :
// remove a relation between two items in the IAM
// returns an error if :
//	- the link does not exist
func RemoveItemLink(
	idb database.IAMDatabase,
	iType model.ItemType,
	sParent model.Item,
	sChild model.Item,
) error {
	if err := checkIfItemsAreConsistents(iType, sParent, sChild); err != nil {
		return err
	}

	return askDBForItemLinks(idb, sParent.Name, sChild.Name, sParent.Type, true,
		func(db *gorm.DB, qs model.ItemLink) error {
			res := db.Delete(&qs)
			return res.Error
		},
		func(db *gorm.DB, qs model.ItemLink) error {
			return errors.New("the connection link does not exist")
		},
	)
}

// AddItemArchitecture :
// add an architecture to the IAM
// returns an error if :
//	- tabs given have not the same size
//	- one of the items already exists in the iam
//	- one of the items has an empty name
//	- one of the links alrady exists
// We can ignore some of this error with the other parameters
func AddItemArchitecture(
	idb database.IAMDatabase,
	iType model.ItemType,
	parents []model.Item,
	childs []model.Item,
	ignoreAlreadyExistsItem bool,
	ignoreAlreadyExistsLinks bool,
) error {
	// TODO: implement + test
	return nil
}
