package main

import (
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "go-jwt-api/docs"

	"go-jwt-api/config"
	//"go-jwt-api/models"
	"go-jwt-api/routes"
	"go-jwt-api/seed"
)

// @title Go JWT API
// @version 1.0
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	r := gin.Default()

	// Load env & connect DB
	config.LoadEnv()
	config.ConnectDB()

	// Auto migrate คือสร้างตารางในฐานข้อมูลอัตโนมัติ ตามโครงสร้างของ struct ที่กำหนดใน models
	config.AutoMigrate()

	// Seed data (dev only) 
	if os.Getenv("APP_ENV") == "dev" {
		seed.SeedItems() // เรียกใช้ฟังก์ชัน SeedItems จากแพ็กเกจ seed เพื่อเพิ่มข้อมูลตัวอย่างในตาราง items
	}

	// Routes
	routes.Setup(r) // เรียกใช้ฟังก์ชัน Setup จากแพ็กเกจ routes เพื่อกำหนดเส้นทางต่างๆ ของ API

	// Swagger
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Swagger (เปิดเฉพาะ dev)
	if os.Getenv("APP_ENV") != "prod" {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) 
		// เพิ่มเส้นทางสำหรับ Swagger UI ในสภาพแวดล้อมที่ไม่ใช่ production โดยใช้ ginSwagger.WrapHandler เพื่อให้สามารถเข้าถึงเอกสาร API ได้ผ่านทาง URL /swagger/*
	}

	// Static files (uploads)
	r.Static("/uploads", "./uploads") // ให้บริการไฟล์สถิติเพื่อให้สามารถเข้าถึงไฟล์ที่อัปโหลดผ่าน URL /uploads/*

	// Port
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	// Run server (ต้องท้ายสุด)
	r.Run(":" + port) // เริ่มต้นเซิร์ฟเวอร์ Gin บนพอร์ตที่กำหนด (ค่าพอร์ตเริ่มต้นคือ 8080 หากไม่ได้ตั้งค่าในตัวแปรสภาพแวดล้อม APP_PORT)
}

// LoadEnv โหลดตัวแปรสภาพแวดล้อมจากไฟล์ .env
// ConnectDB เชื่อมต่อกับฐานข้อมูลโดยใช้ GORM
// AutoMigrate สร้างตารางในฐานข้อมูลตามโครงสร้างของ struct ใน models
// SeedItems เพิ่มข้อมูลตัวอย่างในตาราง items (สำหรับสภาพแวดล้อมการพัฒนาเท่านั้น)
// Setup กำหนดเส้นทางต่างๆ ของ API
// ginSwagger.WrapHandler ให้บริการ Swagger UI สำหรับเอกสาร API
// r.Static ให้บริการไฟล์สถิติเพื่อเข้าถึงไฟล์ที่อัปโหลด
// r.Run เริ่มต้นเซิร์ฟเวอร์ Gin บนพอร์ตที่กำหนด

//ไปอ่านต่อการทำงานของแต่ละฟังก์ชันได้ที่ไฟล์อื่นๆ ในโปรเจ็กต์นี้ เช่น config/env.go, config/db.go, routes/routes.go เป็นต้น