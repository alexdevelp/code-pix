package main

import (
	"os"

	"github.com/alexdevelp/code-pix/application/grpc"
	"github.com/alexdevelp/code-pix/infrastructure/db"
	"github.com/jinzhu/gorm"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))

	grpc.StartGrpcServer(database, 50051)
}
