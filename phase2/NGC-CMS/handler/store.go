package handler

import (
	"net/http"
	"ngc-cms/entity"
	"ngc-cms/helper"
	"ngc-cms/repository"
	"ngc-cms/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Handler repository.UserQuery
}

func (h UserHandler) Register(c *gin.Context) {
	// bind request json
	var reqStore entity.Store
	err := c.ShouldBindJSON(&reqStore)
	if err != nil {
		utils.ErrorMessage(c,&utils.ErrBindingJSON)
		return
	}
	
	// check if exist by email
	_, err = h.Handler.FindbyEmail(reqStore.Email)
	if err == nil {
		utils.ErrorMessage(c,&utils.ErrDuplicateData)
		return
	}
	
	// hash
	reqStore.Password,err = helper.HashPassword(reqStore.Password)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	
	// insert data
	newStore,err := h.Handler.Insert(reqStore)
	if err != nil {
		utils.ErrorMessage(c,&utils.ErrInsertData)
		return
	}

	utils.SuccessWithData(c,http.StatusCreated,newStore)
}

func (h UserHandler) ViewAll(c *gin.Context) {
	stores,err := h.Handler.FindAll()
	if err != nil {
		utils.ErrorMessage(c,&utils.ErrGetData)
		return
	}
	
	utils.SuccessWithData(c,http.StatusCreated,stores)
}

func (h UserHandler) View(c *gin.Context) {
	idStr := c.Param("id")
	id,err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorMessage(c, &utils.ErrInsertParam)
		return
	}

	store,err := h.Handler.FindbyId(id)
	if err != nil {
		utils.ErrorMessage(c,&utils.ErrGetData)
		return
	}

	utils.SuccessWithData(c,http.StatusOK,store)
}

// func (h UserHandler) Login(c *gin.Context) {
// 	// fmt.Fprint(c.Writer,"masuk login")
// 	var store entity.Store
// 	err := c.ShouldBindJSON(&store)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest,gin.H{
// 			"Message":"Failed",
// 			"Detail":err,
// 		})
// 		return
// 	}

// 	storeDb := h.Handler.FindbyEmail(store.Email)

// 	if !helper.CheckPasswordHash(store.Password,storeDb.Password){
// 		c.JSON(http.StatusBadRequest,gin.H{
// 			"Message":"Failed login",
// 			"Detail":"Wrong email or password",
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusAccepted,gin.H{
// 		"Message":"Succes Login",
// 		"Detail":storeDb.Name,
// 	})
// }
