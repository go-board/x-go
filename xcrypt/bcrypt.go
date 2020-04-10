package xcrypt

import (
	"golang.org/x/crypto/bcrypt"
)

var bcryptCost = bcrypt.DefaultCost

// SetBCryptCost set global bcrypt cost, and after, all encrypt operation
// use this call before changed again
func SetBCryptCost(cost int) {
	if cost >= bcrypt.MinCost && cost <= bcrypt.MaxCost {
		bcryptCost = cost
	}
}

// BCryptHash encrypt original password with bcryptCost to encrypted data,
// then store encrypted data into database.
func BCryptHash(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, bcryptCost)
}

// BCryptValidate validate password and encryptedData retrieved from database.
func BCryptValidate(password []byte, encryptedData []byte) bool {
	return bcrypt.CompareHashAndPassword(encryptedData, password) == nil
}
