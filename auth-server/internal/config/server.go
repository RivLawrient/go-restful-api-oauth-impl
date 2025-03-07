package config

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func NewServer(ech *echo.Echo) *http.Server {
	server := http.Server{
		Addr:    os.Getenv("SERVER_ADDR"),
		Handler: ech,
	}

	fmt.Println("server running at ", server.Addr)
	return &server
}
