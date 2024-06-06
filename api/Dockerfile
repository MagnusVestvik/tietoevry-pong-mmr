FROM golang:alpine AS build

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /usr/local/bin/app


FROM scratch

COPY --from=build /usr/local/bin/app /usr/local/bin/app

ENTRYPOINT ["/usr/local/bin/app"]
