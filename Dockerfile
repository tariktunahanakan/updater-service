FROM golang:latest

ARG CACHEBUST=1

RUN apt-get update && apt-get install -y curl ca-certificates git

WORKDIR /src
COPY . .

RUN go build \
      -ldflags "-X main.BuildTime=$(date -u '+%Y-%m-%dT%H:%M:%S') -X main.Commit=$(git rev-parse HEAD 2>/dev/null || echo unknown)" \
      -o /usr/local/bin/updater .

EXPOSE 8080
CMD ["updater"]
