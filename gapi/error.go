package gapi

import (
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func fieldViolation(field string, err error) *errdetails.BadRequest_FieldViolation {
	return &errdetails.BadRequest_FieldViolation{
		Field:       field,
		Description: err.Error(),
	}
}

func invalidArgumentError(violations []*errdetails.BadRequest_FieldViolation) error {
	br := &errdetails.BadRequest{FieldViolations: violations}
	si := status.New(codes.InvalidArgument, "invalid parameters")

	st, err := si.WithDetails(br)
	if err != nil {
		return si.Err()
	}
	return st.Err()
}

func unauthenticatedError(err error) error {
	return status.Errorf(codes.Unauthenticated, "authentication required: %v", err)
}
