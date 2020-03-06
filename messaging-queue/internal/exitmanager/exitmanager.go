package exitmanager

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type ExitManager struct {
	done         chan struct{}
	osSignals    chan os.Signal
	serverErrors chan error
}

func Init() *ExitManager {
	e := &ExitManager{
		done:         make(chan struct{}),
		osSignals:    make(chan os.Signal, 1),
		serverErrors: make(chan error, 1),
	}
	go e.listen()
	return e
}

func (e *ExitManager) listen() {
	defer e.shutdown()
	signal.Notify(e.osSignals, syscall.SIGINT, syscall.SIGTERM)
	select {
	case sig := <-e.osSignals:
		fmt.Println(sig)
	case err := <-e.serverErrors:
		fmt.Println(err)
	}
}

func (e *ExitManager) shutdown() {
	close(e.done)
}

func (e *ExitManager) Wait() {
	<-e.done
}
func (e *ExitManager) ServerError(err error) {
	e.serverErrors <- err
}
