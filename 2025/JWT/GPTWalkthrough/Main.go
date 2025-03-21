package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// 2. Generate a Secret Key
// You'll need a secret key to sign and verify JWTs.
var secretKey = []byte("your-secret-key")

// 3. Create a Function to Generate JWT Tokens
// You'll issue JWTs when a user logs in or registers.
func GenerateJWT(username string) (string, error) {
	// Define the expiration time
	expirationTime := time.Now().Add(time.Hour * 24) // Token expires in 24 hours

	// Create claims with user data
	claims := jwt.MapClaims{
		"username": username,
		"exp":      expirationTime.Unix(),
	}

	// Generate the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// 4. Middleware to Verify JWT
// You'll need a middleware to protect routes.
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract token from the "Authorization" header
		authHeader := c.GetHeader("Authorization")
		println(authHeader)
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}

		// Split "Bearer <token>"
		tokenString := strings.Split(authHeader, " ")
		if len(tokenString) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		// Parse the token
		token, err := jwt.Parse(tokenString[1], func(token *jwt.Token) (interface{}, error) {
			// Ensure token is signed with HMAC method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return secretKey, nil
		})

		// Check if token is valid
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Extract claims and set them in context
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("username", claims["username"])
		}

		c.Next()
	}
}

// 5. Use Middleware to Protect Routes
// Example of applying JWT authentication to an endpoint:
func main() {
	r := gin.Default()

	// Login route (generate JWT)
	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")

		// Normally, you would authenticate user credentials from a database
		if username == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
			return
		}

		// Generate JWT token
		token, err := GenerateJWT(username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	})

	// Protected route
	authRoutes := r.Group("/protected")
	authRoutes.Use(JWTMiddleware())
	authRoutes.GET("/dashboard", func(c *gin.Context) {
		username := c.MustGet("username").(string)
		c.JSON(http.StatusOK, gin.H{"message": "Welcome, " + username})
	})

	r.Run(":8080") // Start server on port 8080
}
