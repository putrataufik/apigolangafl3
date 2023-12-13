package kamarcontroller

import (
	"api-golang/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var kamar []models.Kamar
	models.DB.Find(&kamar)

	c.JSON(http.StatusOK, gin.H{"kamar": kamar})
}

func Show(c *gin.Context) {
	var kamar models.Kamar
	id := c.Param("id")

	if err := models.DB.First(&kamar, id).Error; err != nil {

		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data Tidak Ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"kamar": kamar})
}

func Update(c *gin.Context) {
	var kamar models.Kamar

	nomorKamar := c.Param("nomor_kamar")

	if err := c.ShouldBindJSON(&kamar); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&kamar).Where("nomor_kamar = ?", nomorKamar).Updates(&kamar).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak Dapat Update Kamar"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"kamar": "Kamar Berhasil TerUpdate"})

}

func GetKamarKosong(c *gin.Context) {
	var kamars []models.Kamar

	if err := models.DB.Select("nomor_kamar, status_kamar").Where("status_kamar = ?", "Kosong").Find(&kamars).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	var kamarInfoList []models.KamarInfo
	for _, k := range kamars {
		kamarInfoList = append(kamarInfoList, models.KamarInfo{
			NomorKamar: k.NomorKamar,
			Status:     k.StatusKamar,
		})
	}

	if len(kamarInfoList) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Kamar Kosong Tidak Ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"kamar_kosong": kamarInfoList})
}

func PesanKamarKos(c *gin.Context) {
	var pesanRequest models.PesanKamarRequest
	nomorKos := c.Param("nomorKos")

	if err := c.ShouldBindJSON(&pesanRequest); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Incorrect type of input on field"})
		return
	}

	var kamar models.Kamar
	result := models.DB.Where("nomor_kamar = ?", nomorKos).First(&kamar)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "nomor kos not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": result.Error.Error()})
		return
	}
	kamar.StatusKamar = pesanRequest.StatusKos
	kamar.PemilikKamar = pesanRequest.NamaPemesan
	kamar.WaktuMasuk = pesanRequest.WaktuMasuk
	kamar.WaktuKeluar = pesanRequest.WaktuKeluar

	if err := models.DB.Save(&kamar).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"nomor_kamar": nomorKos,
		"statusKos":   pesanRequest.StatusKos,
		"namaPemesan": pesanRequest.NamaPemesan,
		"waktuMasuk":  pesanRequest.WaktuMasuk,
		"waktuKeluar": pesanRequest.WaktuKeluar,
		"message":     "Successful",
	})
}
