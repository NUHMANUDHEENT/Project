package routers

import (
	"project1/package/controller"
	"project1/package/handler"

	"github.com/gin-gonic/gin"
)

func UserGroup(r *gin.RouterGroup) {
	//==============user authenticatio==============
	r.GET("/user/login", controller.UserLogin)
	r.POST("/user/signup", controller.UserSignUp)
	r.POST("/user/signup/otp", controller.OtpCheck)
	r.POST("/user/signup/resend_otp", controller.ResendOtp)

	// ================= product page ===============
	r.GET("/", controller.ProductsPage)
	r.GET("/product/:ID", controller.ProductDetails)
	r.POST("/product/rating", controller.RatingStore)
	r.POST("/product/review", controller.ReviewStore)

	//============= authentication google ======================
	r.GET("/login", handler.Googlelogin)

	//================user profile=================
	r.GET("/user/profile/:ID", controller.UserProfile)
	r.POST("/user/address", controller.AddressStore)
	r.PATCH("/user/address/:ID", controller.AddressEdit)
	r.DELETE("/user/address/:ID", controller.AddressDelete)

	//================= User cart ======================
	r.POST("/cart/:ID",controller.CartStore)
	r.GET("/cart",controller.CartView)
}
