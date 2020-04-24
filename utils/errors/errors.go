package errors

import (
	"fmt"

	"github.com/google/uuid"
)

//myError : custom error struct
type myError struct {
	id      string //Error Id auto generated
	message string //Error description
	*stack         //Error stackTrace
	error          //chain this object to store list of errorst
}

//Errors : errors type
type Errors struct {
	Type   string    // error type
	errors []myError //error list
}

//New : new custom error
func New(data string) error {
	return &myError{
		id:      uuid.New().String(),
		message: data,
		stack:   callers(),
	}
}

//Wrap : error with additional stact details
func Wrap(err error, message string) error {

	if err == nil {
		return nil
	}

	newErr := &myError{
		id:      uuid.New().String(),
		message: message,
		stack:   callers(),
		error:   err,
	}

	return newErr
}

//Error : for implementation
func (er myError) Error() string {
	return er.message
}

//Append : error appending
func (errs Errors) Append(err error) {
	// var actualError error
	// var ok bool
	if len(errs.errors) == 0 {
		errs.errors = []myError{}
	}
	if err != nil {
		errs.Type = "Error list"
		if actualError, ok = err.(myError); ok {
			errs.errors = append(errs.errors, actualError)
		}
		fmt.Printf("test: %+v", actualError)
		// errs.errors = append(errs.errors.([]error), err)
	}
}

//Error : error list data
func (errs Errors) Error() string {
	return errs.Type
}

// func (er myError) GetErrors() []error{
// 	var errs []error
// 	if er.
// 	errs = append(errs,er.error.( ))
