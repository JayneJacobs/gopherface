package common

import (
	"github.com/JayneJacobs/gopherface/common/datastore"

	"go.isomorphicgo.org/go/isokit"
)

type Env struct {
	DB          datastore.Datastore
	TemplateSet *isokit.TemplateSet
}
