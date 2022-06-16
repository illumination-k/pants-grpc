package main

import (
	"fmt"

	"github.com/illumination-k/protos/user"
)

func main() {
	user := user.UserRequest{
		Uuid: "adeka",
	}
	fmt.Println(user.Uuid)
}
