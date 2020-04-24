package main

import (
	"fmt"

	"github.com/paranthamansam/go/utils/errors"
)

func main() {
	if err := TestError(); err != nil {
		fmt.Println(err.Error())

		fmt.Printf("%+v", err)
	}

	if errs := TestErrors(); errs != nil {
		fmt.Println(errs.Error())
	}

}

//TestError : test the custom error
func TestError() error {
	if err := errors.New("sample error"); err != nil {
		err1 := errors.Wrap(err, "new Error")
		return err1
	}
	// fmt.Printf("our data : %v ", err)
	return nil
}

//TestErrors : Error list
func TestErrors() error {
	errs := errors.Errors{}
	errs.Append(errors.New("one"))
	errs.Append(errors.New("two"))
	errs.Append(TestError())

	return errs
}
