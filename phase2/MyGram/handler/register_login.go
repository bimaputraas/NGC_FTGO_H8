package handler

import (
	"mygram/helpers"
	"mygram/models"
	"mygram/repository"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Repository repository.UserRepository
}

func (h UserHandler) Register(c *gin.Context) {
	// bind
	var reqRegister models.User
	err := c.ShouldBindJSON(&reqRegister)
	if err != nil {
		// panic(err)
		helpers.ResponseWritter(c,404,err.Error())
		return
	}
	// hash
	reqRegister.Password,err = helpers.HashPassword(reqRegister.Password)
	if err != nil {
		// panic(err)
		helpers.ResponseWritter(c,500,"Failed hash")
		return
	}
	
	// query insert
	user,err := h.Repository.Insert(reqRegister)
	if err != nil {
		// panic(err)
		helpers.ResponseWritter(c,500,err.Error())
		return
	}

	helpers.ResponseWritterWithData(c, 201, "Success Register", user)
}


// func Authentication() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 	 verifyToken, err := helpers.VerifyToken(c)
// 	 _ = verifyToken
// 	 // return
   
// 	 if err != nil {
// 	  c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
// 	   "error":   "Unauthenticated",
// 	   "message": err.Error(),
// 	  })
// 	  return
// 	 }
// 	 c.Set("userData", verifyToken)
// 	 c.Next()
// 	}
//    }

// func VerifyToken(c *gin.Context) (interface{}, error) {
// 	errReponse := errors.New("sign in to proceed")
// 	headerToken := c.Request.Header.Get("Authorization")
// 	bearer := strings.HasPrefix(headerToken, "Bearer")
// 	if !bearer {
// 	 return nil, errReponse
// 	}
// 	stringToken := strings.Split(headerToken, " ")[1]
   
// 	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
// 	 if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
// 	  return nil, errReponse
// 	 }
// 	 return []byte(secretKey), nil
// 	})
   
// 	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
// 	 return nil, errReponse
// 	}
   
// 	return token.Claims.(jwt.MapClaims), nil
//    }