package errors

import (
	"github.com/google/uuid"
)

//myError : custom error struct
type myError struct {
	id      string //Error Id auto generated
	message string //Error description
	*stack         //Error stackTrace
	error          //chain this object to store list of errors
}

type errorList struct {
	[]error
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

	data := &myError{
		id:      uuid.New().String(),
		message: message,
		stack:   callers(),
		error:   err,
	}
	data.error
}

func (er myError) Extend


// func (er myError) GetErrors() []error{

// 	var errs []error
// 	if er.
// 	errs = append(errs,er.error.( ))


}

//Error : for implementation
func (er myError) Error() string {
	return er.message
}
