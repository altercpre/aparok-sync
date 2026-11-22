//go:generate mockgen -destination mock_treesyncer/mock_treesyncer.go github.com/altercpre/aparok-sync/commonspace/object/treesyncer TreeSyncer
package treesyncer

import (
	"context"

	"github.com/altercpre/aparok-sync/app"
	"github.com/altercpre/aparok-sync/commonspace/object/tree/treechangeproto"
	"github.com/altercpre/aparok-sync/net/peer"
)

const CName = "common.object.treesyncer"

type TreeSyncer interface {
	app.ComponentRunnable
	StartSync()
	StopSync()
	ShouldSync(peerId string) bool
	SyncAll(ctx context.Context, p peer.Peer, existing, missing []string) error
}

// PullFilter is an optional extension of the registered TreeSyncer. When a
// head update arrives for a tree that does not exist locally, objectsync
// consults it before queueing the fetch: returning false swallows the update
// and the tree is not pulled. rootChange is the tree's raw root (carried by
// every head update; its header — including changeType — is cleartext) and
// heads are the sender's current heads. TreeSyncers that do not implement
// the interface keep the default always-pull behavior.
type PullFilter interface {
	ShouldPull(ctx context.Context, objectId string, rootChange *treechangeproto.RawTreeChangeWithId, heads []string) bool
}
