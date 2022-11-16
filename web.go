package g

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/eoe2005/g/glog"
	"github.com/eoe2005/g/gmonitor"
	"github.com/gin-gonic/gin"
)

func RunWeb(routerRegister func(*gin.Engine)) {
	initConfig()
	// glog.RegisterErrorLog()
	r := gin.New()
	mids := []gin.HandlerFunc{
		glog.AccessLog(),
		gin.Recovery(),
	}
	r.Use(mids...)
	for _, c := range localCall {
		c()
	}
	gmonitor.RegisterGin(r)
	routerRegister(r)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := &http.Server{
		Addr:    ":8888",
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(fmt.Sprintf("listen: %v", err))
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	fmt.Fprintln(gin.DefaultWriter, "[GIN-debug] shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 10 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		panic(fmt.Sprintf("Server forced to shutdown: %v", err))
	}

	fmt.Fprintln(gin.DefaultWriter, "[GIN-debug] Server exiting")
}
