package utils

import (
	"net/http"
)

var (
	ErrBindingJSON = APIErrors{
		StatusCode: http.StatusBadRequest,
		ResponseCode: "0001",
		ResponseDescription:"Failed binding data"}
	ErrDuplicateData = APIErrors{
		StatusCode: http.StatusBadRequest,
		ResponseCode: "0002",
		ResponseDescription:"Data is already exist"}
	ErrInsertData = APIErrors{
		StatusCode: http.StatusBadRequest,
		ResponseCode: "0003",
		ResponseDescription:"Failed insert new data"}
	ErrGetData = APIErrors{
		StatusCode: http.StatusBadRequest,
		ResponseCode: "0004",
		ResponseDescription:"Data does not exist"}
	ErrDeleteData = APIErrors{
		StatusCode: http.StatusBadRequest,
		ResponseCode: "0005",
		ResponseDescription:"Error delete data"}
	ErrInsertParam = APIErrors{
		StatusCode: http.StatusBadRequest,
		ResponseCode: "0005",
		ResponseDescription:"Failed insert param"}
)