package controllers

import (
	"hanifanggit/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type PersonInput struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Phonenumber string `json:"phonenumber"`
	City        string `json:"city"`
	Address     string `json:"address"`
	Province    string `json:"province"`
}

//Tampil data Person
func PersonTampil(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var mhs []models.Person
	db.Find(&mhs)
	c.JSON(http.StatusOK, gin.H{"data": mhs})
}

//Tambah Data Person
func PersonTambah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//validasi input/masukkan
	var dataInput PersonInput
	if err := c.ShouldBindJSON(&dataInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//proses input
	mhs := models.Person{
		Name:        dataInput.Name,
		Phonenumber: dataInput.Phonenumber,
		City:        dataInput.City,
		Address:     dataInput.Address,
		Province:    dataInput.Province,
	}

	db.Create(&mhs)

	c.JSON(http.StatusOK, gin.H{"data": mhs})
}

//Ubah Data Person
func PersonUbah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//cek dulu datanya
	var mhs models.Person
	if err := db.Where("id = ?", c.Param("id")).First(&mhs).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data Person tidak ditemukan"})
		return
	}

	//validasi input/masukkan
	var dataInput PersonInput
	if err := c.ShouldBindJSON(&dataInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//proses ubah data
	db.Model(&mhs).Update(dataInput)

	c.JSON(http.StatusOK, gin.H{"data": mhs})
}

//Hapus Data Person
func PersonHapus(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//cek dulu datanya
	var mhs models.Person
	if err := db.Where("id = ?", c.Param("id")).First(&mhs).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data Person tidak ditemukan"})
		return
	}

	//proses hapus data
	db.Delete(&mhs)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
