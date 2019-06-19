package main

import (
	"context"
	"math"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/dchest/uniuri"
	"github.com/gomematic/gomematic-api/pkg/config"
	"github.com/gomematic/gomematic-api/pkg/middleware/requestid"
	"github.com/gomematic/gomematic-api/pkg/router"
	"github.com/gomematic/gomematic-api/pkg/service"
	"github.com/gomematic/gomematic-api/pkg/service/teams"
	"github.com/gomematic/gomematic-api/pkg/service/users"
	"github.com/oklog/oklog/pkg/group"
	"github.com/rs/zerolog/log"
	"gopkg.in/urfave/cli.v2"
)

// Server provides the sub-command to start the server.
func Server(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:   "server",
		Usage:  "start integrated server",
		Flags:  serverFlags(cfg),
		Before: serverBefore(cfg),
		Action: serverAction(cfg),
	}
}

func serverFlags(cfg *config.Config) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "metrics-addr",
			Value:       "0.0.0.0:8090",
			Usage:       "address to bind the metrics",
			EnvVars:     []string{"GOMEMATIC_API_METRICS_ADDR"},
			Destination: &cfg.Metrics.Addr,
		},
		&cli.StringFlag{
			Name:        "metrics-token",
			Value:       "",
			Usage:       "token to make metrics secure",
			EnvVars:     []string{"GOMEMATIC_API_METRICS_TOKEN"},
			Destination: &cfg.Metrics.Token,
		},
		&cli.StringFlag{
			Name:        "server-addr",
			Value:       "0.0.0.0:8080",
			Usage:       "address to bind the server",
			EnvVars:     []string{"GOMEMATIC_API_SERVER_ADDR"},
			Destination: &cfg.Server.Addr,
		},
		&cli.BoolFlag{
			Name:        "server-pprof",
			Value:       false,
			Usage:       "enable pprof debugging",
			EnvVars:     []string{"GOMEMATIC_API_SERVER_PPROF"},
			Destination: &cfg.Server.Pprof,
		},
		&cli.BoolFlag{
			Name:        "server-docs",
			Value:       true,
			Usage:       "enable swagger documentation",
			EnvVars:     []string{"GOMEMATIC_API_SERVER_DOCS"},
			Destination: &cfg.Server.Docs,
		},
		&cli.StringFlag{
			Name:        "server-host",
			Value:       "http://localhost:8080",
			Usage:       "external access to server",
			EnvVars:     []string{"GOMEMATIC_API_SERVER_HOST"},
			Destination: &cfg.Server.Host,
		},
		&cli.StringFlag{
			Name:        "server-root",
			Value:       "/",
			Usage:       "path to access the server",
			EnvVars:     []string{"GOMEMATIC_API_SERVER_ROOT"},
			Destination: &cfg.Server.Root,
		},
		&cli.StringFlag{
			Name:        "db-dsn",
			Value:       "boltdb://storage/gomematic.db",
			Usage:       "database dsn",
			EnvVars:     []string{"GOMEMATIC_API_DB_DSN"},
			Destination: &cfg.Database.DSN,
		},
		&cli.StringFlag{
			Name:        "upload-dsn",
			Value:       "file://storage/uploads/",
			Usage:       "uploads dsn",
			EnvVars:     []string{"GOMEMATIC_API_UPLOAD_DSN"},
			Destination: &cfg.Upload.DSN,
		},
		&cli.DurationFlag{
			Name:        "session-expire",
			Value:       time.Hour * 24,
			Usage:       "session expire duration",
			EnvVars:     []string{"GOMEMATIC_API_SESSION_EXPIRE"},
			Destination: &cfg.Session.Expire,
		},
		&cli.StringFlag{
			Name:        "session-secret",
			Value:       uniuri.NewLen(32),
			Usage:       "session encription secret",
			EnvVars:     []string{"GOMEMATIC_API_SESSION_SECRET"},
			Destination: &cfg.Session.Secret,
		},
		&cli.BoolFlag{
			Name:        "admin-create",
			Value:       true,
			Usage:       "create an initial admin user",
			EnvVars:     []string{"GOMEMATIC_API_ADMIN_CREATE"},
			Destination: &cfg.Admin.Create,
		},
		&cli.StringFlag{
			Name:        "admin-username",
			Value:       "admin",
			Usage:       "initial admin username",
			EnvVars:     []string{"GOMEMATIC_API_ADMIN_USERNAME"},
			Destination: &cfg.Admin.Username,
		},
		&cli.StringFlag{
			Name:        "admin-password",
			Value:       "admin",
			Usage:       "initial admin password",
			EnvVars:     []string{"GOMEMATIC_API_ADMIN_PASSWORD"},
			Destination: &cfg.Admin.Password,
		},
		&cli.StringFlag{
			Name:        "admin-email",
			Value:       "",
			Usage:       "initial admin email",
			EnvVars:     []string{"GOMEMATIC_API_ADMIN_EMAIL"},
			Destination: &cfg.Admin.Email,
		},
		&cli.BoolFlag{
			Name:        "tracing-enabled",
			Value:       false,
			Usage:       "enable open tracing",
			EnvVars:     []string{"GOMEMATIC_API_TRACING_ENABLED"},
			Destination: &cfg.Tracing.Enabled,
		},
		&cli.StringFlag{
			Name:        "tracing-endpoint",
			Value:       "",
			Usage:       "open tracing endpoint",
			EnvVars:     []string{"GOMEMATIC_API_TRACING_ENDPOINT"},
			Destination: &cfg.Tracing.Endpoint,
		},
	}
}

