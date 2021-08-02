go mod init github.com/U24Lab/sample-gql-go

go get github.com/99designs/gqlgen

go run github.com/99designs/gqlgen init

go mod tidy

If required:
go get github.com/vektah/gqlparser/v2@v2.1.0
go get github.com/vektah/gqlparser/v2/ast@v2.1.0
