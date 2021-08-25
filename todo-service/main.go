package main

import (
	"errors"
	"fmt"
	"todo-service/errors"
	"todo-service/errors/app"
	"todo-service/errors/bussines"
)

func main() {
	configErr := app.NewBadConfigErr(myerrors.BaseError{
		Message: "db connection missing",
		Code:    0,
	})

	configErr2 := app.NewBadConfigErr(myerrors.BaseError{
		Message: "db connection missing",
		Code:    0,
	})
	var badConfig *app.BadConfigErr
	fmt.Printf("%v \n", errors.As(configErr, &badConfig))
	fmt.Printf("%v \n", errors.Is(configErr, configErr2))
	printErr(configErr)
	printErr(configErr2)

	badFieldErr := bussines.NewBadFieldForEntity(
		myerrors.BaseError{
			Err: configErr,
			Message: "bad field value limitField",
			Code:    1000,
		}, "user",
		12345,
	)
	printErr(badFieldErr)
	fmt.Printf("%v", errors.Is(configErr, configErr2))
}

func printErr(theError error) {
	fmt.Println(theError)
}