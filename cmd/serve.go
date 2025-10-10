package cmd

import (
	"fmt"
	"os"

	"github.com/gowalillah/ecommerce/config"
	"github.com/gowalillah/ecommerce/infra/db"
	"github.com/gowalillah/ecommerce/product"
	"github.com/gowalillah/ecommerce/repo"
	"github.com/gowalillah/ecommerce/rest"
	productHandler "github.com/gowalillah/ecommerce/rest/handlers/product"
	"github.com/gowalillah/ecommerce/rest/middleware"
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

	// repos
	productRepo := repo.NewProductRepo(*dbCon)
	userRepo := repo.NewUserRepo(*dbCon)

	// domains
	productSrc := product.NewService(productRepo)
	userSrc := user.NewService(userRepo)

	middlewares := middleware.NewMiddlewares(cnf)

	// handlers
	productHdl := productHandler.NewHandler(middlewares, productSrc)

	server := rest.NewServer(cnf, userSrc, productHdl)

	server.Start()

}
