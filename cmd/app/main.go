package main

import (
	"github.com/sergera/marketplace-order-notifier-worker/internal/evt"
)

func main() {
	listener := evt.NewOrderListener()
	listener.Listen()
}
