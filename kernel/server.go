package kernel

import (
	"go.uber.org/zap"
	"net/http"
	"tools/services/zlog"
)

func handleServer() {
	defer wg.Done()
	zlog.Warn(moduleName, "handle server begin")
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		zlog.Error(moduleName, "handle server", zap.Error(err))
		return
	}
	zlog.Warn(moduleName, "handle server end")
}
