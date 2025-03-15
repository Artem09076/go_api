FROM golang:1.24.1-alpine3.21 AS builder

ENV GOCACHE=/root/.cache/go-build

WORKDIR /myapp

COPY . .

RUN --mount=type=cache,target="/root/.cache/go-build" \
   go build -o apps ./main.go



FROM alpine:3.21

WORKDIR /mysuperapp

COPY --from=builder /myapp/apps ./

EXPOSE 6080


CMD [ "/mysuperapp/apps" ]

