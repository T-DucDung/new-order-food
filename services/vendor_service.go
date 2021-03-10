package services

import (
	"new-order-food/models"
	"new-order-food/requests"
)

func GetListVendor() ([]models.Vendor, error) {
	c := models.Vendor{}
	return c.GetListVendor()
}

func UpdateVendor(vendor models.Vendor) error {
	return vendor.UpdateVendor()
}

func CreateVendor(vendor requests.RequestVendor) error {
	v := models.Vendor{
		Name:    vendor.Name,
		Address: vendor.Address,
		Email:   vendor.Email,
		Phone:   vendor.Phone,
	}
	return v.CreateVendor()
}