func serverBefore(cfg *config.Config) cli.BeforeFunc {
	return func(c *cli.Context) error {
		setupLogger(cfg)
		return nil
	}
}

func serverAction(cfg *config.Config) cli.ActionFunc {
	return func(c *cli.Context) error {
		tracing, err := setupTracing(cfg)

		if err != nil {
			log.Fatal().
				Err(err).
				Msg("failed to setup tracing")
		}

		if tracing != nil {
			defer tracing.Close()
		}

		uploads, err := setupUploads(cfg)

		if err != nil {
			log.Fatal().
				Err(err).
				Msg("failed to setup uploads")
		}

		log.Info().
			Fields(uploads.Info()).
			Msg("preparing uploads")

		if uploads != nil {
			defer uploads.Close()
		}

		storage, err := setupStorage(cfg)

		if err != nil {
			log.Fatal().
				Err(err).
				Msg("failed to setup database")
		}

		log.Info().
			Fields(storage.Info()).
			Msg("preparing database")

		if storage != nil {
			defer storage.Close()
		}

		for i := 0; i < 10; i++ {
			err := storage.Ping()

			if err != nil {
				dur := time.Duration(math.Pow(2, float64(i))) * time.Second

				log.Warn().
					Str("retry", dur.String()).
					Msg("database ping failed")

				time.Sleep(dur)
			}
		}

		if err := storage.Migrate(); err != nil {
			log.Fatal().
				Err(err).
				Msg("failed to migrate database")
		}

		if cfg.Admin.Create {
			err := storage.Admin(
				cfg.Admin.Username,
				cfg.Admin.Password,
				cfg.Admin.Email,
			)

			if err != nil {
				log.Warn().
					Err(err).
					Str("username", cfg.Admin.Username).
					Str("password", cfg.Admin.Password).
					Str("email", cfg.Admin.Email).
					Msg("failed to create admin")
			} else {
				log.Info().
					Str("username", cfg.Admin.Username).
					Str("password", cfg.Admin.Password).
					Str("email", cfg.Admin.Email).
					Msg("admin successfully stored")
			}
		}

		registry := service.New()

		registry.Teams = teams.NewService(storage.Teams())
		registry.Teams = teams.NewLoggingService(registry.Teams, requestid.Get)
		registry.Teams = teams.NewTracingService(registry.Teams, requestid.Get)

		registry.Users = users.NewService(storage.Users())
		registry.Users = users.NewLoggingService(registry.Users, requestid.Get)
		registry.Users = users.NewTracingService(registry.Users, requestid.Get)

		var gr group.Group

		{
			server := &http.Server{
				Addr:         cfg.Server.Addr,
				Handler:      router.Server(cfg, uploads, registry),
				ReadTimeout:  5 * time.Second,
				WriteTimeout: 10 * time.Second,
			}

			gr.Add(func() error {
				log.Info().
					Str("addr", cfg.Server.Addr).
					Msg("starting http server")

				return server.ListenAndServe()
			}, func(reason error) {
				ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
				defer cancel()

				if err := server.Shutdown(ctx); err != nil {
					log.Error().
						Err(err).
						Msg("failed to shutdown http gracefully")

					return
				}

				log.Info().
					Err(reason).
					Msg("http shutdown gracefully")
			})
		}

		{
			server := &http.Server{
				Addr:         cfg.Metrics.Addr,
				Handler:      router.Metrics(cfg),
				ReadTimeout:  5 * time.Second,
				WriteTimeout: 10 * time.Second,
			}

			gr.Add(func() error {
				log.Info().
					Str("addr", cfg.Metrics.Addr).
					Msg("starting metrics server")

				return server.ListenAndServe()
			}, func(reason error) {
				ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
				defer cancel()

				if err := server.Shutdown(ctx); err != nil {
					log.Error().
						Err(err).
						Msg("failed to shutdown metrics gracefully")

					return
				}

				log.Info().
					Err(reason).
					Msg("metrics shutdown gracefully")
			})
		}

		{
			stop := make(chan os.Signal, 1)

			gr.Add(func() error {
				signal.Notify(stop, os.Interrupt)

				<-stop

				return nil
			}, func(err error) {
				close(stop)
			})
		}

		return gr.Run()
	}
}
