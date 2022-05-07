package models

import (
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Voucher struct {
	gorm.Model
	ID          int64  `json:"id" gorm:"primary_key;auto_increment;not_null"`
	Code        string `json:"code"`
	MaxUses     uint   `json:"max_uses"`
	UserID      int64  `json:"user_id"`
	ExpiredTime int64  `json:"expired_time"`
}

func GetAllVouchers(vouchers *[]Voucher) (err error) {
	if err = db.Find(vouchers).Error; err != nil {
		return err
	}
	return nil
}

func GetVoucherByCode(voucher *Voucher, code string) (err error) {
	if err := db.Where("code = ?", code).First(voucher).Error; err != nil {
		return err
	}
	return nil
}

func GetVoucher(voucher *Voucher, id string) (err error) {
	if err := db.Where("id = ?", id).First(voucher).Error; err != nil {
		return err
	}
	return nil
}

func AddVoucher(voucher *Voucher) (err error) {
	userVoucher := UserVoucher{ID: int64(voucher.UserID), VoucherId: voucher.ID, Uses: 0}
	if err = AddUserVoucher(&userVoucher); err != nil {
		return err
	}
	if err = db.Model(&Voucher{}).Create(voucher).Error; err != nil {
		return err
	}
	return nil
}

func UpdateVoucher(voucher *Voucher, id string) (err error) {
	var userVoucher UserVoucher
	UpdateUserVoucher(&userVoucher, strconv.FormatInt(voucher.UserID, 10), strconv.FormatInt(voucher.ID, 10))
	db.Where("id = ?", id).Update(voucher)

	return nil
}

func DeleteVoucher(voucher *Voucher, id string) (err error) {
	db.Where("id = ?", id).Delete(voucher)
	return nil
}
