package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/elyarsadig/studybud-go/configs"
	"github.com/elyarsadig/studybud-go/internal/application"
	"github.com/elyarsadig/studybud-go/migrations"
	confighandler "github.com/elyarsadig/studybud-go/pkg/configHandler"
	"github.com/elyarsadig/studybud-go/pkg/encryption"
	"github.com/elyarsadig/studybud-go/pkg/errorHandler"
	"github.com/elyarsadig/studybud-go/pkg/logger"
	redispkg "github.com/elyarsadig/studybud-go/pkg/redis"
	"github.com/elyarsadig/studybud-go/pkg/unmarshaller"
	"github.com/elyarsadig/studybud-go/transport"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	configFile := flag.String("c", "", "Path to config file")
	migrate := flag.Bool("migrate", false, "Run DB migrations")
	seed := flag.Bool("seed", false, "seed DB")
	flag.Parse()

	if *configFile == "" {
		log.Fatal("config file must be set with '-c'")
	}

	cfg, err := confighandler.New[configs.ExtraData](*configFile)
	if err != nil {
		log.Fatal(err)
	}

	err = loadEnvVar(cfg.FromEnvFile)
	if err != nil {
		log.Fatal(err)
	}

	errHandler, err := errorHandler.NewError()
	if err != nil {
		log.Fatal(err)
	}

	logger, err := initLogging(cfg)
	if err != nil {
		log.Fatal(err)
	}

	serviceInfo, err := loadServiceAccessControl()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	db, err := initDB(ctx, logger, cfg.Database.Name, cfg.Database.Host, cfg.Database.Port)
	if err != nil {
		log.Fatal(err)
	}

	redisClient := initRedis(cfg)

	if *migrate {
		err = migrations.Migrate(db, logger)
		if err != nil {
			log.Fatal(err)
		}
	}

	if *seed {
		err := migrations.Seed(db)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Successfully seeded the database")
	}

	router := transport.NewHTTPServer(cfg.HttpAddress, logger)

	aes, err := encryption.NewAES[string]([]byte(os.Getenv("SESSION_PRIVATE_KEY")))
	if err != nil {
		log.Fatal(err)
	}

	app, err := application.New(
		ctx,
		router,
		errHandler,
		cfg,
		db,
		redisClient,
		logger,
		serviceInfo,
		aes,
		time.Minute*time.Duration(cfg.ExtraData.SessionExpireDuration),
	)

	if err != nil {
		log.Fatal(err)
	}

	err = app.Run(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func initDB(ctx context.Context, logging logger.Logger, dbName, dbHost, dbPort string) (*gorm.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	logging.InfoContext(ctx, "attempting to connect to DB...")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)
	config := postgres.New(postgres.Config{
		DSN: dsn,
	})

	db, err := gorm.Open(config, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	logging.InfoContext(ctx, "successfully connected to DB",
		"db-address", fmt.Sprintf("%s:%s", dbHost, dbPort))

	return db, nil
}

func loadEnvVar(fileMode bool) error {
	if fileMode {
		currentDir, err := os.Getwd()
		if err != nil {
			return err
		}

		envFilePath := filepath.Join(currentDir, ".env")

		err = godotenv.Load(envFilePath)
		if err != nil {
			return fmt.Errorf("failed to load .env file")
		}
		return nil
	}
	return nil
}

func initRedis(cfg *confighandler.Config[configs.ExtraData]) *redispkg.Redis {
	password := os.Getenv("REDIS_PASSWORD")
	client := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", cfg.Redis.Address, cfg.Redis.Port),
		Password:     password,
		DB:           cfg.Redis.DB,
		ReadTimeout:  cfg.Redis.ReadTimeout,
		WriteTimeout: cfg.Redis.WriteTimeout,
		PoolSize:     cfg.Redis.PoolSize,
		MinIdleConns: cfg.Redis.MinIdleConns,
	})
	redistClient := redispkg.NewRedis(client)
	return &redistClient
}

func initLogging(cfg *confighandler.Config[configs.ExtraData]) (logger.Logger, error) {
	logger, err := logger.New(logger.JSON, logger.DebugLevel)
	if err != nil {
		return nil, err
	}
	return logger, nil
}

func loadServiceAccessControl() (*configs.ServiceInfo, error) {
	cfg := new(configs.ServiceInfo)
	unmarshaller := &unmarshaller.YamlUnmarshaller{
		Data: configs.ServiceInfoYAML,
	}

	return cfg, unmarshaller.Unmarshal(cfg)
}
