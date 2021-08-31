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
		Message: "db connection missing for db A",
		Code:    1,
	})

	configErr2 := app.NewBadConfigErr(myerrors.BaseError{
		Message: "db connection missing for db B",
		Code:    1,
	})
	var badConfig *app.BadConfigErr
	// As is type assertion
	fmt.Printf("As = %v \n", errors.As(configErr, &badConfig))
	// Is compares against value, we need to implement method is for our err
	fmt.Printf("Is = %v \n", errors.Is(configErr, configErr2))
	printErr(configErr)
	printErr(configErr2)

	badFieldErr := bussines.NewBadFieldForEntity(
		myerrors.BaseError{
			Err:     configErr,
			Message: "bad field value limitField",
			Code:    1000,
		}, "user",
		12345,
	)
	printErr(badFieldErr)
	fmt.Printf("%v\n", errors.Is(configErr, configErr2))
	fmt.Printf("%v\n", errors.As(badFieldErr, &badConfig))
}

func printErr(theError error) {
	fmt.Println(theError)
}
