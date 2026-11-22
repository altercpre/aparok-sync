package metric

import (
	"context"
	"github.com/altercpre/aparok-sync/app/logger"
	"go.uber.org/zap"
	"testing"
)

func TestLog(t *testing.T) {
	m := &metric{rpcLog: logger.NewNamed("rpcLog"), appField: zap.String("appName", "test")}
	m.RequestLog(context.Background(), "rpc")
}
