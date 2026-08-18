FROM --platform=$BUILDPLATFORM golang:1.22-alpine AS build
WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -o /out/rig-integrity ./cmd/rig-integrity
FROM alpine:3.20
RUN adduser -D -u 10001 rig
USER rig
COPY --from=build /out/rig-integrity /usr/local/bin/rig-integrity
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/rig-integrity"]
