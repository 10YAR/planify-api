FROM golang:1.19 as BUILDER

# Active le comportement de module indépendant
ENV GO111MODULE=on

# Build en 2 étapes
ENV CGO_ENABLED=0
ENV GOOS=$GOOS
ENV GOARCH=$GOARCH

WORKDIR /goserver
COPY /goserver/go.mod ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /build/buildedApp ./goserver/main.go

FROM scratch as FINAL

COPY --from=BUILDER /build/buildedApp .

ENTRYPOINT ["./buildedApp"]