
# Rest API Server Skeleton

![](https://i0.wp.com/phocode.com/wp-content/uploads/2016/08/golang.sh-600x600.png?fit=300%2C300&ssl=1)

![](https://img.shields.io/github/stars/huynhphuchuy/server.svg) ![](https://img.shields.io/github/forks/huynhphuchuy/server.svg) ![](https://img.shields.io/github/tag/huynhphuchuy/server.svg) ![](https://img.shields.io/github/release/huynhphuchuy/server.svg) ![](https://img.shields.io/github/issues/huynhphuchuy/server.svg)

## Project Structure

    github.com/huynhphuchuy/server
    ├── cmd/
    │   ├── cli/
    │   │   ├── email-template/
    │   │   ├── stress-test/
    │   │   └── cli.go
    │   └── api/
    │       ├── routes/
    │       │   └── handlers/
    │       ├── tests/
    │       └── server.go
    ├── internal/
    │   ├── config/
    │   │   ├── config.go
    │   │   ├── dev.yaml
    │   │   └── prod.yaml
    │   ├── helpers/
    │   │   ├── generator/
    │   │   └── messages/
    │   ├── registrations/
    │   │   └── user.go
    │   └── platform/
    │       ├── smtp/
    │       ├── mongo/
    │       └── auth/
    └── vendor/
        ├── github.com/
        └── golang.org/

## Features

- Integrated simple registration and bearer authentication
- Request logger to db
- Email verification
- Email template
- CLI tools

## How to run?

1. Copy *.yaml.sample template from `internal/config/samples/` to `internal/config/` and rename to *.yaml
2. Run `dep ensure` to install needed packages to vendor folder
3. Start with `go run cmd/api/server.go -e dev` or `go run cmd/api/server.go -e prod` depends on the environment
4. Run stress test cli tool with following command `echo "GET http://localhost:6969" | ./vegeta attack -duration=0 -connections=1000000`
5. Enjoy the show!

