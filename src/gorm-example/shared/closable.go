package shared

import "log"

type Closable interface {
	Close() error
}

func CloseWithErrorLogging(closeable Closable) {
	if err := closeable.Close(); err != nil {
		log.Print(err)
	}
}
