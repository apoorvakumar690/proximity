package http

import (
	"sync"

	"proximity/apis/http/ping"
	"proximity/apis/middleware"
	log "proximity/pkg/utils/logger"
	"proximity/shared"

	"github.com/gin-gonic/gin"
)

// StartServer starts the http server using the dependencies passed to it.
// It also initializes the routes
func StartServer(deps *shared.Deps, wg *sync.WaitGroup, fatalError chan error) error {
	address := deps.Config.Get().Server.HTTP.Address

	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	// Injects apm to trace http requests in gin
	// router.Use(middleware.ApmMiddleware(apm.APM))
	// Adds panic handler as a middleware
	router.Use(middleware.HandlePanic)

	// Initializes Ping routes
	ping.NewPingRoute(router)
	// Initialize all the routes
	// httpUser.NewUserRoute(router, deps)

	log.Debug("HTTP Server listening on : " + address + ", Version: " + shared.VERSION)

	// Start the server
	err := router.Run(address)
	if err != nil {
		fatalError <- err
	}

	// Go routine finished (can also be deferred)
	wg.Done()

	return nil
}
