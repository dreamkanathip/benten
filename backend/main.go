package main

import (
	"github.com/dreamkanathip/project-66/controller"
	"github.com/dreamkanathip/project-66/entity"
	"github.com/gin-gonic/gin"
)

func main() {
	entity.SetupDatabase()
	r := gin.Default()
	r.Use(CORSMiddleware())
	//Appointment Routes
		r.GET("/appointments", controller.ListAppointment)
		r.GET("/appointments/:id", controller.GetAppointment)
		r.POST("/appointments", controller.CreateAppointment)
		r.DELETE("/appointments/:id", controller.DeleteAppointment)
	//Member Routes
	r.GET("/members", controller.ListMembers)         // GET /members
	r.GET("/members/:id", controller.GetMember)       // GET /members/:id
	r.POST("/members", controller.CreateMember)        // POST /members
	r.DELETE("/members/:id", controller.DeleteMember)  // DELETE /members/:id
	//Dentist Routes
	r.GET("/dentists", controller.ListDentists)         // GET /dentists
	r.GET("/dentists/:id", controller.GetDentist)       // GET /dentists/:id
	r.POST("/dentists", controller.CreateDentist)        // POST /dentists
	r.DELETE("/dentists/:id", controller.DeleteDentist)  // DELETE /dentists/:id
	//Run at port sever
	r.Run()
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}