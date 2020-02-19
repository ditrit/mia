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

	for {
		key, empty := setToVisit.Pop()
		if !empty {
			break
		}

		// fmt.Println(key)

		resSet.Add(mapNameToItem[key].ID)

		for index := range parentTable[key] {
			name := parentTable[key][index].Name
			setToVisit.Add(name)
		}
	}

	return resSet.ToSlice(), nil
}

func getWantedAssignments(
	idb database.IAMDatabase,
	subjectIDs []uint64,
) ([]model.Assignment, error) {
	var assigns []model.Assignment

	if len(subjectIDs) == 0 {
		return assigns, nil
	}

	tx := idb.DB()
	tx = tx.Where("id_subject IN (?)", subjectIDs)
	res := tx.Find(&assigns)

	return assigns, res.Error
}

func getWantedPermission(
	idb database.IAMDatabase,
	objectsIDs []uint64,
	act constant.Action,
) ([]model.Permission, error) {
	var perms []model.Permission

	if len(objectsIDs) == 0 {
		return perms, nil
	}

	tx := idb.DB()
	tx = tx.Where("action = ?", act)
	tx = tx.Where("id_object IN (?)", objectsIDs)

	res := tx.Find(&perms)

	return perms, res.Error
}

//Enforce :
// the enforce function
// TODO description
// nolint: funlen, gocyclo
func Enforce(
	idb database.IAMDatabase,
	subjectName string,
	domainName string,
	objectName string,
	action constant.Action,
) (constant.Effect, error) {
	var (
		subj               model.Item
		domain             model.Item
		object             model.Item
		ancestorsSubj      []uint64
		ancestorsDomain    []uint64
		ancestorsDomainSet utils.IDSet
		ancestorsObject    []uint64
		assigns            []model.Assignment
		perms              []model.Permission
		err                error
		effects            []constant.Effect
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

	fmt.Printf("subj %s %d\n", subj.Name, subj.ID)
	fmt.Printf("domain %s %d\n", domain.Name, domain.ID)
	fmt.Printf("object %s %d\n", object.Name, object.ID)

	// Step 2 : get Ancestors

	ancestorsSubj, err = getAncestorOf(idb, subj)

	if err != nil {
		return constant.EFFECT_DENY, err
	}

	ancestorsDomain, err = getAncestorOf(idb, domain)

	if err != nil {
		return constant.EFFECT_DENY, err
	}

	ancestorsDomainSet = utils.NewIDSetFromSlice(ancestorsDomain)

	ancestorsObject, err = getAncestorOf(idb, object)

	if err != nil {
		return constant.EFFECT_DENY, err
	}

	fmt.Printf("len ancestors subj %d\n", len(ancestorsSubj))
	fmt.Println(ancestorsSubj)
	fmt.Printf("len ancestors domain %d\n", len(ancestorsDomain))
	fmt.Println(ancestorsDomain)
	fmt.Printf("len ancestors objects %d\n", len(ancestorsObject))
	fmt.Println(ancestorsObject)

	// Step 3 : getAssignments

	assigns, err = getWantedAssignments(idb, ancestorsSubj)

	if err != nil {
		return constant.EFFECT_DENY, err
	}

	fmt.Printf("number assigns : %d\n", len(assigns))

	// Step 4 : getPermissions for given action

	perms, err = getWantedPermission(idb, ancestorsObject, action)

	if err != nil {
		return constant.EFFECT_DENY, err
	}

	fmt.Printf("number perms : %d\n", len(perms))

	// Step 5 : resolve effects that apply

	for indexAssign := range assigns {
		for indexPerm := range perms {
			assign := assigns[indexAssign]
			perm := perms[indexPerm]

			if assign.IDRole != perm.IDRole {
				continue
			}

			if ancestorsDomainSet.Contains(assign.IDDomain) && ancestorsDomainSet.Contains(perm.IDDomain) {
				effects = append(effects, perm.Effect)
			}
		}
	}

	fmt.Printf("number effects : %d\n", len(effects))

	// Step 6 : dealing with effects

	if len(effects) == 0 {
		return constant.EFFECT_DENY, nil
	}

	for _, eff := range effects {
		if eff == constant.EFFECT_DENY {
			return constant.EFFECT_DENY, nil
		}
	}

	return constant.EFFECT_ALLOW, nil
}
