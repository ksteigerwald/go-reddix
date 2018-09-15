package store
import (
  "github.com/rs/xid"
)

type User struct {
  Username string `json:"username"`
  Password string `json:"pasword"`
}

type JwtToken struct {
  Token string `json:"token"`
}

type Exception struct {
  Message string `json:"message"`
}

//Add time stamp
type Thread struct {
  ID string `json:"id"`
  Title string `json:"title"`
  Link string `json:"link"`
  Content string `json:"content"`
  Rating uint8 `json:"rating"`
}

func NewThread() *Thread {
  id := xid.New()
  return &Thread{ID: id.String() }
}
// Array of Threads
type Threads []Thread
