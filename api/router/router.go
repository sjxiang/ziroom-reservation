package router

import (
	"github.com/gin-gonic/gin"

	"github.com/sjxiang/ziroom-reservation/api/controller"
	"github.com/sjxiang/ziroom-reservation/pkg/mws"
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
	// auth           = app.Group("/api")
	apiv1          := engine.Group("/api/v1", mws.JWTAuthentication(r.Controller.Store.User))
	// admin          = apiv1.Group("/admin", api.AdminAuth)

	// // auth
	// auth.Post("/auth", authHandler.HandleAuthenticate)

	// user handler
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

	// bookings handlers
	// apiv1.Get("/booking/:id", bookingHandler.HandleGetBooking)
	// apiv1.Get("/booking/:id/cancel", bookingHandler.HandleCancelBooking)

	// admin handlers
	// admin.Get("/booking", bookingHandler.HandleGetBookings)
	// routerGroup := engine.Group("/api/v1")

}
