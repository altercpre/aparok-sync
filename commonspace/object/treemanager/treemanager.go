//go:generate mockgen -destination mock_treemanager/mock_treemanager.go github.com/altercpre/aparok-sync/commonspace/object/treemanager TreeManager
package treemanager

import (
	"context"

	"github.com/altercpre/aparok-sync/app"
	"github.com/altercpre/aparok-sync/commonspace/object/tree/objecttree"
	"github.com/altercpre/aparok-sync/commonspace/object/tree/treestorage"
)

const CName = "common.object.treemanager"

type TreeManager interface {
	app.ComponentRunnable
	GetTree(ctx context.Context, spaceId, treeId string) (objecttree.ObjectTree, error)
	ValidateAndPutTree(ctx context.Context, spaceId string, payload treestorage.TreeStorageCreatePayload) error
	MarkTreeDeleted(ctx context.Context, spaceId, treeId string) error
	DeleteTree(ctx context.Context, spaceId, treeId string) error
}
