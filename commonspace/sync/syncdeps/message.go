package syncdeps

import "github.com/altercpre/aparok-sync/commonspace/spacesyncproto"

type Message interface {
	ObjectId() string
	MsgSize() uint64
	ObjectType() spacesyncproto.ObjectType
}
