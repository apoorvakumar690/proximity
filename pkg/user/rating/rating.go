package rating

import (
	"proximity/config"
	httpPkg "proximity/pkg/clients/http"
)

// Rater is implemented by any value that contains the required methods
type Rater interface {
	Get(GetRequest) (*GetResponse, error)
}

// NewRating creates a new instance of Rating
func NewRating(conf config.IConfig, httpRequester httpPkg.IRequest) Rater {
	return &Rating{config: conf, httpRequester: httpRequester}
}

// Rating contains methods to the perform operations on ratings
type Rating struct {
	config        config.IConfig
	httpRequester httpPkg.IRequest
}

// Get makes the request to get the ratings
func (r *Rating) Get(req GetRequest) (*GetResponse, error) {

	res := GetResponse{}

	return &res, nil
}
