package main

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func logSomething() {
	log.Debug().
		Str("Scale", "833 cents").
		Float64("Interval", 833.09).
		Msg("Fibonacci is everywhere")
	log.Warn().
		Str("Name", "Tom").
		Send()
	fmt.Println("")
}

func main() {
	zerolog.TimestampFunc = func() time.Time {
		return time.Now().In(time.Local)
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "2006-01-02 15:04:05 Mon"})
	logSomething()

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "2006-01-02 15:04:05 Mon"}).With().Caller().Logger()
	logSomething()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	logSomething()
}
