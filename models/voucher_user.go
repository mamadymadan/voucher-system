package models

import "github.com/jinzhu/gorm"

type UserVoucher struct {
	gorm.Model
	ID        int64 `json:"id"`
	VoucherId int64 `json:"voucher_id" gorm:"not_null"`
	Uses      int64 `json:"uses"`
}

func AddUserVoucher(userVoucher *UserVoucher) (err error) {
	if err = db.Create(userVoucher).Error; err != nil {
		return err
	}
	return nil
}

func GetUserVoucher(userVoucher *UserVoucher, id string, voucherId string) (err error) {
	if err := db.Where("id = ? AND voucher_id = ?", id, voucherId).First(userVoucher).Error; err != nil {
		return err
	}
	return nil
}

func GetUserVouchers(userVouchers *[]UserVoucher, id string) (err error) {
	if err := db.Where("id = ?", id).Find(userVouchers).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUserVoucher(userVoucher *UserVoucher, id string, voucherId string) (err error) {
	db.Where("id = ? AND voucher_id = ?", id, voucherId).Delete(userVoucher)
	return nil
}

func UpdateUserVoucher(userVoucher *UserVoucher, id string, voucherId string) (err error) {
	db.Where("id = ? AND voucher_id = ?", id, voucherId).Update(userVoucher)
	return nil
}
