module github.com/explodes/explodio/tokyo

go 1.15

replace github.com/explodes/explodio/stand => ../stand

require (
	github.com/explodes/explodio/stand v0.0.0-00010101000000-000000000000
	github.com/go-pg/pg/v10 v10.7.3
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.1.2
	golang.org/x/crypto v0.0.0-20201117144127-c1f2f97bffc9
	google.golang.org/grpc v1.27.0
	google.golang.org/protobuf v1.25.0
)
