package step1

import "github.com/chararch/gobatch/file"

type Person struct {
	FirstName  string `order:"0" db:"first_name"`
	LastName   string `order:"1" db:"last_name"`
	Email      string `order:"2" db:"email"`
	Profession string `order:"3" db:"profession"`
}

var CsvFile = file.FileObjectModel{
	FileStore:      &file.LocalFileSystem{},
	FileName:       "resources/sample-persons-1k.csv", // path is relative to the main package
	Type:           file.CSV,
	Encoding:       "utf-8",
	FieldSeparator: ",",
	Header:         false,
	ItemPrototype:  &Person{},
}
