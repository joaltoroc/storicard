package main

import (
	"context"

	"github/joaltoroc/storicard/internal/app"
)

func main() {
	app.NewApp(context.Background()).Run()
}
