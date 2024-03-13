package config

import "os"

type Env struct {
	DriverName        string
	DataSource        string
	HttpServerAddress string
	SecretJWT         string
}

func NewEnv() *Env {
	driverName := os.Getenv("DRIVER_NAME")
	dataSource := os.Getenv("DATA_SOURCE")
	httpServerAddress := os.Getenv("HTTP_SERVER_ADDRESS")
	secretJWT := os.Getenv("SECRET_JWT")

	return &Env{
		DriverName:        driverName,
		DataSource:        dataSource,
		HttpServerAddress: httpServerAddress,
		SecretJWT:         secretJWT,
	}
}
