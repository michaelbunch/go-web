# -----------------------------------------------------
# Build container for compiling the Golang application
# -----------------------------------------------------
FROM golang:alpine AS GoBuilder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY . .

RUN go mod download
RUN go build -o web cmd/web.go

# -----------------------------------------------------
# Build container for generating frontend assets
# -----------------------------------------------------
FROM node:current as FrontendBuilder

RUN mkdir -p /app
COPY ./frontend /app

WORKDIR /app

RUN npm install
RUN npm run build

# -----------------------------------------------------
# Final container
# -----------------------------------------------------

FROM scratch

COPY --from=GoBuilder /build/web /
COPY --from=FrontendBuilder /app/dist /public

WORKDIR /

EXPOSE 3000

ENTRYPOINT ["/web"]
