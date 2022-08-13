package services

import (
	"new-order-food/models"
	"new-order-food/responses"
)

func GetListCategory() ([]responses.CategoryRes, error) {
	c := models.Category{}
	return c.GetListCategory()
}

func UpdateCategory(category models.Category) error {
	return category.UpdateCategory()
}

func CreateCategory(nameCategory string) error {
	c := models.Category{Name: nameCategory}
	return c.CreateCategory()
}
