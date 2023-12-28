# syntax=docker/dockerfile:1

FROM golang:1.21 AS build-stage

# Set destination for COPY
WORKDIR /app

# Download Go modules 
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code.
COPY . ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /subscriptions

FROM build-stage AS run-test-stage
RUN go test -v ./...

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /subscriptions /subscriptions

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/subscriptions"]