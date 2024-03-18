package main

import (
	"context"

	"github/joaltoroc/storicard/internal/app"
)

func main() {
	_ = app.NewApp(context.Background()).Run()
}
