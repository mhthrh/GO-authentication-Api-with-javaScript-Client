package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mhthrh/Voiceflex/Utils/ConfigUtil"
	"github.com/mhthrh/Voiceflex/Utils/DbUtil/DbPool"
	"github.com/mhthrh/Voiceflex/Utils/LogUtil"
	"github.com/mhthrh/Voiceflex/View"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	cfg := ConfigUtil.ReadConfig("Files/ConfigCoded.json")
	if cfg == nil {
		log.Fatalln("Cant read Config,By")
	}
	logger := LogUtil.New()
	sm := mux.NewRouter()
	db := DbPool.New(&DbPool.DbInfo{
		Host:            cfg.DB[0].Host,
		Port:            cfg.DB[0].Port,
		User:            cfg.DB[0].User.UserName,
		Pass:            cfg.DB[0].User.Password,
		Dbname:          cfg.DB[0].Dbname,
		Driver:          cfg.DB[0].Driver,
		ConnectionCount: 10,
		RefreshPeriod:   20,
	})
	View.RunApiOnRouter(sm, logger, db, cfg)

	server := http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.Server.IP, cfg.Server.Port),
		Handler:      sm,
		ErrorLog:     log.New(LogUtil.LogrusErrorWriter{}, "", 0),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 20 * time.Second,
		IdleTimeout:  180 * time.Second,
	}

	go func() {
		fmt.Printf("Starting server on  %s:%d\n", cfg.Server.IP, cfg.Server.Port)
		err := server.ListenAndServe()
		if err != nil {
			logger.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	log.Println("Got signal:", <-c)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(ctx)
}
