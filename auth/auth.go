package auth

import (
	//"github.com/dgrijalva/jwt-go"
	"log"

	//"github.com/golang-jwt/jwt"
	"github.com/golang-jwt/jwt/v5"

	"github.com/labstack/echo/v4"
	"myweb/models"
	"net/http"
	"time"
)

const (
	accessTokenCookieName  = "access-token"
	refreshTokenCookieName = "refresh-token"
	jwtSecretKey           = "some-secret-key"
	jwtRefreshSecretKey    = "some-refresh-secret-key"
)

func GetJwtSecret() string {
	return jwtSecretKey
}
func GetRefreshJwtSecret() string {
	return jwtRefreshSecretKey
}

type Claims struct {
	Name string `json:"name"`
	jwt.RegisteredClaims
}

func GenerateTokenAndSetCookie(user *models.User, c echo.Context) error {
	accessToken, t, err := generateAccessToken(user)
	if err != nil {
		return err
	}
	setTokenCookie(accessTokenCookieName, accessToken, t, c)
	setUserCookie(user, t, c)

	refreshToken, t, err := generateRefreshToken(user)
	if err != nil {
		return err
	}
	setTokenCookie(refreshTokenCookieName, refreshToken, t, c)
	return nil
}

func setUserCookie(user *models.User, t time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = "user"
	cookie.Value = user.Name
	cookie.Expires = t
	cookie.Path = "/"
	c.SetCookie(cookie)
}

func setTokenCookie(name string, token string, t time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = token
	cookie.Expires = t
	cookie.Path = "/"
	cookie.HttpOnly = true
	c.SetCookie(cookie)
}

func generateAccessToken(user *models.User) (string, time.Time, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	return generateToken(user, expirationTime, []byte(jwtSecretKey))
}

func generateRefreshToken(user *models.User) (string, time.Time, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	return generateToken(user, expirationTime, []byte(jwtRefreshSecretKey))
}

func generateToken(user *models.User, expirationTime time.Time, secret []byte) (string, time.Time, error) {

	claims := &Claims{
		Name: user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//create the string token
	tokenString, err := token.SignedString(secret)
	if err != nil {
		log.Println("ee generate token", err.Error())
		return "", time.Now(), err
	}

	return tokenString, expirationTime, nil
}

// JWT Error will be execute when user try to access a proteccted path
func JWTErrorChecked(err error, c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, c.Echo().Reverse("userSignInForm"))
}

// TokenRefresherMiddleware middleware, yang menyegarkan token JWT jika token akses akan kedaluwarsa.
func TokenRefreshMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Jika pengguna tidak diautentikasi (tidak ada data token pengguna dalam konteks), jangan lakukan apa pun.
		if c.Get("user") == nil {
			return next(c)
		}
		// Gets user token from the context.
		u := c.Get("user").(*jwt.Token)
		claims := u.Claims.(*Claims)

		// Kami memastikan bahwa token baru tidak dikeluarkan hingga waktu yang cukup berlalu.
		// Dalam hal ini, token baru hanya akan dikeluarkan jika token lama ada di dalamnya
		//if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) < 15*time.Minute {
		if claims.ExpiresAt.Sub(time.Now()) < 15*time.Minute {
			// Get the refresh toke from the cookie
			rc, err := c.Cookie(refreshTokenCookieName)
			if err == nil && rc != nil {
				tkn, err := jwt.ParseWithClaims(rc.Value, claims, func(token *jwt.Token) (interface{}, error) {
					return []byte(GetRefreshJwtSecret()), nil
				})
				if err != nil {
					if err == jwt.ErrSignatureInvalid {
						c.Response().Writer.WriteHeader(http.StatusUnauthorized)
					}
				}
				if tkn != nil && tkn.Valid {
					// If everything is good, update tokens.
					_ = GenerateTokenAndSetCookie(&models.User{
						Name: claims.Name,
					}, c)
				}
			}
		}
		return next(c)
	}
}
