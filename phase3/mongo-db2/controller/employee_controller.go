package controller

import (
	"ngc2_p3/dto"
	"ngc2_p3/helper"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c Controller) AddEmployee(ctx echo.Context) error {
	// bind
	var reqBody dto.ReqBodyAddEmployee
	ctx.Bind(&reqBody)

	// validate
	err := helper.ValidateStruct(reqBody)
	if err != nil {
		return dto.ErrorResponse(400, err.Error())
	}

	// create
	newEmployee,err := c.Repository.InsertEmployee(reqBody)
	if err != nil {
		return dto.ErrorResponse(400, err.Error())	
	}

	// success
	dto.SuccessResponseResponseWithData(ctx,201,"Success add employee",newEmployee)
	return nil
}

func (c Controller) ViewEmployees(ctx echo.Context) error {
	// find many
	employees,err := c.Repository.FindEmployees()
	if err != nil {
		return dto.ErrorResponse(400, err.Error())
	}
	
	// success
	dto.SuccessResponseResponseWithData(ctx,200,"Success view all employees",employees)
	return nil
}

func (c Controller) ViewEmployee(ctx echo.Context) error {
	// get param path
	idStr := ctx.Param("id")
	idObjectId,err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return dto.ErrorResponse(400, err.Error())	
	}
	
	// find
	employee,err := c.Repository.FindEmployeById(idObjectId)
	if err != nil {
		return dto.ErrorResponse(400, err.Error())	
	}

	// success
	dto.SuccessResponseResponseWithData(ctx,200,"Success view employee",employee)
	return nil
}

func (c Controller) UpdateEmployee(ctx echo.Context) error {
	// bind
	var reqBody dto.ReqBodyUpdateEmployee
	ctx.Bind(&reqBody)

	// validate
	err := helper.ValidateStruct(reqBody)
	if err != nil {
		return dto.ErrorResponse(400, err.Error())
	}


	// get param path
	idStr := ctx.Param("id")
	idObjectId,err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return dto.ErrorResponse(400, err.Error())	
	}
	
	// find
	err = c.Repository.UpdateEmployeeById(idObjectId,reqBody)
	if err != nil {
		return dto.ErrorResponse(400, err.Error())	
	}

	// success
	dto.SuccessResponse(ctx,200,"Success update employee with _id = "+idObjectId.Hex())
	return nil
}

func (c Controller) DeleteEmployee(ctx echo.Context) error {
	// get param path
	idStr := ctx.Param("id")
	idObjectId,err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return dto.ErrorResponse(400, err.Error())	
	}
	
	// find
	err = c.Repository.DeleteEmployeeById(idObjectId)
	if err != nil {
		return dto.ErrorResponse(400, err.Error())	
	}

	// success
	dto.SuccessResponse(ctx,200,"Success update employee with _id = "+idObjectId.Hex())
	return nil
}