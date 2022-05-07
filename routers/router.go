package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/mamadymadan/voucher-system/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(gin.DebugMode)

	api := r.Group("/v1/vouchers")
	{
		api.POST("/", v1.AddVoucher)
		api.GET("/", v1.VoucherList)
		api.GET("/:id", v1.GetVoucher)
		api.PUT("/:id", v1.UpdateVoucher)
		api.DELETE("/:id", v1.DeleteVoucher)
	}

	return r
}
