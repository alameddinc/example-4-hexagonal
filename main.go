package main

import (
	"fmt"
	"hexagonal/internal/core/services/userService"
	"hexagonal/internal/repository/inMemo"
)

func main() {
	userRepo := inMemo.NewUserInMemo()
	userService := userService.NewUserService(userRepo)

	fmt.Println(userService.Create(1, "Alper"))
	fmt.Println(userService.Create(2, "Alameddin"))
	fmt.Println(userService.Create(2, "Alameddin"))
	fmt.Println(userService.Get(1))
	fmt.Println(userService.Get(2))
	fmt.Println(userService.Get(3))
	fmt.Println("Deneme")
}
