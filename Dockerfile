FROM golang:1.21-bookworm

WORKDIR /src
COPY . .

ARG COMMIT=unknown
ARG BUILD_TIME=1970-01-01T00:00:00Z
RUN CGO_ENABLED=0 go build -trimpath -buildvcs=false \
      -ldflags "-X main.BuildTime=${BUILD_TIME} -X main.Commit=${COMMIT}" \
      -o /usr/local/bin/updater .

EXPOSE 8080
CMD ["updater"]