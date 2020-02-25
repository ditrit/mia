// Package engine :
//  contains the function getAncestorOf
package engine

import (
	"errors"
	"iam/core/database"
	"iam/core/model"
	"iam/core/utils"
)

//nolint: unparam
// remove this above if haveToOpenConnection is needed
func getAncestorOf(
	idb database.IAMDatabase,
	haveToOpenConnection bool,
	itemName string,
	iType model.ItemType,
) (utils.IDSet, error) {
	resSet := utils.NewIDSet()
	vertices, parentTable, err := GetItemArchitecture(idb, false, iType)
	mapNameToItem := make(map[string]model.Item)

	if err != nil {
		return resSet, err
	}

	if haveToOpenConnection {
		idb.OpenConnection()
		defer idb.CloseConnection() //nolint: errcheck
	}

	// Verify if item is in vertices
	found := false

	for i := range vertices {
		name := vertices[i].Name
		if name == itemName {
			found = true
		}

		mapNameToItem[name] = vertices[i]
	}

	if !found {
		return resSet, errors.New("SNH: item wasn't found in the whole architecture")
	}

	setToVisit := utils.NewStringSet()
	alreadyVisited := utils.NewStringSet()

	setToVisit.Add(itemName)
	alreadyVisited.Add(itemName)

	for {
		key, empty := setToVisit.Pop()
		if !empty {
			break
		}

		resSet.Add(mapNameToItem[key].ID)

		for index := range parentTable[key] {
			name := parentTable[key][index].Name

			if !alreadyVisited.Contains(name) {
				setToVisit.Add(name)
				alreadyVisited.Add(name)
			}
		}
	}

	return resSet, nil
}
