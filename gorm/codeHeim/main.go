package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	gorm.Model
	ID         uint64      `gorm:"primaryKey" json:"id"`
	Username   string      `gorm:"size:64" json:"username"`
	Password   string      `gorm:"size:255" json:"-"`
	Notes      []Note      `json:"notes,omitempty"`
	CreditCard *CreditCard `json:"credit_card,omitempty"`
}

type Note struct {
	gorm.Model
	ID      uint64 `gorm:"primaryKey" json:"id"`
	Name    string `gorm:"size:255" json:"name"`
	Content string `gorm:"type:text" json:"content"`
	UserID  uint64 `gorm:"index" json:"user_id"`
	User    User   `json:"-"`
}

type CreditCard struct {
	gorm.Model
	Number string `json:"number"`
	UserID uint64 `json:"user_id"`
	User   User   `json:"-"`
}

var DB *gorm.DB

func connectDB() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)

	dsn := "host=localhost port=5432 user=nilmadhab dbname=ecommerce sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})

	if err != nil {
		panic("Failed to connect to databse!")
	}

	DB = database
}

func dbMigrate() {
	DB.AutoMigrate(&User{}, &Note{}, &CreditCard{})
}

func seedData() {
	// Create users
	users := []User{
		{
			Username: "john_doe",
			Password: "hashed_password_123",
		},
		{
			Username: "jane_smith",
			Password: "hashed_password_456",
		},
		{
			Username: "bob_wilson",
			Password: "hashed_password_789",
		},
	}

	for i := range users {
		DB.Create(&users[i])
	}

	// Create notes for users
	notes := []Note{
		{Name: "Shopping List", Content: "Milk, Bread, Eggs", UserID: users[0].ID},
		{Name: "Meeting Notes", Content: "Discuss Q4 targets", UserID: users[0].ID},
		{Name: "Recipe", Content: "Pasta with tomato sauce", UserID: users[1].ID},
		{Name: "Travel Plans", Content: "Visit Paris in December", UserID: users[2].ID},
	}

	for _, note := range notes {
		DB.Create(&note)
	}

	// Create credit cards for users
	creditCards := []CreditCard{
		{Number: "4111-1111-1111-1111", UserID: users[0].ID},
		{Number: "5500-0000-0000-0004", UserID: users[1].ID},
	}

	for _, card := range creditCards {
		DB.Create(&card)
	}

	log.Println("Seed data created successfully!")
}

// API Handlers
func getUsers(c *gin.Context) {
	var users []User
	DB.Preload("Notes").Preload("CreditCard").Find(&users)
	c.JSON(http.StatusOK, users)
}

func getUserByID(c *gin.Context) {
	id := c.Param("id")
	var user User
	if err := DB.Preload("Notes").Preload("CreditCard").First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func getNotes(c *gin.Context) {
	var notes []Note
	DB.Find(&notes)
	c.JSON(http.StatusOK, notes)
}

type CreateUserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func createUser(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := User{Username: input.Username, Password: input.Password}
	if err := DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

type CreateNoteInput struct {
	Name    string `json:"name" binding:"required"`
	Content string `json:"content" binding:"required"`
	UserID  uint64 `json:"user_id" binding:"required"`
}

func createNote(c *gin.Context) {
	var input CreateNoteInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if user exists
	var user User
	if err := DB.First(&user, input.UserID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	note := Note{Name: input.Name, Content: input.Content, UserID: input.UserID}
	if err := DB.Create(&note).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create note"})
		return
	}

	c.JSON(http.StatusCreated, note)
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	// CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))

	// API routes
	api := r.Group("/api")
	{
		api.GET("/users", getUsers)
		api.GET("/users/:id", getUserByID)
		api.POST("/users", createUser)
		api.GET("/notes", getNotes)
		api.POST("/notes", createNote)
	}

	return r
}

func main() {
	connectDB()
	dbMigrate()

	// Check if we need to seed (only if no users exist)
	var count int64
	DB.Model(&User{}).Count(&count)
	if count == 0 {
		seedData()
	}

	// Start the server
	r := setupRouter()
	log.Println("Server starting on http://localhost:8080")
	r.Run(":8080")
}
