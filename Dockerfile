FROM golang:1.17-alpine

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
	
WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

# Build the Go app
RUN go build -o ./out/service-trap .

# Run the binary program produced by `go install`
CMD ["./out/service-trap"]