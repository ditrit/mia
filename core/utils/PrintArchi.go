// Package utils :
// Contains the PrintArchi function that just prints the architecture given
// Used in debugging purposes
package utils

import (
	"fmt"
)

// PrintArchi :
func PrintArchi(vertices []string, edges map[string][]string) {
	fmt.Println("====================")
	fmt.Println("Vertices: [")

	for _, elem := range vertices {
		fmt.Printf("\t'%s',\n", elem)
	}

	fmt.Println("]")
	fmt.Println("")
	fmt.Println("Edges: [")

	for _, child := range vertices {
		for _, parent := range edges[child] {
			fmt.Printf("\t'%s' --> '%s',\n", child, parent)
		}
	}

	fmt.Println("]")
	fmt.Println("====================")
}
