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
	tokoRepository := repository.NewTokoRepository(db)

	userService := service.NewUserService(userRepository)
	produkService := service.NewProdukService(produkRepository)
	tokoService := service.NewTokoService(tokoRepository)

	userHandler := handler.NewUserHandler(userService)
	produkHandler := handler.NewProdukHandler(produkService)
	tokoHandler := handler.NewTokoHandler()

	auth := r.Group("/", middleware.AuthMiddleware(userService))
	toko := auth.Group("/", middleware.AccessToko(tokoService))

	// Ambil Info User
	auth.GET("/user", userHandler.Index)

	// Membuat Toko Baru
	auth.POST("/toko", tokoHandler.Store)
	// Update Info Toko
	toko.PUT("/toko", tokoHandler.Update)
	// Hapus Toko
	toko.DELETE("/toko", tokoHandler.Destroy)

	// Ambil List Produk Semua
	auth.GET("/produk", produkHandler.Index)
	// Ambil Produk Berdasarkan ID
	auth.GET("/produk/:id", produkHandler.Show)
	// Membuat Produk Baru
	toko.POST("/produk", produkHandler.Store)
	// Mengedit Produk
	toko.PUT("/produk/:id", produkHandler.Update)
	// Menghapus Produk
	toko.DELETE("/produk/:id", produkHandler.Destroy)

	r.Run(":8080")
}
