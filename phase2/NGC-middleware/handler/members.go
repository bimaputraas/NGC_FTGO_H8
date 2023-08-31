package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"ngc5-p2/entity"
	"ngc5-p2/helper"

	"github.com/golang-jwt/jwt/v5"
	"github.com/julienschmidt/httprouter"
)
type MemberHandler struct {
	HandlerDB *sql.DB
}

// Create a new member
func (h MemberHandler) Register(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	// decode
	var newMember entity.Members
	json.NewDecoder(r.Body).Decode(&newMember)
	
	// check email if already exist
	ctx := context.Background()
	query1 := `SELECT * FROM members WHERE email = ?`
	row,err := h.HandlerDB.QueryContext(ctx,query1,newMember.Email)
	if err != nil {
		InternalErrorResponse(w,err)
		return
	}
	if row.Next(){
		WriteResponse(w,400,"Email already exist, please try again")
		return
	}

	// validate
	err = helper.Validate(newMember)
	if err != nil {
		WriteResponse(w,400,err.Error())
		return
	}

	// hashing password
	hashedPassword,err := helper.HashPassword(newMember.Password)
	if err != nil {
		InternalErrorResponse(w,err)
		return
	}

	query2 := `INSERT INTO members(email,password,fullname,age,occupation,role) VALUES(?,?,?,?,?,?);`
	_,err = h.HandlerDB.ExecContext(ctx,query2,newMember.Email,hashedPassword,newMember.Fullname,newMember.Age,newMember.Occupation,newMember.Role)
	if err != nil {
		InternalErrorResponse(w,err)
		return
	}

	// 201 created
	WriteResponse(w,201,"New member created successfully")
}

// Login
func (h MemberHandler) Login(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	// decode
	var requestLogin entity.Members
	err := json.NewDecoder(r.Body).Decode(&requestLogin)
	if err != nil {
		WriteResponse(w,400,"Invalid input")
		return
	}

	// get data from db
	ctx := context.Background()
	query := `SELECT id,password FROM members WHERE email = ?`
	row,err := h.HandlerDB.QueryContext(ctx,query,requestLogin.Email)
	if err != nil {
		InternalErrorResponse(w,err)
		return
	}
	if !row.Next(){
		WriteResponse(w,400,"Wrong email or password, please try again ")
		return
	}

	// get id and password member from DB
	var member entity.Members
	err = row.Scan(&member.Id,&member.Password)
	if err != nil {
		InternalErrorResponse(w,err)
		return
	}

	// compare hashed password and plain password
	err = helper.CheckPasswordHash(requestLogin.Password,member.Password)
	if err != nil {
		WriteResponse(w,400,"Wrong email or password, please try again ")
		return
	}

	// generate token
	token,err := helper.GenerateToken(jwt.MapClaims{
		"userId": member.Id,
	})
	if err != nil {
		InternalErrorResponse(w,err)
		return
	}

	// 201 created
	WriteResponse(w,200,map[string]string{
		"Message": "Login Success",
		"Token" : token,
	})
}

