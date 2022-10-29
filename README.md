dockerdemo - showing how to link containers in docker compose

## Quick start

  $ git clone git@github.com:gregoryv/dockerdemo
  $ cd dockerdemo
  $ CGO_ENABLED=0 go build -ldflags="-s -w" .
  $ docker build -t x .
  $ docker-compose up
