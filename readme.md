# Go / Vue Project Template

This is a template for bootstrapping a Go project that uses a Vuejs/TailwindCSS frontend.

## Features

- Go
- Vuejs
- TailwindCSS
- JWT auth support

## Build

The full frontend and backend of the application are built in containers using a multistep containers.

```
docker build . -t michaelbunch/go-web:latest
```

This will build the build both the Go code and Vuejs code, then merge them into a single container.
