package server

import (
	"PDI-COBRANCA/build/package/configs"
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
)

var (
	E   = echo.New()
	cfg = configs.ConfigsApp{}
)

func init() {
	err := cleanenv.ReadEnv(&cfg)
	fmt.Printf("%+v", cfg)
	if err != nil {
		E.Logger.Fatal("Unable to load configuration")
	}
}
func StartServer() {

	E.Logger.Print(fmt.Sprintf("Listening on port %s....", cfg.Port))
	E.Logger.Fatal(E.Start(fmt.Sprintf("localhost:%s", cfg.Port)))
}
