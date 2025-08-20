package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// User struct maps to the 'users' table in MySQL
type User struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Name    string `json:"name"`
	Mobile  string `json:"mobile"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

// Bus struct maps to the 'buses' table in MySQL
type Bus struct {
	ID            uint    `json:"id" gorm:"primaryKey"`
	Name          string  `json:"name"`
	From          string  `json:"from"`
	To            string  `json:"to"`
	DepartureTime string  `json:"departure_time"`
	ArrivalTime   string  `json:"arrival_time"`
	Amount        float64 `json:"amount"`
}

// Booking struct maps to the 'bookings' table in MySQL
type Booking struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	UserID        uint   `json:"user_id"`
	BusID         uint   `json:"bus_id"`
	DateOfJourney string `json:"date_of_journey"`
}

// Cancellation struct maps to the 'cancellations' table in MySQL
type Cancellation struct {
	ID uint `json:"id" gorm:"primaryKey"`
}

func initDatabase() {
	// Update the DSN with your MySQL credentials
	dsn := "shreeshail:Mysql@123@tcp(127.0.0.1:3306)/ticket_booking?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	// Auto-migrate the tables
	DB.AutoMigrate(&User{}, &Bus{}, &Booking{}, &Cancellation{})
	log.Println("Database connected and tables migrated!")
}

func main() {
	initDatabase()

	router := gin.Default()

	// User Routes
	router.POST("/users", createUser)
	router.GET("/users", getAllUsers)

	// Bus Routes
	router.POST("/buses", createBus)
	router.GET("/buses", getAllBuses)

	// Booking Routes
	router.POST("/bookings", createBooking)
	router.GET("/bookings", getAllBookings)

	// Cancellation Routes
	router.POST("/cancellations", createCancellation)

	log.Println("Server running on http://localhost:8080")
	router.Run(":8080")
}

// === User Handlers ===
func createUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func getAllUsers(c *gin.Context) {
	var users []User
	if err := DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// === Bus Handlers ===
func createBus(c *gin.Context) {
	var bus Bus
	if err := c.ShouldBindJSON(&bus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := DB.Create(&bus).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create bus"})
		return
	}
	c.JSON(http.StatusCreated, bus)
}

func getAllBuses(c *gin.Context) {
	var buses []Bus
	if err := DB.Find(&buses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch buses"})
		return
	}
	c.JSON(http.StatusOK, buses)
}

// === Booking Handlers ===
func createBooking(c *gin.Context) {
	var booking Booking
	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := DB.Create(&booking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create booking"})
		return
	}
	c.JSON(http.StatusCreated, booking)
}

func getAllBookings(c *gin.Context) {
	var bookings []Booking
	if err := DB.Find(&bookings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bookings"})
		return
	}
	c.JSON(http.StatusOK, bookings)
}

// === Cancellation Handlers ===
func createCancellation(c *gin.Context) {
	var cancellation Cancellation
	if err := c.ShouldBindJSON(&cancellation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := DB.Create(&cancellation).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create cancellation"})
		return
	}
	c.JSON(http.StatusCreated, cancellation)
}
