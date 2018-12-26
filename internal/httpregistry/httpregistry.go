package httpregistry

import "github.com/sks/mqttfaas/internal/types"

//Registry ...
type Registry struct {
	store []types.HTTPRequest
}

//New ...
func New(requests []types.HTTPRequest) *Registry {
	return &Registry{
		store: requests,
	}
}

//GetAllRelavantRequests ...
func (r *Registry) GetAllRelavantRequests(topic string) []types.HTTPRequest {
	returnVal := []types.HTTPRequest{}
	for i := range r.store {
		if r.store[i].ShouldTriggerFor(topic) {
			returnVal = append(returnVal, r.store[i])
		}
	}
	return returnVal
}
