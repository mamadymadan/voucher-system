package v1

import (
	"github.com/mamadymadan/voucher-system/helpers"
	"github.com/mamadymadan/voucher-system/models"

	"github.com/gin-gonic/gin"
)

func VoucherList(c *gin.Context) {
	var vouchers []models.Voucher
	err := models.GetAllVouchers(&vouchers)
	if err != nil {
		helpers.RespondJSON(c, 404, vouchers)
	} else {
		helpers.RespondJSON(c, 200, vouchers)
	}
}

func AddVoucher(c *gin.Context) {
	var voucher models.Voucher
	c.BindJSON(&voucher)
	err := models.AddVoucher(&voucher)
	if err != nil {
		helpers.RespondJSON(c, 404, voucher)
	} else {
		helpers.RespondJSON(c, 200, voucher)
	}
}

func GetVoucher(c *gin.Context) {
	id := c.Params.ByName("id")
	var voucher models.Voucher
	err := models.GetVoucher(&voucher, id)
	if err != nil {
		helpers.RespondJSON(c, 404, voucher)
	} else {
		helpers.RespondJSON(c, 200, voucher)
	}
}

func UpdateVoucher(c *gin.Context) {
	var voucher models.Voucher
	id := c.Params.ByName("id")
	err := models.GetVoucher(&voucher, id)
	if err != nil {
		helpers.RespondJSON(c, 404, voucher)
	}
	c.BindJSON(&voucher)
	err = models.UpdateVoucher(&voucher, id)
	if err != nil {
		helpers.RespondJSON(c, 404, voucher)
	} else {
		helpers.RespondJSON(c, 200, voucher)
	}
}

func DeleteVoucher(c *gin.Context) {
	var voucher models.Voucher
	id := c.Params.ByName("id")
	err := models.DeleteVoucher(&voucher, id)
	if err != nil {
		helpers.RespondJSON(c, 404, voucher)
	} else {
		helpers.RespondJSON(c, 200, voucher)
	}
}
