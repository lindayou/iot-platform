package define

import (
	"github.com/golang-jwt/jwt/v4"
)

type UserClaim struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Identity string `json:"identity"`
	jwt.RegisteredClaims
}

var (
	JwtKey = "iot-platform"
)

var (
	EmqxKey    = "1f9c5b734fe27865"
	EmqxSecret = "lV9C2iefOp9Cr9BeiB5rr3N9CBolJjKk3HruhqEpHQxsuVD"
)

type M map[string]interface{}
