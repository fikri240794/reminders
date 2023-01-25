package servers

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/rs/zerolog/log"
)

func (s *RESTfulAPIServer) Serve() {
	var (
		terminateListener   func() chan os.Signal
		serverErrListener   chan error
		shutdownErrListener chan error
		err                 error
	)

	terminateListener = func() chan os.Signal {
		var signalListener chan os.Signal = make(
			chan os.Signal,
			1,
		)

		signal.Notify(
			signalListener,
			os.Interrupt,
		)

		return signalListener
	}
	serverErrListener = make(chan error)

	go func() {
		serverErrListener <- s.server.Listen(fmt.Sprintf(":%d", s.cfg.RESTfulAPIServer.Port))
	}()

	select {
	case err = <-serverErrListener:
		if s.server != nil {
			log.Error().Msgf("Terminating %s with error: %s\n", s.cfg.RESTfulAPIServer.Name, err.Error())
			shutdownErrListener = make(chan error)

			go func() {
				shutdownErrListener <- s.server.ShutdownWithTimeout(time.Duration(s.cfg.RESTfulAPIServer.GracefullyShutdownDurationInSec) * time.Second)
			}()

			err = <-shutdownErrListener

			if err != nil {
				log.Error().Msgf("Unexpected error while terminating. error: %s\n", err.Error())
			}
		} else {
			log.Error().Msgf("Terminating %s with unexpected error: %s\n", s.cfg.RESTfulAPIServer.Name, err.Error())
		}

		log.Fatal().Msgf("%s terminated\n", s.cfg.RESTfulAPIServer.Name)
	case <-terminateListener():
		log.Warn().Msgf("Terminating %s\n", s.cfg.RESTfulAPIServer.Name)
		shutdownErrListener = make(chan error)

		go func() {
			shutdownErrListener <- s.server.ShutdownWithTimeout(time.Duration(s.cfg.RESTfulAPIServer.GracefullyShutdownDurationInSec) * time.Second)
		}()

		err = <-shutdownErrListener

		if err != nil {
			log.Error().Msgf("Unexpected error while terminating. error: %s\n", err.Error())
		}

		log.Fatal().Msgf("%s terminated\n", s.cfg.RESTfulAPIServer.Name)
	}
}
