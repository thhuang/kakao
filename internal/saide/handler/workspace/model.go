package workspace

import (
	"github.com/google/uuid"
)

type getWorkspaceResponse struct {
	Id            string    `json:"id"`
	FileStructure directory `json:"file_sturcture"`
}

type directory map[string][]interface{}

type file map[string]string

func (d directory) isValid() bool {
	if len(d) != 1 {
		return false
	}

	for _, vs := range d {
		for _, v := range vs {
			switch t := v.(type) {
			case directory:
				if !t.isValid() {
					return false
				}
			case file:
				if !t.isValid() {
					return false
				}
			default:
				return false
			}
		}
	}

	return true
}

func (f file) isValid() bool {
	if len(f) != 1 {
		return false
	}

	for _, v := range f {
		if _, err := uuid.Parse(v); err != nil {
			return false
		}
	}

	return true
}
