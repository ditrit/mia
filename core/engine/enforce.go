// Package engine :
// Resolving queries
package engine

import (
	"errors"
	"fmt"
	"iam/core/constant"
	"iam/core/database"
	"iam/core/model"
	"iam/core/utils"
)

func getAncestorOf(
	idb database.IAMDatabase,
	item model.Item,
) ([]uint64, error) {
	res := []uint64{}
	resSet := utils.NewIDSet()
	vertices, parentTable, err := GetItemArchitecture(idb, false, item.Type)
	mapNameToItem := make(map[string]model.Item)

	if err != nil {
		return res, err
	}

	// Verify if item is in vertices
	found := false

	for i := range vertices {
		name := vertices[i].Name
		if name == item.Name {
			found = true
		}

		mapNameToItem[name] = vertices[i]
	}

	if !found {
		return res, errors.New("SNH: item wasn't found in the whole architecture")
	}

	setToVisit := utils.NewStringSet()

	setToVisit.Add(item.Name)

	for key, empty := setToVisit.Pop(); !empty; {
		resSet.Add(mapNameToItem[key].ID)

		for index := range parentTable[key] {
			name := parentTable[key][index].Name
			setToVisit.Add(name)
		}
	}

	return resSet.ToSlice(), nil
}

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
	var (
		subj          model.Item
		domain        model.Item
		object        model.Item
		ancestorsSubj []uint64
		err           error
	)

	idb.OpenConnection()
	defer idb.CloseConnection()

	// Step 1 : get Subject, Domain and Object

	subj, err = GetItem(idb, false, model.ITEM_TYPE_SUBJ, subjectName)

	if err != nil {
		return constant.EFFECT_DENY, err
	}

	domain, err = GetItem(idb, false, model.ITEM_TYPE_DOMAIN, domainName)

	if err != nil {
		return constant.EFFECT_DENY, err
	}

	object, err = GetItem(idb, false, model.ITEM_TYPE_OBJ, objectName)

	if err != nil {
		return constant.EFFECT_DENY, err
	}

	fmt.Printf("subj %s %d", subj.Name, subj.ID)
	fmt.Printf("domain %s %d", domain.Name, domain.ID)
	fmt.Printf("object %s %d", object.Name, object.ID)

	ancestorsSubj, err = getAncestorOf(idb, subj)

	if err != nil {
		return constant.EFFECT_DENY, err
	}

	fmt.Printf("len ancestors subj %d\n", len(ancestorsSubj))

	//TODO
	return constant.EFFECT_DENY, errors.New("not implemented")
}
