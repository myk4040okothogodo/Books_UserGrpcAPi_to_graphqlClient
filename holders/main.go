package main

import (
    "context"
    "log"
    "os"
    "os/signal"
    "syscall"
    "time"
    "github.com/myk4040okothogodo/tutorial/db"
    "github.com/myk4040okothogodo/tutorial/holders/server"
)


func main() {
    ctx, canceFn := context.WithTimeout(context.Background(), time.Second*5)
    defer cancelFn()

    database, err := db.Connect(ctx, db.GetDbConfig())
    if err != nil {
        log.Fatalf("db.OpenDatabase failed with error: %s", err)
    }

    srv, err := server.NewServer(ctx, database)
    if err != nil {
        log.Fatalf("NewServer failed with error: %s", err)
    }

    srv.Run()

    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
    signal := <-sigChan
    log.Printf("shutting down customers servers with signal: %s", signal)
}
