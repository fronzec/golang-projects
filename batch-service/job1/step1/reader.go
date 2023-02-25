package step1

import "github.com/chararch/gobatch/file"

type Person struct {
	FirstName  string `order:"0"`
	LastName   string `order:"1"`
	Email      string `order:"2"`
	Profession string `order:"3"`
}

var CsvFile = file.FileObjectModel{
	FileStore:      &file.LocalFileSystem{},
	FileName:       "resources\\sample-persons-1k.csv", // route is relative path to main file
	Type:           file.CSV,
	Encoding:       "utf-8",
	FieldSeparator: ",",
	Header:         false,
	ItemPrototype:  &Person{},
}
