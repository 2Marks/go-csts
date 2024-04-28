package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/2marks/csts/internal/auth"
	supportrequests "github.com/2marks/csts/internal/supportRequests"
	"github.com/2marks/csts/internal/users"
	"github.com/2marks/csts/middlewares"
	"github.com/gorilla/mux"
)

type ApiServer struct {
	db   *sql.DB
	addr string
}

func NewApiServer(db *sql.DB, addr string) *ApiServer {
	return &ApiServer{
		db:   db,
		addr: addr,
	}
}

func (a *ApiServer) Run() error {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	/* start middleware declarations */
	authMiddleware := middlewares.NewAuthMiddleware(users.NewRepository(a.db))
	subRouter.Use(authMiddleware.Authenticate)
	/* end middleware declarations */

	/* start auth routes*/
	authRepository := auth.NewRepository(a.db)
	authService := auth.NewService(authRepository)
	authHandler := auth.NewHandler(authService)
	authHandler.RegisterRoutes(subRouter)
	/* end auth routes */

	/* start user routes*/
	userRepository := users.NewRepository(a.db)
	userService := users.NewService(userRepository)
	userHandler := users.NewHandler(userService)
	userHandler.RegisterRoutes(subRouter)
	/* end user routes*/

	/* start support request routes*/
	supportRequestRepository := supportrequests.NewRepository(a.db)
	supportRequestService := supportrequests.NewService(supportRequestRepository)
	supportRequestHandler := supportrequests.NewHandler(supportRequestService)
	supportRequestHandler.RegisterRoutes(subRouter)
	/* end support request routes*/

	fmt.Println("Server Listening on", a.addr)

	return http.ListenAndServe(a.addr, router)
}
