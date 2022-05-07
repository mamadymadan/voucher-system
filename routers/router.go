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

	api_voucher := r.Group("/v1/vouchers")
	{
		api_voucher.POST("/", v1.AddVoucher)
		api_voucher.GET("/", v1.VoucherList)
		api_voucher.GET("/:id", v1.GetVoucher)
		api_voucher.PUT("/:id", v1.UpdateVoucher)
		api_voucher.DELETE("/:id", v1.DeleteVoucher)
	}

	return r
}
