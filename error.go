package fsm

import "errors"

var (
	ErrMachineCreationFailed = errors.New("machine creation failed")
	ErrTransitionFailed      = errors.New("transition failed")
	ErrEventDeclined         = errors.New("event declined")
	ErrStateUndefined        = errors.New("state undefined")
	ErrInitialStateUndefined = errors.New("initial state undefined")
)
