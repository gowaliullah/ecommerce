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

	// repos
	productRepo := repo.NewProductRepo(*dbCon)
	userRepo := repo.NewUserRepo(dbCon)

	// domains
	userSrc := user.NewService(userRepo)
	productSrc := product.NewService(productRepo)

	middlewares := middleware.NewMiddlewares(cnf)

	// handlers
	productHdl := productHandler.NewHandler(middlewares, productSrc)
	usrHnadler := userHandler.NewHandler(cnf, userSrc)

	server := rest.NewServer(cnf, productHdl, usrHnadler)
	server.Start()
}
