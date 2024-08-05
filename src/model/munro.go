package model

type Munro struct {
	Id       string  `json:Id`
	Name     string  `json:Name`
	Height   float32 `json:Height`
	Location string  `json:Location`
}

/*
func (m *Munro) NewMunro() *Munro {
	return &Munro{
		id:       "1",
		name:     "Ben Nevis",
		height:   1345.0,
		location: "Grampian Mountains",
	}
}
*/
