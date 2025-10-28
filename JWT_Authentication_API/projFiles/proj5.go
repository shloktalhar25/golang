//Authentication (JWT)

//This is one of the most powerful and reusable backend skills you can learn. Once you master JWT authentication,
// you’ll be able to secure any Go + Gin API project

//Auth Service
// You’ll create an API that allows users to:

// Register (Sign Up)
// Login (Get JWT Token)
// Access Protected Routes (only if logged in)

// Concept:Description
// JWT (JSON Web Token):Secure tokens used to verify a user’s identity
// bcrypt:Used to hash and verify passwords
// Middleware:Gin’s way to check tokens before letting requests through

package main

import (
	"net/http"
	"time"

	

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5" // for JWT JWT token creation and validation
	"golang.org/x/crypto/bcrypt"   // for password hashing and checking
	"gorm.io/driver/sqlite"        // for sqlite driver
	"gorm.io/gorm"                 // for gorm (go Object Relational Mapping
)

// Model for User

// user models represents a user in our database

type User struct {
	ID string `json:"id" gorm:"primaryKey"` //auto increment primary key
	Username string `json:"username" gorm:"unique"` // unique username
	Password string `json:"-"` // exclude password from json responses
}


// global variable to hold db connection so that we can use it in handler functions

var DB *gorm.DB 


// secret key for signing JWT tokens
var jwtKey = []byte("secret_key") // secret key for signing JWT tokens (in production use env variable)


// connect to database 

func ConnectDatabase() {
	var err error
	DB , err = gorm.Open(sqlite.Open("auth.db"),&gorm.Config{})
	if err != nil {
		panic("failed to connect to database" ) //stop execution if the database fails to connect
	}
	DB.AutoMigrate(&User{}) // what it does is it checks the current state of the
	//  User model and compares it to the database schema. If there are any differences, 
	// it automatically creates or modifies the necessary tables and columns in the database to match 
	// the model definition.
}


// JWT HELPER FUNCTIONS
// this function generates a JWT token for a given username
func GenerateToken(username string)(string ,error) {
	// create a new token with claimns(data stored inside token)
	token:= jwt.NewWithClaims(jwt.SigningMethodES256,jwt.MapClaims{
		"username": username,// store username in token claims
		"exp" : time.Now().Add(time.Hour* 1).Unix(), //token expires in 1 hour.Also stores expiration time in claims
		//claims are pieces of information asserted about a subject (e.g., user) and are encoded within the JWT token.

	})

	// sign the token with the secret key and return it
	return token.SignedString(jwtKey)
}


// Register Routes

//post/register -> register a new user
func register(c *gin.Context) {
	var input User 

	// Bind the Json input to user struct(from request body) and check for errors
	// it basically stores the incoming json payload into the input struct

	if err :=c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error": "invalid input"})
		return
	}

	// hash password before storing it in DB

	hashed , err := bcrypt.GenerateFromPassword([]byte(input.Password),bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"Error":"could not hash password"})
		return
	}

	// create new user with hashed password if no errors
	// this is else block of the above if err != nil condition
	// if no error occurs we proceed to create the new user with hashed password
	user := User{
		Username: input.Username,
		Password: string(hashed), 
	}

	// now save the user to the database 
	result :=DB.Create(&user)
	// check for errors during user creation (e.g., duplicate username)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error": "username already exists"})
		return
	}
	//else return success message
	c.JSON(http.StatusBadGateway,gin.H{"message":"registration successful"})
	
}



// Post/Login 
// POST /login
func login(c *gin.Context) {
	var input User

	// Get JSON input
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	var user User
	// Search DB for username
	result := DB.Where("username = ?", input.Username).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	// Compare provided password with stored hashed password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	// Generate JWT token
	token, err := GenerateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}

	// Return token
	c.JSON(http.StatusOK, gin.H{"token": token})
}


// GET /protected
func protected(c *gin.Context) {
	// Get token from header
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
		return
	}

	// Parse and verify the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil // verify with secret key
	})

	// Check if invalid or expired
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		return
	}

	// Token valid → authorized
	c.JSON(http.StatusOK, gin.H{"message": "You are authorized!"})
}



   





func main() {
	router := gin.Default() // create Gin router with default middleware
	ConnectDatabase()       // connect and migrate DB

	// Public routes (no auth)
	router.POST("/register", register)
	router.POST("/login", login)

	// Protected route (requires JWT)
	router.GET("/protected", protected)

	// Start server on localhost:8080
	router.Run(":8080")
}


