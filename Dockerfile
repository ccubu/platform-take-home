FROM golang:1.23.2-alpine AS builder
RUN apk add --no-cache ca-certificates git alpine-sdk linux-headers

# Create the user and group files that will be used in the running 
# container to run the process as an unprivileged user.
RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group
WORKDIR /src
COPY go.mod go.sum ./
COPY . .
RUN go mod download
RUN CGO_ENABLED=1 go build -race -installsuffix 'static' -o /platform/takehomeserver ./cmd/server/main.go 

FROM alpine AS final
RUN apk --no-cache add ca-certificates
COPY --from=builder /user/group /user/passwd /etc/
COPY --from=builder /platform /platform

USER nobody:nobody
