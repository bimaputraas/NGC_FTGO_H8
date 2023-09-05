package handler

import (
	"fmt"
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

// @BasePath /api/v1

// Register User godoc
// @Summary Register user
// @Description do Register user
// @ID Create-users
// @Accept json
// @Produce json
// @Created 201 {object} utils.SuccessWithData
// @InternalError 500 {object} utils.APIErrors
// @BindingError 400 {object} utils.APIErrors
// @DuplicateError 400 {object} utils.APIErrors
// @InsertError 400 {object} utils.APIErrors
// @GetError 400 {object} utils.APIErrors
// @Router /users/register/ [post]
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

// View all User godoc
// @Summary View all user
// @Description do view all user
// @ID Viewall-users
// @Accept json
// @Produce json
// @Success 200 {object} utils.SuccessWithData
// @InternalError 500 {object} utils.APIErrors
// @GetError 400 {object} utils.APIErrors
// @Router /users/ [get]
func (h UserHandler) ViewAll(c *gin.Context) {
	stores,err := h.Handler.FindAll()
	if err != nil {
		utils.ErrorMessage(c,&utils.ErrGetData)
		return
	}
	
	utils.SuccessWithData(c,http.StatusCreated,stores)
}

// View all user by id godoc
// @Summary View all userby id 
// @Description do view all user by id
// @ID Viewall-users
// @Accept json
// @Produce json
// @Param id path int true "Users ID"
// @Success 200 {object} utils.SuccessWithData
// @InternalError 500 {object} utils.APIErrors
// @GetError 400 {object} utils.APIErrors
// @Router /users/{id} [get]
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

func (h UserHandler) Login(c *gin.Context) {
	var logStore entity.Store
	err := c.ShouldBindJSON(&logStore)
	if err != nil {
		utils.ErrorMessage(c,&utils.ErrBindingJSON)
		return
	}
	
	dbStore,err := h.Handler.FindbyEmail(logStore.Email)
	if err != nil {
		utils.ErrorMessage(c,&utils.ErrGetData)
	}

	if !helper.CheckPasswordHash(logStore.Password,dbStore.Password){
		fmt.Fprint(c.Writer, "Wrong password or email")
	}

	utils.SuccessWithData(c,http.StatusOK,"Success Login")

}

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
