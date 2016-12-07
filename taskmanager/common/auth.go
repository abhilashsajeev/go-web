package common
import (
    "log"
    "time"
    "io/ioutil"
    "net/http"
    jwt "github.com/dgrijalva/jwt-go"
    request "github.com/dgrijalva/jwt-go/request"
)
// using asymmetric crypto/RSA keys
const (
// openssl genrsa -out app.rsa 1024
    privKeyPath = "keys/app.rsa"
// openssl rsa -in app.rsa -pubout > app.rsa.pub
    pubKeyPath = "keys/app.rsa.pub"
)
// private key for signing and public key for verification
var (
    verifyKey, signKey []byte
)

type CustomClaims struct {
    Iss string `json:"iss"`
    userInfo customUserInfo `json:"customUserInfo"`
    jwt.StandardClaims
}

type customUserInfo struct {
    Name string
    Role string
}

// Read the key files before starting http handlers
func initKeys() {
    var err error
    signKey, err = ioutil.ReadFile(privKeyPath)
    if err != nil {
        log.Fatalf("[initKeys]: %s\n", err)
    }
    verifyKey, err = ioutil.ReadFile(pubKeyPath)
    if err != nil {
        log.Fatalf("[initKeys]: %s\n", err)
        panic(err)
    }
}

// Generate JWT token
func GenerateJWT(name, role string) (string, error) {
    expireToken:= time.Now().Add(time.Minute * 20).Unix() 
     // create claims
    claims := CustomClaims{
        "admin",
        customUserInfo {name, role},
        jwt.StandardClaims{
            ExpiresAt: expireToken,
        },
    }
    // create a signer for rsa 256
    t := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), claims)
    // set claims for JWT token
    convertedKey, perror := jwt.ParseRSAPrivateKeyFromPEM(signKey)
    if perror != nil {
        return "", perror
    }
    tokenString, err := t.SignedString(convertedKey)
    if err != nil {
        return "", err
    }
    return tokenString, nil
}


// Middleware for validating JWT tokens
func Authorize(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
    // validate the token
    token, err := request.ParseFromRequestWithClaims(r, request.AuthorizationHeaderExtractor, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
        t, rerror := jwt.ParseRSAPublicKeyFromPEM(verifyKey)
        return t, rerror
    })
    if err != nil {
        switch err.(type) {
            case *jwt.ValidationError: // JWT validation error
            vErr := err.(*jwt.ValidationError)
            switch vErr.Errors {
                case jwt.ValidationErrorExpired: //JWT expired
                DisplayAppError(
                    w,
                    err,
                    "Access Token is expired, get a new Token",
                    401,
                )
                return
                default:
                    DisplayAppError(w,
                        err,
                        "Error while parsing the Access Token!",
                        500,
                        )
                    return
            }
        default:
            DisplayAppError(w,
                err,
                "Error while parsing Access Token!",
                500)
            return
        }
    }
    if token.Valid {
        next(w, r)
    } else {
        DisplayAppError(w,
            err,
            "Invalid Access Token",
            401,
        )
    }
}