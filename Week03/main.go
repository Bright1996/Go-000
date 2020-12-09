package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// 问题 : 基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够 一个退出，全部注销退出。
func main() {
	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		return SignalServer(ctx)
	})

	g.Go(func() error {
		return HttpServer(ctx, ":8080", nil)
	})
	g.Go(func() error {
		return HttpServer(ctx, ":8081", nil)
	})
	g.Go(func() error {
		return HttpServer(ctx, ":8082", nil)
	})

	fmt.Println(g.Wait())
}

func HttpServer(ctx context.Context, addr string, handler http.Handler) error {
	s := http.Server{
		Addr:    addr,
		Handler: handler,
	}

	go func() {
		<-ctx.Done()
		_ = s.Shutdown(ctx)
	}()

	err := s.ListenAndServe()
	return err
}

func SignalServer(ctx context.Context) error {
	osChan := make(chan os.Signal, 1)
	signal.Notify(osChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-ctx.Done():
	case call := <-osChan:
		return errors.New(fmt.Sprintf("os exit syscall : %v", call))
	}

	return nil
}
