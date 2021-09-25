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

	"github.com/joho/godotenv"
)

func Run() {
	godotenv.Load()

	migrate := flag.Bool("migrate", false, "Use For Migrate")
	seed := flag.Bool("seed", false, "Use For Seed")

	flag.Parse()

	db := Database()

	if *migrate {
		db.AutoMigrate(model.User{})
		db.AutoMigrate(model.Toko{})
		db.AutoMigrate(model.KategoriProduk{})
		db.AutoMigrate(model.Produk{})
		db.AutoMigrate(model.Transaksi{})
		db.AutoMigrate(model.DetailTransaksi{})
		fmt.Println("Berhasil Migrate Database")

		return
	}
	if *seed {
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

	r := gin.Default()

	userRepository := repository.NewUserRepository(db)
	produkRepository := repository.NewProdukRepository(db)

	userService := service.NewUserService(userRepository)
	produkService := service.NewProdukService(produkRepository)

	userHandler := handler.NewUserHandler(userService)
	produkHandler := handler.NewProdukHandler(produkService)

	auth := r.Group("/", middleware.AuthMiddleware(userService))

	auth.GET("/user", userHandler.Index)
	auth.GET("/produk", produkHandler.Index)

	r.Run(":8080")
}
