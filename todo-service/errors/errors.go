package myerrors

import (
	"errors"
	"fmt"
	"strconv"

	"todo-service/errors/app"
	"todo-service/errors/bussines"
)

type BaseError struct {
	Err     error
	Message string
	Code    int
}

func NewBaseError(err error, message string, code int) *BaseError {
	return &BaseError{Err: err, Message: message, Code: code}
}

func (e *BaseError) Error() string {
	return strconv.Itoa(e.Code) + " " + e.Message + ": " + e.Err.Error()
}

func TestErrors() {
	configErr := app.NewBadConfigErr(BaseError{
		Message: "db connection missing for db A",
		Code:    1,
	})

	configErr2 := app.NewBadConfigErr(BaseError{
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
		BaseError{
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
