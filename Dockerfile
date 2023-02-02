FROM golang:1.19 as BUILDER

# Active le comportement de module indépendant
ENV GO111MODULE=on

# Build en 2 étapes
ENV CGO_ENABLED=0
ENV GOOS=$GOOS
ENV GOARCH=$GOARCH

WORKDIR /api
COPY ./api ./
RUN go mod download && go mod verify && go build -v -o /build/planifyApi .

FROM alpine:3.17.1 as FINAL

COPY --from=BUILDER /build/planifyApi .

CMD ./planifyApi