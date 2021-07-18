package configs

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Env struct {
	AppPort       string
	DbPort        string
	DbHost        string
	DbAutoMigrate bool
	DbName        string
	DbUser        string
	GrpcPort      string
}

var (
	EnvFile *Env
)

func init() {
	_ = godotenv.Load()

	env := &Env{}
	env.AppPort = os.Getenv("APP_PORT")
	env.DbAutoMigrate, _ = strconv.ParseBool(os.Getenv("DB_AUTO_MIGRATE"))
	env.DbHost = os.Getenv("DB_HOST")
	env.DbPort = os.Getenv("DB_PORT")
	env.DbName = os.Getenv("DB_NAME")
	env.DbUser = os.Getenv("DB_USER")
	env.GrpcPort = os.Getenv("GRPC_PORT")

	EnvFile = env

}
