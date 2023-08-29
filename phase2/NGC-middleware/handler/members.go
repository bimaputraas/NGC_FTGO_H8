package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"ngc5-p2/entity"
	"ngc5-p2/helper"

	"github.com/julienschmidt/httprouter"
)

// Create a new member
func (h Handler) Register(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	// decode
	var newMember entity.Members
	err := json.NewDecoder(r.Body).Decode(&newMember)
	if err != nil {
		WriteResponse(w,400,"Invalid input")
		return
	}

	// validate
	err = helper.Validate(newMember)
	if err != nil {
		WriteResponse(w,400,err)
	}

	// hashing password
	hashedPassword,err := helper.HashPassword(newMember.Password)
	if err != nil {
		InternalErrorResponse(w,err)
		return
	}

	ctx := context.Background()
	query := `INSERT INTO members(email,password,fullname,age,occupation,role) VALUES(?,?,?,?,?,?);`
	_,err = h.HandlerDB.ExecContext(ctx,query,newMember.Email,hashedPassword,newMember.Fullname,newMember.Age,newMember.Occupation,newMember.Role)
	if err != nil {
		InternalErrorResponse(w,err)
		return
	}

	// 201 created
	WriteResponse(w,201,"New member created successfully")
}

// Login
func (h Handler) Login(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	// decode
	var requestLogin entity.Members
	err := json.NewDecoder(r.Body).Decode(&requestLogin)
	if err != nil {
		WriteResponse(w,400,"Invalid input")
		return
	}

	// get email and password from db
	ctx := context.Background()
	query := `SELECT password FROM members WHERE email = ?`
	row,err := h.HandlerDB.QueryContext(ctx,query,requestLogin.Email)
	if err != nil {
		InternalErrorResponse(w,err)
		return
	}
	if !row.Next(){
		WriteResponse(w,400,"Wrong email or password, please try again")
		return
	}

	// compare hashed password and plain password
	var plainPassword string
	err = row.Scan(&plainPassword)
	if err != nil {
		InternalErrorResponse(w,err)
		return
	}
	err = helper.CheckPasswordHash(plainPassword,requestLogin.Password)
	if err != nil {
		WriteResponse(w,400,"Wrong email or password, please try again")
		return
	}

	// 201 created
	WriteResponse(w,200,"Login success!")
}

