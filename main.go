package main

import (
	"fmt"

	"github.com/sam619/go/utils/errors"
)

func main() {
	if err := TestError(); err != nil {
		fmt.Println(err.Error())

		fmt.Printf("%+v", err)
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

//returnError : return error
