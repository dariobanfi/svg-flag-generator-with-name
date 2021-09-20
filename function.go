package p

import (
	"fmt"
	"net/http"
)

func GetFlag(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}
