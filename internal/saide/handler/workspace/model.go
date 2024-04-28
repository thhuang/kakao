package workspace

import (
	"github.com/google/uuid"
)

type getWorkspaceResponse struct {
	Id            string    `json:"id"`
	FileStructure Directory `json:"file_structure"`
}

type Directory map[string][]interface{}

type File map[string]string

func (d Directory) isValid() bool {
	if len(d) != 1 {
		return false
	}

	for _, vs := range d {
		for _, v := range vs {
			switch t := v.(type) {
			case Directory:
				if !t.isValid() {
					return false
				}
			case File:
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

func (f File) isValid() bool {
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
