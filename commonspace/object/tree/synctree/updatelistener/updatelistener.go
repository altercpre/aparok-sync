//go:generate mockgen -destination mock_updatelistener/mock_updatelistener.go github.com/altercpre/aparok-sync/commonspace/object/tree/synctree/updatelistener UpdateListener
package updatelistener

import (
	"github.com/altercpre/aparok-sync/commonspace/object/tree/objecttree"
)

type UpdateListener interface {
	Update(tree objecttree.ObjectTree) error
	Rebuild(tree objecttree.ObjectTree) error
}
