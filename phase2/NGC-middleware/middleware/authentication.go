package middleware

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"ngc5-p2/entity"
	"ngc5-p2/handler"
	"ngc5-p2/helper"

	"github.com/julienschmidt/httprouter"
)

func Authentication(next httprouter.Handle, db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Println("MASUK AUTHENTICATION")
		// cek token nya ada apa ga?
		token := r.Header.Get("access-token")

		// token valid apa ga?
		claims, err := helper.DecodeToken(token)
		if err != nil {
			handler.WriteResponse(w, http.StatusUnauthorized,"invalid token, please provide valid access token")
			return
		}

		// cek informasi dalam token nya beneran ada apa ga?
		userId := int(claims["userId"].(float64))
		loggedInUser, err := findUserById(userId, db)
		if err != nil {
			handler.WriteResponse(w, http.StatusUnauthorized,"invalid token, please provide valid access token")
			return
		}

		ctx := context.WithValue(r.Context(), "loggedInUser", loggedInUser)

		next(w, r.WithContext(ctx), p)
	}
}

func AuthorizeSuperAdmin(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Println("AUTHORIZATION")
		userVal := r.Context().Value("loggedInUser")
		loggedInUser := userVal.(entity.Members)

		if loggedInUser.Role != "superadmin" {
			handler.WriteResponse(w, http.StatusForbidden,"forbidden action. only superadmin can access this endpoint")
			return
		}

		next(w, r, p)
	}
}

func findUserById(id int, db *sql.DB) (entity.Members, error) {
	query := `
		SELECT * FROM members WHERE id = ?
	`
	rows, err := db.QueryContext(context.Background(), query, id)
	if err != nil {
		panic(err)
	}

	userFound := entity.Members{}

	if !rows.Next() {
		return userFound, errors.New("user not found")
	}

	rows.Scan(&userFound.Id, &userFound.Email, &userFound.Password, &userFound.Fullname,&userFound.Age,&userFound.Occupation,&userFound.Role)
	return userFound, err
}