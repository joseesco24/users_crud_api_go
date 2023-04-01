# Users Crud Api Python

A really simple CRUD GraphQL API based on Docker and Golang.

**Note:** In develop mode and running locally the docs are available at this [**url**](http://localhost:10050/graphql)

<br/>

## Project Commands

**Note:** Before running any of these commands be sure that your **CWD** is **users_crud_api_go/src** directory.

### Run App

```bash
go run main.go
```

### Download Dependencies

```bash
go mod download
```

### Clean Unused Dependencies

```bash
go mod tidy
```

### Update Dependencies

```bash
go get -u
```

<br/>

## Docker Project Commands

**Note:** Before running any of these commands be sure that your **CWD** is **users_crud_api_go** directory.

### Docker Login Into Github Container Registry

```bash
docker login -u joseesco24 -p < authentication token > ghcr.io
```

### Docker Push The Image To Github Container Registry

```bash
docker push ghcr.io/joseesco24/users_crud_api_go:latest
```

### Docker Pull The Image From Github Container Registry

```bash
docker pull ghcr.io/joseesco24/users_crud_api_go:latest
```

<br/>

## Docker Compose Project Commands

**Note:** Before running any of these commands be sure that your **CWD** is **users_crud_api_go** directory.

### Docker Compose Build Image Using Compose File

```bash
docker-compose -f compose.build.yaml build
```

### Docker Compose Start Dbs Services Using Compose File

```bash
docker-compose -f compose.databases.yaml up
```

### Docker Compose Stop Dbs Services Using Compose File

```bash
docker-compose -f compose.databases.yaml down
```

### Docker Compose Start Project Using Compose File

```bash
docker-compose -f compose.project.yaml up
```

### Docker Compose Stop Project Using Compose File

```bash
docker-compose -f compose.project.yaml down
```

<br/>