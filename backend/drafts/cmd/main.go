package main

import (
	"log/slog"
	"os"
	"os/signal"

	"github.com/davg/drafts/internal/aplication"
)

func main() {
	h := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	mainlogger := slog.New(h)

	app := aplication.New(mainlogger)

	app.Start()

	ch := make(chan os.Signal, 1)

	signal.Notify(ch, os.Interrupt)

	<-ch

	app.GracefulStop()
}
