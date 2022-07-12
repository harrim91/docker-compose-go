# docker-compose-go

Allows you to programmatically interact with Docker Compose from your Go code.

## Usage

```go
package main

import (
 "log"

 "github.com/harrim91/docker-compose-go/client"
)

func main() {
  // Create a new Docker Compose client
  // All commands issued with this client will include these global config flags
  compose := client.New(&client.GlobalOptions{
    Files: []string{
      "/path/to/docker-compose.yml",
    },
  })

  // Query the version of Docker Compose being used
  v, err := compose.Version()

  if err != nil {
    log.Fatalln(err)
  }

  log.Printf("docker compose version: %s", v.Version)

  // Run `docker compose config`
  // The global client config with the command specific options means
  // this will run `docker compose --file /path/to/docker-compose.yml config`

  config, err := compose.Config(&client.ConfigOptions{})

  if err != nil {
    log.Fatalln(err)
  }

  log.Printf("docker compose config: %s", string(config))

  // Run `docker compose build`
  // this will run `docker compose --file /path/to/docker-compose.yml build`
  buildCh, err := compose.Build(nil, os.Stdout)

  if err != nil {
    log.Fatalln(err)
  }

  err = <-buildCh

  if err != nil {
    log.Fatalln(err)
  }

  // Run `docker compose up`
  // Merging the global client config with the command specific options means
  // this will run `docker compose --file /path/to/docker-compose.yml up --detach`
  upCh, err := compose.Up(&client.UpOptions{
    Detach: true,
  }, nil)

  if err != nil {
    log.Fatalln(err)
  }

  // Returns a channel, so you can do other things while you wait
  // The channel will emit one thing and then close once the command has executed.
  err = <-upCh

  if err != nil {
    log.Fatalln(err)
  }

  // Run `docker compose down`
  // Here's we're passing os.Stdout as an io.Writer to capture the streamed output from the command
  // os.Stdout means you'll see it in your terminal with all the pretty colours
  // We're also passing in an extra config to override the client config just for this command
  // Merging it all together means this will run:
  // `docker compose --file /path/to/docker-compose.yml --verbose down --rmi local`
  verbose := true

  downCh, err := compose.Down(
    &client.DownOptions{
      RemoveImages: client.RemoveImageFlagLocal,
    },
    os.Stdout,
    &client.GlobalOptions{
      Verbose: &verbose
    })

  if err != nil {
    log.Fatalln(err)
  }

  // another channel
  err = <-downCh

  if err != nil {
    log.Fatalln(err)
  }
}
```
