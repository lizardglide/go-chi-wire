package main

import (
	"go-chi-wire/internal/di"
	"log"
	"net/http"
)

func main() {
	userRouter := di.InitializeRouter()
	//userRouter := router.NewUserRouter(userHandler)
	log.Println("UserRouter initialized and running at port :8080")
	if err := http.ListenAndServe(":8080", userRouter); err != nil {
		log.Fatalf("Failed to start router: %v", err)
	}
}
