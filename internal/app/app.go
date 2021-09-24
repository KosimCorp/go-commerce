package app

import (
	"ecommerce/internal/pkg/handler"
	"ecommerce/internal/pkg/middleware"
	"ecommerce/internal/pkg/model"
	"ecommerce/internal/pkg/repository"
	"ecommerce/internal/pkg/service"
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Run() {
	db := Database()
	r := gin.Default()

	migrate := flag.Bool("migrate", false, "Use For Migrate")

	flag.Parse()

	if *migrate {
		db.AutoMigrate(model.User{})
		db.AutoMigrate(model.Toko{})
		db.AutoMigrate(model.KategoriProduk{})
		db.AutoMigrate(model.Produk{})
		db.AutoMigrate(model.Transaksi{})
		db.AutoMigrate(model.DetailTransaksi{})

		// Seeder
		password, _ := bcrypt.GenerateFromPassword([]byte("password"), 14)
		db.Create(&model.User{
			Username: "user",
			Password: string(password),
			Role:     "customer",
			Name:     "User",
			Alamat:   "Alamat",
			Avatar:   "avatar",
			Phone:    "123123123",
		})
		fmt.Println("Berhasil Memanggil Seeder")

		return
	}

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)

	userHandler := handler.NewUserHandler(userService)

	auth := r.Group("/", middleware.AuthMiddleware(userService))

	auth.GET("/user", userHandler.Index)

	r.Run(":8080")
}
