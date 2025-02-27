FROM --platform=$BUILDPLATFORM golang:1.24-alpine AS build
WORKDIR /src
COPY . .
ARG TARGETOS TARGETARCH
RUN GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /out/migration-tool github.com/ChristophBe/migration-tool/cmd/migration-tool

FROM bash:5.2-alpha
COPY --from=build /out/migration-tool /bin
ENTRYPOINT ["/bin/migration-tool"]
CMD ["-help"]