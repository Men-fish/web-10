package main

import (
	"flag"
	"log"
	"web-10/internal/count/api"
	"web-10/internal/count/config"
	"web-10/internal/count/provider"
	"web-10/internal/count/usecase"

	_ "github.com/lib/pq"
)

func main() {

	configPath := flag.String("config-path", "../../configs/hello_example.yaml", "путь к файлу конфигурации")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	prv := provider.NewProvider(cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.DBname)
	use := usecase.NewUsecase(cfg.Usecase.DefaultMessageCount, prv)
	srv := api.NewServer(cfg.IP, cfg.Port, cfg.API.MaxMessageSize, use)

	srv.Run()
}

// curl -X GET http://localhost:8082/count
// curl -X POST -H "Content-Type: application/json" -d '{"count": 5}' http://localhost:8082/count
