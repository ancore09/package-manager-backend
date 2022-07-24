package model

type Dependency struct {
	Id       uint64       `json:"id"`
	DependOn []Dependency `json:"dependOn,omitempty"`
}

type Tree struct {
	Dependencies []Dependency `json:"q"`
}
