package controllers

import (
	"net/http"

	// "../services"
	"strconv"

	"../structs"
	"github.com/gin-gonic/gin"
)

type TaxDetail struct {
	TaxType    string
	Refundable bool
	TaxPay     int
}

func (idb *InDB) GetItem(c *gin.Context) {
	var (
		item   structs.Item
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&item).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": item,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetItems(c *gin.Context) {
	var (
		items  []structs.Item
		result gin.H
	)

	idb.DB.Find(&items)
	if len(items) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": items,
			"count":  len(items),
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) CreateItem(c *gin.Context) {
	var (
		item   structs.Item
		tax    structs.Tax
		result gin.H
	)
	item.Name = c.PostForm("name")
	price, err := strconv.Atoi(c.PostForm("price"))
	code, err := strconv.Atoi(c.PostForm("tax_code"))
	if err != nil {
		result = gin.H{
			"result":  "error occured",
			"message": err,
		}
		c.JSON(http.StatusOK, result)
	}
	item.Price = price
	tax.Code = code

	idb.DB.Create(&item)
	idb.DB.Create(&tax)
	// taxDetail := services.taxService(item.price, tax.Code)

	var taxDetail TaxDetail

	if code == 1 {
		tax := price * 10 / 100
		taxDetail = TaxDetail{
			TaxType:    "food",
			Refundable: true,
			TaxPay:     tax,
		}
	}

	if code == 2 {
		tax := 10 + (price * 2 / 100)
		taxDetail = TaxDetail{
			TaxType:    "tobacco",
			Refundable: false,
			TaxPay:     tax,
		}
	}

	if code == 3 {
		tax := 0
		if price >= 100 {
			tax = (price - 100) * 1 / 100
		}
		taxDetail = TaxDetail{
			TaxType:    "entertainment",
			Refundable: false,
			TaxPay:     tax,
		}
	}

	result = gin.H{
		"name":       item.Name,
		"price":      item.Price,
		"tax_code":   tax.Code,
		"tax_detail": taxDetail,
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) UpdateItem(c *gin.Context) {
	id := c.Query("id")
	var (
		item    structs.Item
		newItem structs.Item
		result  gin.H
	)

	err := idb.DB.First(&item, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	newItem.Name = c.PostForm("name")
	newItem.Price, err = strconv.Atoi(c.PostForm("price"))
	err = idb.DB.Model(&item).Updates(newItem).Error
	if err != nil {
		result = gin.H{
			"result": "update failed",
		}
	} else {
		result = gin.H{
			"result": "successfully updated data",
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) DeleteItem(c *gin.Context) {
	var (
		item   structs.Item
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.First(&item, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	err = idb.DB.Delete(&item).Error
	if err != nil {
		result = gin.H{
			"result": "delete failed",
		}
	} else {
		result = gin.H{
			"result": "Data deleted successfully",
		}
	}

	c.JSON(http.StatusOK, result)
}
