// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package transcribeservice

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeBadRequestException for service response error code
	// "BadRequestException".
	//
	// Your request didn't pass one or more validation tests. For example, if the
	// entity that you're trying to delete doesn't exist or if it is in a non-terminal
	// state (for example, it's "in progress"). See the exception Message field
	// for more information.
	ErrCodeBadRequestException = "BadRequestException"

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	//
	// There is already a resource with that name.
	ErrCodeConflictException = "ConflictException"

	// ErrCodeInternalFailureException for service response error code
	// "InternalFailureException".
	//
	// There was an internal error. Check the error message and try your request
	// again.
	ErrCodeInternalFailureException = "InternalFailureException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// Either you have sent too many requests or your input file is too long. Wait
	// before you resend your request, or use a smaller file and resend the request.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeNotFoundException for service response error code
	// "NotFoundException".
	//
	// We can't find the requested resource. Check the name and try your request
	// again.
	ErrCodeNotFoundException = "NotFoundException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"BadRequestException":      newErrorBadRequestException,
	"ConflictException":        newErrorConflictException,
	"InternalFailureException": newErrorInternalFailureException,
	"LimitExceededException":   newErrorLimitExceededException,
	"NotFoundException":        newErrorNotFoundException,
}
