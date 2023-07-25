package hash

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// 哈希加密
func Make(password string) string {

	bytePassword := []byte(password)

	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}

	return string(hash)
}

// 哈希验证
func Check(hashedPassword string, password string) bool {
	byteHashedPassword := []byte(hashedPassword)
	bytePassword := []byte(password)

	err := bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)

	return err == nil
}
