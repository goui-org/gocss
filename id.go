package gocss

import "strconv"

var id = 0

func generateId() string {
	id++
	return "_" + strconv.Itoa(id)
}
