package model

type Package struct {
	Id   uint64
	Name string
}

type Packages []*Package

func (p Packages) FilterByName(name string) *Package {
	for _, pack := range p {
		if pack.Name == name {
			return pack
		}
	}
	return nil
}
