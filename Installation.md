# Installation

## Installing dependencies

```bash
go mod tidy
```

or :

```bash
go mod download
```

or :

```bash
go run main.go
```

### Generate JWT Secret

open a terminal and enter :

```bash
openssl rand -base64 24
```

then copy the generated string and put it in the `.env` file :

```env
JWT_SECRET=dZvG+ecNWSRb6WNAX6l/sC2gh2qzlLl3
```

## Running development server

```bash
go run main.go
```

## Testing

```bash
go test ./tests/...
```



## Building

autobuild :

```bash
go build -o myapp main.go
```

build for linux :

```bash
GOOS=linux GOARCH=amd64 go build -o myapp-linux main.go
```
