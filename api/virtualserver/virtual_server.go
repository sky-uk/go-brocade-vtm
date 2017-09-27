package virtualserver

import (
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/sky-uk/go-rest-api"
	"net/http"
)

// VirtualServerEndpoint : virtual server uri endpoint
var VirtualServerEndpoint = "/api/tm/3.8/config/active/virtual_servers/"

// NewCreate : used to create a new virtual server
func NewCreate(virtualServerName string, virtualServer VirtualServer) *rest.BaseAPI {
	createVirtualServerAPI := rest.NewBaseAPI(http.MethodPut, VirtualServerEndpoint+virtualServerName, virtualServer, new(VirtualServer), new(api.VTMError))
	return createVirtualServerAPI
}

// NewGetAll : used to retrieve a list of virtual servers and their HRef
func NewGetAll() *rest.BaseAPI {
	getAllVirtualServerAPI := rest.NewBaseAPI(http.MethodGet, VirtualServerEndpoint, nil, new(VirtualServersList), new(api.VTMError))
	return getAllVirtualServerAPI
}

// NewGet - used to retrieve a single virtual server
func NewGet(virtualServerName string) *rest.BaseAPI {
	getVirtualServerAPI := rest.NewBaseAPI(http.MethodGet, VirtualServerEndpoint+virtualServerName, nil, new(VirtualServer), new(api.VTMError))
	return getVirtualServerAPI
}

// NewUpdate : used to update a virtual server
func NewUpdate(name string, virtualServer VirtualServer) *rest.BaseAPI {
	updateVirtualServerAPI := rest.NewBaseAPI(http.MethodPut, VirtualServerEndpoint+name, virtualServer, new(VirtualServer), new(api.VTMError))
	return updateVirtualServerAPI
}

// NewDelete : used to delete a virtual server
func NewDelete(name string) *rest.BaseAPI {
	deleteVirtualServerAPI := rest.NewBaseAPI(http.MethodDelete, VirtualServerEndpoint+name, nil, nil, new(api.VTMError))
	return deleteVirtualServerAPI
}
