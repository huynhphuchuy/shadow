go run cmd/api/server.go -e dev
echo "GET http://localhost:6969" | ./vegeta attack -duration=0 -connections=1000000 

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