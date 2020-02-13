//package main :
// File main.go
package main

import (
	"iam/core"
	"iam/core/model"
)

func main() {
	iam := core.NewIAM("test.db", true)

	s, _ := model.NewSubject("coucou")

	_ = iam.AddSubject(*s)
}
