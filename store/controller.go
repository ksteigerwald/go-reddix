package store

import (
  "encoding/json"
  //"io"
  "io/ioutil"
  "log"
  "fmt"
  "net/http"
  "strings"

  //"strconv"

  //"github.com/gorilla/mux"
  "github.com/gorilla/context"
  "github.com/dgrijalva/jwt-go"

//  "github.com/ipfs/go-ipfs-api"
)

//Controller ...
type Controller struct {
  Repository Repository
}

/* Middleware handler to handle all requests for authentication */
func AuthenticationMiddleware(next http.HandlerFunc) http.HandlerFunc {
  return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
    authorizationHeader := req.Header.Get("authorization")
    if authorizationHeader == "" {
      json.NewEncoder(w).Encode(Exception{Message: "An authorization header is required"})
      return
    }

    bearerToken := strings.Split(authorizationHeader, " ")
    if len(bearerToken) != 2 {
      return
    }

    token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
      if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
        return nil, fmt.Errorf("There was an error")
      }
      return []byte("secret"), nil
    })

    if error != nil {
      json.NewEncoder(w).Encode(Exception{Message: error.Error()})
      return
    }

    if token.Valid {
      log.Println("TOKEN WAS VALID")
      context.Set(req, "decoded", token.Claims)
      next(w, req)
    } else {
      json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization token"})
    }

  })
}

// Get Authentication token GET /
func (c *Controller) GetToken(w http.ResponseWriter, req *http.Request) {
    var user User
    _ = json.NewDecoder(req.Body).Decode(&user)
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username": user.Username,
        "password": user.Password,
    })

    log.Println("Username: " + user.Username);
    log.Println("Password: " + user.Password);

    tokenString, error := token.SignedString([]byte("secret"))
    if error != nil {
        fmt.Println(error)
    }
    json.NewEncoder(w).Encode(JwtToken{Token: tokenString})
}


func (c *Controller) CreateObj(w http.ResponseWriter, r *http.Request) {
  body, err := ioutil.ReadAll(r.Body)

  if err != nil {
    log.Fatalln("Error CreateObj", err)
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  var t Thread =*NewThread()
  //t.ID := xid.New()
  err = json.Unmarshal(body, &t)
  if err != nil {
     panic(err)
  }
  log.Println(t.ID)
  log.Println(t.Title)
}
