package cmd

import (
	"fmt"
	"os"

	"github.com/gowalillah/ecommerce/cart"
	"github.com/gowalillah/ecommerce/category"
	"github.com/gowalillah/ecommerce/config"
	"github.com/gowalillah/ecommerce/infra/db"
	"github.com/gowalillah/ecommerce/product"
	"github.com/gowalillah/ecommerce/repo"
	"github.com/gowalillah/ecommerce/rest"
	cartHandler "github.com/gowalillah/ecommerce/rest/handlers/cart"
	categoryHandler "github.com/gowalillah/ecommerce/rest/handlers/category"
	productHandler "github.com/gowalillah/ecommerce/rest/handlers/product"
	userHandler "github.com/gowalillah/ecommerce/rest/handlers/user"
	"github.com/gowalillah/ecommerce/rest/middleware"
	"github.com/gowalillah/ecommerce/user"
)

func Serve() {

	cnf := config.GetConfig()

	dbCon, err := db.NewConnection(cnf.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = db.MigrateDB(dbCon, "./migrations")
	if err != nil {
		fmt.Println("Error from migration", err)
		os.Exit(1)
	}

	// middlewares
	middlewares := middleware.NewMiddlewares(cnf)

	// repos
	userRepo := repo.NewUserRepo(dbCon)
	productRepo := repo.NewProductRepo(*dbCon)
	categoryRepo := repo.NewCategoryRepo(*dbCon)
	cartRepo := repo.NewCartRepo(*dbCon)

	// services
	userSrc := user.NewService(userRepo)
	productSrc := product.NewService(productRepo)
	categorySrc := category.NewService(categoryRepo)
	cartSrc := cart.NewService(cartRepo)

	// handlers
	usrHnadler := userHandler.NewHandler(middlewares, cnf, userSrc)
	catHandler := categoryHandler.NewHandler(middlewares, categorySrc)
	productHdl := productHandler.NewHandler(middlewares, productSrc)
	cartHdl := cartHandler.NewHandler(middlewares, cartSrc)

	server := rest.NewServer(cnf, usrHnadler, catHandler, productHdl, cartHdl)
	server.Start()
}
