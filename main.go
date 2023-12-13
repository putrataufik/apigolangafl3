package main

import (
	"api-golang/controllers/kamarcontroller"
	"api-golang/controllers/usercontroller"
	"api-golang/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()
	r.GET("/api/kamar", kamarcontroller.Index)
	r.GET("/api/kamar/:id", kamarcontroller.Show)
	r.PUT("/api/kamar/:nomor_kamar", kamarcontroller.Update)
	r.GET("/api/kamar/kosong", kamarcontroller.GetKamarKosong)
	r.PUT("/api/kamar/pesankamarkos/:nomorKos", kamarcontroller.PesanKamarKos)

	//user
	r.POST("/api/createuser", usercontroller.CreateUser)
	r.GET("/api/getuser/:id", usercontroller.GetUser)
	r.Run()
}
