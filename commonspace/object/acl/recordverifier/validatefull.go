package recordverifier

import (
	"github.com/altercpre/aparok-sync/consensus/consensusproto"
)

type ValidateFull struct{}

func NewValidateFull() RecordVerifier {
	return &ValidateFull{}
}

func (a *ValidateFull) VerifyAcceptor(_ *consensusproto.RawRecord) error {
	return nil
}

func (a *ValidateFull) ShouldValidate() bool {
	return true
}
