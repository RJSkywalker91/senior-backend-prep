# ---------- build stage ----------
FROM golang:1.25 AS build
WORKDIR /src

# copy module files first for better caching
COPY go.mod ./
RUN go mod download

# bring in source
COPY cmd ./cmd
COPY internal ./internal

# build the main from ./cmd (CGO off for static binary)
ARG VERSION=dev
ARG COMMIT=none
ARG BUILT_AT=unknown
ENV CGO_ENABLED=0 GOOS=linux
RUN go build -trimpath -ldflags="-s -w -X main.buildVersion=$VERSION -X main.buildCommit=$COMMIT -X main.buildTime=$BUILT_AT" \
    -o /out/app ./cmd

# ---------- final stage ----------
FROM gcr.io/distroless/static:nonroot

# copy the built binary to the path we'll run
COPY --from=build /out/app /app

USER nonroot:nonroot
EXPOSE 8080
ENTRYPOINT ["/app"]
