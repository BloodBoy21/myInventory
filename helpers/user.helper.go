package helpers

import (
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"myInventory/models"
	"myInventory/sql"
	"myInventory/utils"
	"time"
)

var db = sql.GetClient()

func ExistsUser(email string) bool {
	var user models.User
	db.Where("email = ?", email).First(&user)
	return user.ID != 0
}

func GetUserById(userId int32) (models.User, error) {
	var user models.User
	err := db.Where("id = ?", userId).First(&user).Error
	return user, err
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func CreateToken(user models.User) (string, error) {

	secret := utils.GetEnv("JWT_SECRET")
	if secret == "" {
		return "", nil
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(secret))

}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	secret := utils.GetEnv("JWT_SECRET")
	if secret == "" {
		return nil, nil
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}
	return claims, nil
}
