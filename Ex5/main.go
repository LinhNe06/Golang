package main

import (
	"fmt"
	"log"
	config "mongo-server/Config"
	routers "mongo-server/Routers"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Không tìm thấy file .env, sử dụng biến môi trường hệ thống")
	}

	err := config.ConnectMySQL()
	if err != nil {
		log.Fatalf("Khởi động thất bại: %v", err)
	}

	if err := config.ConnectMongoDB(); err != nil {
		log.Fatalf("Khởi động thất bại: %v", err)
	}

	routers.UserRouter()
	routers.OrderRouter()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "9999"
	}

	fmt.Printf("Máy chủ đang chạy tại: http://localhost:%s\n", port)
}
