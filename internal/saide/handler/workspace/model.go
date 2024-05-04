package workspace

import (
	"github.com/google/uuid"
	"github.com/thhuang/kakao/internal/saide/service/workspace"
)

type GetWorkspaceResponse struct {
	Id            string    `json:"id"`
	FileStructure Directory `json:"file_structure"`
}

func NewGetWorkspaceResponse(workspace *workspace.Workspace) (*GetWorkspaceResponse, error) {
	// TODO: parse the file structure from the workspace object

	fileStructure := Directory{
		"/": {
			File{
				"README.md": uuid.NewString(),
			},
			Directory{
				"cmd": {
					File{
						"main.go": uuid.NewString(),
					},
				},
			},
			Directory{
				"pkg": {
					Directory{
						"service": {
							Directory{
								"mongo": {
									File{
										"mongo.go": uuid.NewString(),
									},
									File{
										"impl.go": uuid.NewString(),
									},
									File{
										"impl_test.go": uuid.NewString(),
									},
								},
							},
							Directory{
								"redis": {
									File{
										"redis.go": uuid.NewString(),
									},
									File{
										"impl.go": uuid.NewString(),
									},
									File{
										"impl_test.go": uuid.NewString(),
									},
								},
							},
						},
					},
				},
			},
		},
	}

	if !fileStructure.isValid() {
		return nil, ErrInvalidFileStructure
	}

	return &GetWorkspaceResponse{
		Id:            workspace.Id.String(),
		FileStructure: fileStructure,
	}, nil
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
