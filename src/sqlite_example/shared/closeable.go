package shared

import "log"

type Closable interface {
	Close() error
}

func CloseWithErrorReporting(c Closable) {
	err := c.Close()
	if err != nil {
		log.Println(err)
	}
}
