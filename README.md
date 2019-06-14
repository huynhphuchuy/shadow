
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

