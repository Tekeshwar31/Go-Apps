package main

import (
	"net/http"
	"os"
	"log"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	supabase "github.com/nedpals/supabase-go"
)
// Define Todo struct
type response struct {
    Status  string      `json:"status"`
    Data interface{} `json:"data"`
}

type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`  
	Phone string `json:"phone"`  
}

func main() {
	    // Load .env file
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	// Initialize Supabase client
	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")
	supabase := supabase.CreateClient(supabaseUrl,supabaseKey)

	fmt.Println("Supabase URL:", supabaseUrl)
	
	// Create router
	r := gin.Default()
	
	// Simple route
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	// Route to get all users
	r.GET("/users", func(c *gin.Context) {
		var result [] any
		err := supabase.DB.From("users").Select("*").Execute(&result)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Status": "error",
				"error": err,
			})
			return
		}
		 log.Printf("Retrieved todos: %+v", result) 
		 c.JSON(http.StatusOK, response{
			Status:  "Success",
			Data: result,
		})
	})

	// Route to create a new user
	r.POST("/users", func(c *gin.Context) {
		var newUser User
		
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Invalid request payload",
				"details": err.Error(),
			})
			return
		}
	
		// Check if email exists
		var existingUsers [] any
		err := supabase.DB.From("users").
			Select("*").
			Filter("email", "eq", newUser.Email).  // Use the email from request
			Execute(&existingUsers)
		
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Failed to check existing users",
				"details": err.Error(),
			})
			return
		}
		
		if len(existingUsers) > 0 {
			c.JSON(http.StatusConflict, gin.H{  // 409 Conflict is more appropriate
				"error": "Email already exists",
			})
			return
		}
	
		// Insert new user
		var insertedUsers []User 
		err = supabase.DB.From("users").Insert(newUser).Execute(&insertedUsers)  // Note: = instead of :=
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Failed to create user",
				"details": err.Error(),
			})
			return
		}

		if len(insertedUsers) == 0 {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "No user was created",
			})
			return
		}
	
		c.JSON(http.StatusCreated, gin.H{  // 201 Created for successful resource creation
			"message": "User created successfully",
			"user":    insertedUsers[0],  // Return the first inserted user
		})
	})

	// r.PUT("/users/:id", func(c *gin.Context) {
	// 	id := c.Param("id")
	// })

	r.Run(":8082") // Default port 8080
}

