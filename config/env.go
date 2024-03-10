package config

import "os"

type Env struct {
	DriverName        string
	DataSource        string
	HttpServerAddress string
}

func NewEnv() *Env {
	driverName := os.Getenv("DRIVER_NAME")
	dataSource := os.Getenv("DATA_SOURCE")
	httpServerAddress := os.Getenv("HTTP_SERVER_ADDRESS")

	return &Env{
		DriverName:        driverName,
		DataSource:        dataSource,
		HttpServerAddress: httpServerAddress,
	}
}
