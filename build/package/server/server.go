package server

import (
	"PDI-COBRANCA/build/package/configs"
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
)

var (
	e   = echo.New()
	cfg = configs.ConfigsApp{}
)

func init() {
	err := cleanenv.ReadEnv(&cfg)
	fmt.Printf("%+v", cfg)
	if err != nil {
		e.Logger.Fatal("Unable to load configuration")
	}
}
func StartServer() {

	e.Logger.Print(fmt.Sprintf("Listening on port %s....", cfg.Port))
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", cfg.Port)))
}
