package v1

import (
	"github.com/mamadymadan/voucher-system/helpers"
	"github.com/mamadymadan/voucher-system/models"

	"github.com/gin-gonic/gin"
)

func CheckVoucher(c *gin.Context) {
	user_id := c.Params.ByName("user_id")
	voucher_code := c.Params.ByName(("code"))
	var voucher models.Voucher
	err := models.GetVoucherByCode(&voucher, voucher_code)
	if err != nil {
		helpers.RespondJSON(c, 404, voucher)
	}

	var userVouchers []models.UserVoucher
	err = models.GetUserVouchers(&userVouchers, user_id)
	if err != nil {
		helpers.RespondJSON(c, 404, userVouchers)
	}
	for nil, userVouvher := range userVouchers {
		if voucher.ID == int64(userVouvher.VoucherId) {
			helpers.RespondJSON(c, 200, nil)
		}
	}
	helpers.RespondJSON(c, 404, nil)
}
