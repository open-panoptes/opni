package validation

import (
	"fmt"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"buf.build/go/protovalidate"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ValidationError struct {
	message string
}

func (e *ValidationError) Error() string {
	return e.message
}

func (e *ValidationError) GRPCStatus() *status.Status {
	return status.New(codes.InvalidArgument, e.message)
}

func Error(msg string) error {
	return &ValidationError{
		message: msg,
	}
}

func Errorf(format string, args ...interface{}) error {
	return &ValidationError{
		message: fmt.Errorf(format, args...).Error(),
	}
}

func ErrorsFromProto(p *validate.Violations) *protovalidate.ValidationError {
	var errs protovalidate.ValidationError
	if p == nil {
		return &errs
	}

	for _, err := range p.Violations {
		errs.Violations = append(errs.Violations, &protovalidate.Violation{
			Proto: err,
		})
	}

	return &errs
}
