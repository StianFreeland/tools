package kernel

import (
	"context"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"tools/services/zlog"
)

func handleSignal() {
	zlog.Warn(moduleName, "handle signal begin")
	defer wg.Done()
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	sig := <-ch
	zlog.Warn(moduleName, "handle signal", zap.Stringer("sig", sig))
	if err := srv.Shutdown(context.Background()); err != nil {
		zlog.Error(moduleName, "handle signal", zap.Error(err))
		return
	}
	zlog.Warn(moduleName, "handle signal end")
}
