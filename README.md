
# Rest API Server Skeleton

![](https://cdn-images-1.medium.com/max/1600/1*vHUiXvBE0p0fLRwFHZuAYw.gif)

![](https://img.shields.io/github/stars/huynhphuchuy/shadow.svg) ![](https://img.shields.io/github/forks/huynhphuchuy/shadow.svg) ![](https://img.shields.io/github/tag/huynhphuchuy/shadow.svg) ![](https://img.shields.io/github/release/huynhphuchuy/shadow.svg) ![](https://img.shields.io/github/issues/huynhphuchuy/shadow.svg)

## Project Structure

    github.com/huynhphuchuy/shadow
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
        
## Package Oriented Design 

#### vendor/
- For the purpose of this post, all the source code for 3rd party packages need to be vendored (or copied) into the vendor/ folder. This includes packages that will be used from the company Kit project. Consider packages from the Kit project as 3rd party packages.

#### cmd/
- All the programs this project owns belongs inside the cmd/ folder. The folders under cmd/ are always named for each program that will be built. Use the letter d at the end of a program folder to denote it as a daemon. Each folder has a matching source code file that contains the main package.

#### internal/
- Packages that need to be imported by multiple programs within the project belong inside the internal/ folder. One benefit of using the name internal/ is that the project gets an extra level of protection from the compiler. No package outside of this project can import packages from inside of internal/. These packages are therefore internal to this project only.

#### internal/platform/
- Packages that are foundational but specific to the project belong in the internal/platform/ folder. These would be packages that provide support for things like databases, authentication or even marshaling.

## Features

- Integrated simple registration and Bearer authentication with JWT token
- Email verification after registering
- Email template generator and exporter
- Log HTTP request to mongodb
- CLI tools

## API Documentation

- https://documenter.getpostman.com/view/488619/S1ZueBw5?version=latest

## How to run?

### Manually
1. Copy *.yaml.sample template from `internal/config/samples/` to `internal/config/` and rename to *.yaml
2. Run `dep ensure` to install needed packages to vendor folder
3. Start with `go run cmd/api/server.go -e dev` or `go run cmd/api/server.go -e prod` depends on the environment
4. Run stress test cli tool with following command `echo "GET http://localhost:6969" | ./vegeta attack -duration=0 -connections=1000000`
5. Enjoy the show!

### Automatically
1. Run `dep ensure` first
2. Use `sh MENU.sh` to setup or rename the project
3. Use `sh SERVER.sh` to run server with dev environment
4. Use `sh VEGETA.sh` to stress test the server with 1 million requests
5. Drink a cup of tea..

### Docker
- docker-compose up

## Contact me
- https://fb.com/phuchuy1992
- huynhphuchuy@live.com
