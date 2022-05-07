package models

import "github.com/jinzhu/gorm"

type UserVoucher struct {
	gorm.Model
	UserID    int64 `json:"user_id" gorm:"not_null;uniqueIndex:idx_user_voucher"`
	VoucherId int64 `json:"voucher_id" gorm:"not_null"`
}
