package app

import (
	"fmt"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"log"
	"os"
	"os/signal"
	"syscall"
	"test/project/internal/handlers/manager"
	"test/project/pkg/config"
	"test/project/pkg/scyllaDb"
	"time"
)

func InitApp(cfg *config.Configs) error {
	session, err := scyllaDb.NewScyllaDB(cfg)

	if err != nil {
		return err
	}

	defer func() {
		session.Close()
	}()

	app := manager.Manager(session, cfg)

	go func() {
		if err := app.Listen(fmt.Sprintf(":%s", cfg.Listen.Port)); err != nil {
			panic(err)
		}
	}()

	app.Get("/swagger/*", swagger.HandlerDefault)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)
	<-signals
	log.Println("Shutdown Server ...")

	if err := app.ShutdownWithTimeout(5 * time.Second); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	return nil
}
