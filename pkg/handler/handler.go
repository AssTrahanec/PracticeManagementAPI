package handler

import (
	"github.com/Asstrahanec/PracticeManagementAPI/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api", h.userIdentity)
	api.GET("/validateToken", h.validateToken)
	// Student routes
	student := api.Group("/students")
	{
		student.POST("/", h.createStudent)
		student.GET("/", h.getStudents)
		student.GET("/:id", h.getStudent)
		student.PUT("/:id", h.updateStudent)
		student.DELETE("/:id", h.deleteStudent)
	}

	// Company routes
	company := api.Group("/companies")
	{
		company.POST("/", h.createCompany)
		company.GET("/", h.getCompanies)
		company.GET("/:id", h.getCompany)
		company.PUT("/:id", h.updateCompany)
		company.DELETE("/:id", h.deleteCompany)
	}

	// User routes
	user := api.Group("/users")
	{
		user.POST("/", h.createUser)
		user.GET("/", h.getUsers)
		user.GET("/:id", h.getUser)
		user.PUT("/:id", h.updateUser)
		user.DELETE("/:id", h.deleteUser)
	}

	// Request routes
	request := api.Group("/requests")
	{
		request.POST("/", h.createRequest)
		request.GET("/", h.getRequests)
		request.GET("/:id", h.getRequest)
		request.PUT("/:id", h.updateRequest)
		request.DELETE("/:id", h.deleteRequest)
	}

	// Practice routes
	practice := api.Group("/practices")
	{
		practice.POST("/", h.createPractice)
		practice.GET("/", h.getPracticesOfCurrentUser)
		practice.GET("/:id", h.getPractice)
		practice.PUT("/:id", h.updatePractice)
		practice.DELETE("/:id", h.deletePractice)
		practice.GET("/all", h.getPractices)
		practice.GET("/allpublic", h.getPublicPractices)
	}

	// University routes
	university := api.Group("/universities")
	{
		university.POST("/", h.createUniversity)
		university.GET("/", h.getUniversities)
		university.GET("/:id", h.getUniversity)
		university.PUT("/:id", h.updateUniversity)
		university.DELETE("/:id", h.deleteUniversity)
	}

	return router
}
