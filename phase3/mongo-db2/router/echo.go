package router

import (
	"net/http"
	"ngc2_p3/controller"
	"ngc2_p3/repository"
	"ngc2_p3/service"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewEcho(db *mongo.Database) *echo.Echo{
	// init echo
	e := echo.New()

	// hello world
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// init repository
	repository := repository.NewRepository(db)

	// init controller
	controller := controller.NewController(repository)

	// employees group
	employee := e.Group("/employee")
	{	
		// create
		employee.POST("",controller.AddEmployee)
		// view all
		employee.GET("",controller.ViewEmployees)
		// view
		employee.GET("/:id",controller.ViewEmployee)
		// edit
		employee.PUT("/:id",controller.UpdateEmployee)
		// delete
		employee.DELETE("/:id",controller.DeleteEmployee)
	}

	// run count employees scheduler
	service.CountEmployeesScheduler(*repository)

	return e
}