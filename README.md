Structure of Setup of Graphql with gqlgen
Steps:

1. go mod init github.com/U24Lab/sample-gql-go

2. go get github.com/99designs/gqlgen

3. go run github.com/99designs/gqlgen init

4. go mod tidy

5. If required:
   go get github.com/vektah/gqlparser/v2@v2.1.0
   go get github.com/vektah/gqlparser/v2/ast@v2.1.0
