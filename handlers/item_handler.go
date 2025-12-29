package handlers

import (
	"net/http"
	"strconv"
	"strings"
	
	"os"
	"fmt"

	"github.com/gin-gonic/gin"
	"go-jwt-api/config"
	"go-jwt-api/models"
	"go-jwt-api/utils"
)

// @Security BearerAuth
// @Summary Get items (pagination, search, sort, cursor)
// @Tags Item
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Limit"
// @Param sort query string false "price_asc | price_desc | name_asc | name_desc"
// @Param q query string false "Search keyword"
// @Param cursor query int false "Cursor ID"
// @Success 200 {object} map[string]interface{}
// @Router /items [get]
func GetItems(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1")) //strconv.Atoi แปลงสตริงเป็นจำนวนเต็ม
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	sort := c.DefaultQuery("sort", "id_asc")
	q := c.Query("q")
	cursor, _ := strconv.Atoi(c.DefaultQuery("cursor", "0"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	db := config.DB.Model(&models.Item{}) // เริ่มต้นการสร้างคำสั่ง SQL โดยใช้ GORM โดยระบุโมเดลที่ต้องการคือ models.Item

	// Search
	if q != "" {
		db = db.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(q)+"%")
	}

	// Cursor-based pagination
	if cursor > 0 {
		db = db.Where("id > ?", cursor)
	}

	// Sort
	switch sort {
	case "price_asc":
		db = db.Order("price asc")
	case "price_desc":
		db = db.Order("price desc")
	case "name_asc":
		db = db.Order("name asc")
	case "name_desc":
		db = db.Order("name desc")
	default:
		db = db.Order("id asc")
	}

	var total int64
	db.Count(&total)

	offset := (page - 1) * limit

	var items []models.Item
	db.Limit(limit).Offset(offset).Find(&items)

	// map → DTO
	data := make([]models.ItemDTO, 0)
	for _, item := range items {
		data = append(data, models.ItemDTO{
			ID:    item.ID,
			Name:  item.Name,
			Price: item.Price,
		})
	}

	meta := utils.BuildMeta(page, limit, total) // เรียกใช้ฟังก์ชัน BuildMeta จากแพ็กเกจ utils เพื่อสร้างข้อมูลเมตาเกี่ยวกับการแบ่งหน้า (pagination) โดยใช้ค่าของตัวแปร page, limit, total

	// Response

	c.JSON(http.StatusOK, gin.H{
		"data": data,
		"meta": meta,
	})
}


// @Security BearerAuth
// @Summary Create item
// @Tags Item
// @Accept json
// @Produce json
// @Param body body models.Item true "Item"
// @Success 200 {object} models.Item
// @Router /items/create [post]
func CreateItem(c *gin.Context) {
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil { // ShouldBindJSON แปลงข้อมูล JSON ที่ส่งมาจากคำขอ HTTP เป็นโครงสร้างข้อมูล Go
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	if item.Name == "" || item.Price <= 0 { // ตรวจสอบความถูกต้องของข้อมูลที่ได้รับมา
		c.JSON(400, gin.H{"error": "invalid item data"})
		return
	}

	if err := config.DB.Create(&item).Error; err != nil { 
		c.JSON(500, gin.H{"error": "db error"})
		return
	}

	c.JSON(201, item) // ส่งกลับข้อมูลของไอเท็มที่สร้างใหม่พร้อมกับสถานะ HTTP 201 (Created)
}


// @Security BearerAuth
// @Summary Update item
// @Tags Item
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Param body body models.Item true "Item"
// @Success 200 {array} models.ItemDTO
// @Router /items/update/{id} [put]
func UpdateItem(c *gin.Context) {
	var item models.Item
	if err := config.DB.First(&item, c.Param("id")).Error; err != nil {
		c.JSON(404, gin.H{"error": "not found"})
		return
	}

	var input models.Item
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	item.Name = input.Name
	item.Price = input.Price

	config.DB.Save(&item)
	c.JSON(200, item)
}



// UploadItemImage godoc
// @Summary Upload item image
// @Tags items
// @Security BearerAuth
// @Param id path int true "Item ID"
// @Param file formData file true "Item image"
// @Success 200 {object} map[string]string
// @Router /items/{id}/upload [post]
func UploadItemImage(c *gin.Context) {
	id := c.Param("id")

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": "file is required"})
		return
	}

	// สร้างโฟลเดอร์ uploads ถ้ายังไม่มี
	os.MkdirAll("uploads", os.ModePerm)

	filename := fmt.Sprintf("uploads/item_%s_%s", id, file.Filename) 
	//sprintf คือ การจัดรูปแบบสตริงในภาษา Go โดยใช้ฟังก์ชัน fmt.Sprintf ซึ่งจะสร้างสตริงใหม่ตามรูปแบบที่กำหนด โดยในที่นี้จะสร้างชื่อไฟล์ใหม่สำหรับการอัปโหลดภาพของไอเท็ม โดยใช้ ID ของไอเท็มและชื่อไฟล์ต้นฉบับของภาพมาเป็นส่วนประกอบในการตั้งชื่อไฟล์ใหม่

	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// update item
	config.DB.Model(&models.Item{}).
		Where("id = ?", id).
		Update("image", filename)

	c.JSON(200, gin.H{
		"message": "upload success",
		"image":   filename,
	})
}
