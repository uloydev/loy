package loy

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
)

type Loy struct {
	Server *fiber.App
	Logger *zerolog.Logger
}

type Handler func(*Context) error

func New() *Loy {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	logger := zerolog.New(output).With().Timestamp().Stack().Logger()
	logger.Info().Msg("Logger initialized")
	return &Loy{
		Server: fiber.New(),
		Logger: &logger,
	}
}

func (l *Loy) Start() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	logger := l.Logger.Level(zerolog.DebugLevel)
	if os.Getenv("ENV") == "production" {
		logger = l.Logger.Level(zerolog.InfoLevel)
		l.Logger = &logger
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
		l.Logger.Info().Msgf("PORT is not set, using default %s", port)
	}
	l.Middleware(fiberzerolog.New(fiberzerolog.Config{
		Logger: l.Logger,
		Fields: []string{
			fiberzerolog.FieldMethod,
			fiberzerolog.FieldURL,
			fiberzerolog.FieldIP,
			fiberzerolog.FieldUserAgent,
			fiberzerolog.FieldReferer,
			fiberzerolog.FieldRoute,
			fiberzerolog.FieldError,
			fiberzerolog.FieldLatency,
		},
	}))
	l.Logger.Info().Msgf("Starting server on port %s", port)
	return l.Server.Listen(":" + port)
}

func (l *Loy) Middleware(handler fiber.Handler) {
	l.Server.Use(handler)
}

func (l *Loy) Add(method, path string, handle Handler, mw ...fiber.Handler) {
	handlers := mw
	handlers = append(handlers, makeHandleFunc(handle))
	l.Server.Add(method, path, handlers...)
}

func (l *Loy) Get(path string, handle Handler, mw ...fiber.Handler) {
	handlers := mw
	handlers = append(handlers, makeHandleFunc(handle))
	l.Server.Get(path, handlers...)
}

func (l *Loy) Post(path string, handle Handler, mw ...fiber.Handler) {
	handlers := mw
	handlers = append(handlers, makeHandleFunc(handle))
	l.Server.Post(path, handlers...)
}

func (l *Loy) Put(path string, handle Handler, mw ...fiber.Handler) {
	handlers := mw
	handlers = append(handlers, makeHandleFunc(handle))
	l.Server.Put(path, handlers...)
}

func (l *Loy) Delete(path string, handle Handler, mw ...fiber.Handler) {
	handlers := mw
	handlers = append(handlers, makeHandleFunc(handle))
	l.Server.Delete(path, handlers...)
}

func (l *Loy) Patch(path string, handle Handler, mw ...fiber.Handler) {
	handlers := mw
	handlers = append(handlers, makeHandleFunc(handle))
	l.Server.Patch(path, handlers...)
}

func (l *Loy) Head(path string, handle Handler, mw ...fiber.Handler) {
	handlers := mw
	handlers = append(handlers, makeHandleFunc(handle))
	l.Server.Head(path, handlers...)
}

func (l *Loy) Options(path string, handle Handler, mw ...fiber.Handler) {
	handlers := mw
	handlers = append(handlers, makeHandleFunc(handle))
	l.Server.Options(path, handlers...)
}

func makeHandleFunc(handle Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return handle(&Context{c})
	}
}
