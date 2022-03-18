package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Chicocs/challenge-fullStack-digitalRepublic/internal/endpoint"
	"github.com/Chicocs/challenge-fullStack-digitalRepublic/internal/service"
	transport "github.com/Chicocs/challenge-fullStack-digitalRepublic/internal/transport"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigName("../config/database.yml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")
	viper.AddConfigPath(".")
	errViper := viper.ReadInConfig()
	if errViper != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", errViper))
	}

	logrus.Info(viper.GetString("db_type"))

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"ts", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	defer level.Info(logger)

	var db *sqlx.DB
	{
		var err error
		db, err = sqlx.Open(viper.GetString("db_type"), viper.GetString("db_config"))
		if err != nil {
			logrus.Error(err)
			os.Exit(-1)
		}
	}

	var (
		context    context.Context
		services   = service.NewServiceFactory(db, logger)
		endpoint   = endpoint.MakeEndpoints(services, logger)
		serverHTTP = transport.NewService(context, &endpoint, &logger)
		httpAddr   = flag.String("http.addr", ":8080", "HTTP listen address")
		err        = make(chan error)
	)

	go func() {
		server := &http.Server{
			Addr:         *httpAddr,
			Handler:      serverHTTP,
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Second * 60,
		}
		err <- server.ListenAndServe()
	}()

	fatal := level.Error(logger).Log("exit", <-err)
	if fatal != nil {
		logrus.Error(fatal)
	}

}
