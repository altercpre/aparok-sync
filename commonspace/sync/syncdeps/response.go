package syncdeps

import (
	"github.com/altercpre/aparok-sync/util/multiqueue"
)

type Response interface {
	multiqueue.Sizeable
}
