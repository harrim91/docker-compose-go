package main

import (
	"log"

	"github.com/harrim91/docker-compose-go/client"
)

func main() {
	compose := client.New(&client.GlobalOptions{
		Files: []string{
			"/Users/michaelharrison/Projects/poc/kafka-local/cli/kafka/docker-compose.yml",
			"/Users/michaelharrison/Projects/poc/kafka-local/consumer/docker-compose.yml",
		},
	})

	v, err := compose.Version()

	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("docker compose version: %s", v.Version)

	config, err := compose.Config(&client.ConfigOptions{
		Services: true,
	})

	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("docker compose config: %s", string(config))

	// buildCh, err := compose.Build(nil, os.Stdout)

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// err = <-buildCh

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// up, err := compose.Up(&client.UpOptions{
	// 	Detach: true,
	// }, nil)

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// err = <-up

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// down, err := compose.Down(&client.DownOptions{
	// 	RemoveImages: client.RemoveImageFlagLocal,
	// }, os.Stdout)

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// err = <-down

	// if err != nil {
	// 	log.Fatalln(err)
	// }
}
