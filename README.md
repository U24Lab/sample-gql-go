# Structure of Setup of Graphql with `gqlgen`

# Steps:

### 1. Create Project

go mod init github.com/U24Lab/sample-gql-go

### 2. Get `gqlgen`

go get github.com/99designs/gqlgen

### 3. Make `gqlgen` create the structure

go run github.com/99designs/gqlgen init

### 4. Make the Modules Tidy

go mod tidy

### 5. If error shows to get parser module add it

If required:

go get github.com/vektah/gqlparser/v2@v2.1.0

go get github.com/vektah/gqlparser/v2/ast@v2.1.0

### 6. Create DB

6: MongoDB:

use tododb

#### Create User

db.createUser({user:"dev",pwd:"Welcome$1", roles:[{role:"dbOwner",db:"tododb"}]})

#### Create Collection

db.createCollection("TODO")
