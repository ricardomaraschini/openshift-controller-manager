// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	configv1 "github.com/openshift/api/config/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// APIServerLister helps list APIServers.
// All objects returned here must be treated as read-only.
type APIServerLister interface {
	// List lists all APIServers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*configv1.APIServer, err error)
	// Get retrieves the APIServer from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*configv1.APIServer, error)
	APIServerListerExpansion
}

// aPIServerLister implements the APIServerLister interface.
type aPIServerLister struct {
	listers.ResourceIndexer[*configv1.APIServer]
}

// NewAPIServerLister returns a new APIServerLister.
func NewAPIServerLister(indexer cache.Indexer) APIServerLister {
	return &aPIServerLister{listers.New[*configv1.APIServer](indexer, configv1.Resource("apiserver"))}
}
