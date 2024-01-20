package internal

import (
	"errors"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/saufiroja/cqrs/config"
	todoRouter "github.com/saufiroja/cqrs/internal/delivery/http/router"
	"github.com/saufiroja/cqrs/pkg/database"
	"github.com/saufiroja/cqrs/pkg/logger"
	"net/http"
	"os"
)

func Start() {
	conf := config.NewAppConfig()

	// database
	db := database.NewPostgres(conf)
	log := logger.NewLogger()

	module := NewModule(db, log)

	//router
	router := httprouter.New()

	todoRouters := todoRouter.NewRouter(module, router)

	serve := &http.Server{
		Addr:    fmt.Sprintf(":%s", conf.Http.Port),
		Handler: todoRouters,
	}

	fmt.Println("--------------------")
	fmt.Println("Server running on port", conf.Http.Port)
	fmt.Println("--------------------")

	// start server
	err := serve.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		log.Error("server closed")
	} else if err != nil {
		log.Errorf("failed to start server, err: %s", err.Error())
		os.Exit(1)
	}
}
