package models

type Layout struct {
	// TODO: define layouts
}

type Result struct {
	// TODO: define http result
}

type Person struct {
	Name string `csv:"Name"`
	CPF  string `csv:"CPF"`
	City string `csv:"City"`
}
