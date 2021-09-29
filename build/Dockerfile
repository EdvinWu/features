FROM golang:1.16.3-alpine as build


WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o features-be cmd/features/main.go

FROM scratch
COPY --from=build /build/features-be /opt/features
COPY --from=build /build/app-docker-config.yaml /opt/app-config.yaml
COPY --from=build /build/migration /opt/migration
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /opt
CMD ["/opt/features"]
