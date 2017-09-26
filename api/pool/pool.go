package pool

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/sky-uk/go-rest-api"
	"net/http"
)

// PoolEndpoint : pool uri endpoint
var PoolEndpoint = "/api/tm/3.8/config/active/pools/"

// NewCreate : used to create a new pool
func NewCreate(poolName string, pool Pool) *rest.BaseAPI {
	createPoolAPI := rest.NewBaseAPI(http.MethodPut, PoolEndpoint+poolName, pool, new(Pool), new(api.VTMError))
	return createPoolAPI
}

// NewGetAll - Returns all pools
func NewGetAll() *rest.BaseAPI {
	getAllPoolsAPI := rest.NewBaseAPI(http.MethodGet, PoolEndpoint, nil, new(LBPoolList), new(api.VTMError))
	return getAllPoolsAPI
}

// NewGet - Returns a single pool
func NewGet(poolName string) *rest.BaseAPI {
	getPoolAPI := rest.NewBaseAPI(http.MethodGet, PoolEndpoint+poolName, nil, new(Pool), new(api.VTMError))
	return getPoolAPI
}

// NewUpdate : used to update an existing pool
func NewUpdate(poolName string, pool Pool) *rest.BaseAPI {
	updatePoolAPI := rest.NewBaseAPI(http.MethodPut, PoolEndpoint+poolName, pool, new(Pool), new(api.VTMError))
	return updatePoolAPI
}

// NewDelete - used to delete a pool
func NewDelete(poolName string) *rest.BaseAPI {
	deletePoolAPI := rest.NewBaseAPI(http.MethodDelete, PoolEndpoint+poolName, nil, new(Pool), new(api.VTMError))
	return deletePoolAPI
}
