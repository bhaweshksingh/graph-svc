FROM golang:alpine as builder
WORKDIR /graph-svc
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o graph-svc cmd/*.go

FROM scratch
COPY --from=builder /graph-svc/graph-svc .
ENTRYPOINT ["./graph-svc","http-serve"]
