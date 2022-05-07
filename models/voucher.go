package models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Voucher struct {
	gorm.Model
	ID          int64  `json:"id" gorm:"primary_key;auto_increment;not_null"`
	Code        string `json:"code"`
	Uses        uint   `json:"uses"`
	MAxUses     uint   `json:"max_uses"`
	MaxUsesUser uint   `json:"max_uses_user"`
	Type        uint   `json:"type"`
	ExpiredTime int64  `json:"expired_time"`
}

func GetAllVouchers(vouchers *[]Voucher) (err error) {
	if err = db.Find(vouchers).Error; err != nil {
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
	if err = db.Create(voucher).Error; err != nil {
		return err
	}
	return nil
}

func UpdateVoucher(voucher *Voucher, id string) (err error) {
	fmt.Println(voucher)
	db.Save(voucher)
	return nil
}

func DeleteVoucher(voucher *Voucher, id string) (err error) {
	db.Where("id = ?", id).Delete(voucher)
	return nil
}
