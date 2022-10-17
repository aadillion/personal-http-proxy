package main

import (
	"context"
	"fmt"
	"github.com/aadillion/personal-http-proxy/handlers/proxy"
	"github.com/aadillion/personal-http-proxy/services"
	"github.com/aadillion/personal-http-proxy/services/clientsrv"
	"github.com/aadillion/personal-http-proxy/services/proxysrv"
	"github.com/aadillion/personal-http-proxy/services/storesrv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func main() {
	ctx := context.Background()
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGKILL, os.Interrupt)
	m := sync.Map{}
	storeSrv := storesrv.NewStoreSrv(&m)
	clientSrv := clientsrv.NewClientSrv(services.NewHttpClient())
	proxySrv := proxysrv.NewProxySrv(clientSrv, storeSrv)
	httpServer := http.Server{
		Addr:         ":" + strconv.Itoa(8090),
		Handler:      proxy.NewRouter(proxySrv),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	go func() {
		log.Printf("server is starting: http://localhost%s", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatal("server failed")
		}
	}()

	sig := <-signals
	fmt.Println("Got signal: ", sig)
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	_ = httpServer.Shutdown(ctx)

	os.Exit(0)
}
