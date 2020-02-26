// Package engine :
// Resolving queries
package engine

import (
	"fmt"
	"iam/core/constant"
	"iam/core/database"
	"iam/core/model"
	"iam/core/utils"
)

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
// the enforce function resolves if a subject can execute an action on an object in a given context
// nolint: funlen, gocyclo
func Enforce(
	idb database.IAMDatabase,
	subjectName string,
	domainName string,
	objectName string,
	action constant.Action,
) (bool, error) {
	var (
		subj            model.Item
		domain          model.Item
		object          model.Item
		ancestorsSubj   utils.IDSet
		ancestorsDomain utils.IDSet
		ancestorsObject utils.IDSet
		assigns         []model.Assignment
		perms           []model.Permission
		err             error
		effects         []bool
	)

	idb.OpenConnection()
	defer idb.CloseConnection()

	// Step 1 : get Subject, Domain and Object

	subj, err = GetItem(idb, false, model.ITEM_TYPE_SUBJ, subjectName)

	if err != nil {
		return false, err
	}

	domain, err = GetItem(idb, false, model.ITEM_TYPE_DOMAIN, domainName)

	if err != nil {
		return false, err
	}

	object, err = GetItem(idb, false, model.ITEM_TYPE_OBJ, objectName)

	if err != nil {
		return false, err
	}

	fmt.Printf("subj %s %d\n", subj.Name, subj.ID)
	fmt.Printf("domain %s %d\n", domain.Name, domain.ID)
	fmt.Printf("object %s %d\n", object.Name, object.ID)

	// Step 2 : get Ancestors

	ancestorsSubj, err = getAncestorOf(idb, false, subj.Type, subj.Name)

	if err != nil {
		return false, err
	}

	ancestorsDomain, err = getAncestorOf(idb, false, domain.Type, domain.Name)

	if err != nil {
		return false, err
	}

	ancestorsObject, err = getAncestorOf(idb, false, object.Type, object.Name)

	if err != nil {
		return false, err
	}

	fmt.Printf("len ancestors subj %d\n", ancestorsSubj.Size())
	fmt.Println(ancestorsSubj)
	fmt.Printf("len ancestors domain %d\n", ancestorsDomain.Size())
	fmt.Println(ancestorsDomain)
	fmt.Printf("len ancestors objects %d\n", ancestorsObject.Size())
	fmt.Println(ancestorsObject)

	// Step 3 : getAssignments

	assigns, err = getWantedAssignments(idb, ancestorsSubj.ToSlice())

	if err != nil {
		return false, err
	}

	fmt.Printf("number assigns : %d\n", len(assigns))

	// Step 4 : getPermissions for given action

	perms, err = getWantedPermission(idb, ancestorsObject.ToSlice(), action)

	if err != nil {
		return false, err
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

			if ancestorsDomain.Contains(assign.IDDomain) && ancestorsDomain.Contains(perm.IDDomain) {
				effects = append(effects, perm.Effect)
			}
		}
	}

	fmt.Printf("number effects : %d\n", len(effects))

	// Step 6 : dealing with effects

	if len(effects) == 0 {
		return false, nil
	}

	for _, eff := range effects {
		if !eff {
			return false, nil
		}
	}

	return true, nil
}
