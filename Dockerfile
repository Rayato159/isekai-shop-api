FROM golang:1.22-bullseye AS build

WORKDIR /app

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 go build -o /bin/app

FROM debian:bullseye-slim

COPY --from=build /bin/app /bin
COPY ./config.yml /bin/config.yml

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

EXPOSE 8080

CMD ["/bin/app"]