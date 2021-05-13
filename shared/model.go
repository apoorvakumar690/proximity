package shared

import (
	"proximity/config"
	"proximity/pkg/clients/db"
	grpcPkg "proximity/pkg/clients/grpc"
	httpPkg "proximity/pkg/clients/http"
)

// VERSION keeps the version no. (commit id) for global use
var VERSION string

// Deps ... is a shared dependencies struct that contains common singletons
type Deps struct {
	Config        config.IConfig
	Database      *db.Instances
	GrpcConn      grpcPkg.IGrpcConnections
	HTTPRequester httpPkg.IRequest
}
