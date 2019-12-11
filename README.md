# example-app-backend

This is simple backend service written in GoLang. It provides an HTTP Rest API. The api specification can be found here: [api/openapi.yaml](api/openapi.yaml)

## Build on Docker

This folder contains a [Dockerfile](./Dockerfile). This can be used to build the application and bake it into an image.

## Run on Docker

Here is an example command to run the backend on Docker:

```sh
docker container run --name backend -p 3000:3000 -d backend:v1.0.0
```

## Verify deployment

```sh
curl localhost:3000/api/version
```

This command should return a response like this:

```json
{ "version": "1.0.0" }
```
