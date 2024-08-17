package interfaces

import (
	"golang-prisma-api/src/helper"
	"golang-prisma-api/src/utils"
)

type MetaResponse struct {
	Page       int `json:"page"`
	PerPage    int `json:"perPage"`
	TotalData  int `json:"totalData"`
	TotalPages int `json:"totalPages"`
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Error struct {
	Error ErrorResponse `json:"error"`
}

type Response[T any] struct {
	Code    *string       `json:"code"`
	Status  *int          `json:"status"`
	Message string        `json:"message"`
	Data    T             `json:"data"`
	Meta    *MetaResponse `json:"meta,omitempty"`
}

var (
	StatusInternalServerError = func(message string) Error {
		return Error{
			Error: ErrorResponse{
				Code:    utils.INTERNAL_SERVER_ERROR,
				Message: message,
			},
		}
	}
	StatusBadRequest = func(message string) Error {
		return Error{
			Error: ErrorResponse{
				Code:    utils.BAD_REQUEST,
				Message: message,
			},
		}
	}
	StatusUnauthorized = func(message string) Error {
		return Error{
			Error: ErrorResponse{
				Code:    utils.UNAUTHORIZED,
				Message: message,
			},
		}
	}
	StatusForbidden = func(message string) Error {
		return Error{
			Error: ErrorResponse{
				Code:    utils.FORBIDDEN,
				Message: message,
			},
		}
	}
	StatusNotFound = func(message string) Error {
		return Error{
			Error: ErrorResponse{
				Code:    utils.NOT_FOUND,
				Message: message,
			},
		}
	}
)

func OK[T any](message string, data T, meta *MetaResponse) Response[T] {

	return Response[T]{
		Code:    helper.ReturnPoiter[string]("OK"),
		Status:  helper.ReturnPoiter[int](200),
		Message: message,
		Data:    data,
		Meta:    meta,
	}
}
func CREATED[T any](message string, data T) Response[T] {
	return Response[T]{
		Code:    helper.ReturnPoiter[string]("CREATED"),
		Status:  helper.ReturnPoiter[int](201),
		Message: message,
		Data:    data,
	}
}
