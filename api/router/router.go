package router

import (
	"github.com/gin-gonic/gin"

	"github.com/sjxiang/ziroom-reservation/api/controller"
)

type Router struct {
	Controller    *controller.Controller
// 	Authenticator *authenticator.Authenticator
}

func NewRouter(controller *controller.Controller) *Router {
	return &Router{
		Controller:    controller,
	}
}

func (r *Router) RegisterRouters(engine *gin.Engine) {
	auth           := engine.Group("/api")
	apiv1          := engine.Group("/api/v1", r.Controller.JWTAuthentication())
	admin          := apiv1.Group("/admin", r.Controller.AdminAuth())

	// auth
	auth.POST("/auth", r.Controller.Authenticate)

	// user 
	apiv1.GET("/user/:id", r.Controller.GetUserInfo)
	apiv1.PUT("/user/:id", r.Controller.UpdateUserInfo)
	apiv1.DELETE("/user/:id", r.Controller.DeleteUserInfo)
	apiv1.POST("/user", r.Controller.CreateUserInfo)
	apiv1.GET("/user", r.Controller.GetUserList)

	// hotel handlers
	// apiv1.Get("/hotel", hotelHandler.HandleGetHotels)
	// apiv1.Get("/hotel/:id", hotelHandler.HandleGetHotel)
	// apiv1.Get("/hotel/:id/rooms", hotelHandler.HandleGetRooms)

	// rooms handlers
	// apiv1.Get("/room", roomHandler.HandleGetRooms)
	// apiv1.Post("/room/:id/book", roomHandler.HandleBookRoom)
	// // TODO: cancel a booking

	// bookings
	// apiv1.Get("/booking/:id", bookingHandler.HandleGetBooking)
	// apiv1.Get("/booking/:id/cancel", bookingHandler.HandleCancelBooking)

	// admin 
	admin.GET("/booking", nil)

}
