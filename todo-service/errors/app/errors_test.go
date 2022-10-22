package app

import (
	"reflect"
	"testing"

	myerrors "todo-service/errors"
)

func TestNewBadConfigErr(t *testing.T) {
	type args struct {
		baseError myerrors.BaseError
	}
	tests := []struct {
		name string
		args args
		want *BadConfigErr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := NewBadConfigErr(tt.args.baseError); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("NewBadConfigErr() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestBadConfigErr_Error(t *testing.T) {
	type fields struct {
		BaseError myerrors.BaseError
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				e := &BadConfigErr{
					BaseError: tt.fields.BaseError,
				}
				if got := e.Error(); got != tt.want {
					t.Errorf("Error() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestBadConfigErr_Is(t *testing.T) {
	type fields struct {
		BaseError myerrors.BaseError
	}
	type args struct {
		target error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				e := &BadConfigErr{
					BaseError: tt.fields.BaseError,
				}
				if got := e.Is(tt.args.target); got != tt.want {
					t.Errorf("Is() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestNewBadConfigErr1(t *testing.T) {
	type args struct {
		baseError myerrors.BaseError
	}
	tests := []struct {
		name string
		args args
		want *BadConfigErr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := NewBadConfigErr(tt.args.baseError); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("NewBadConfigErr() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}