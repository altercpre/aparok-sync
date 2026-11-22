package headupdater

import "github.com/altercpre/aparok-sync/commonspace/object/acl/list"

type AclUpdater interface {
	UpdateAcl(aclList list.AclList)
}
