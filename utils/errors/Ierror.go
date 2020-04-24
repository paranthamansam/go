package errors

type Ierror interface {
	New(data string) error
	Wrap(err error, message string) error
	error
}
