ARG GO_VERSION=1.17
ARG NODE_VERSION=14

FROM --platform=$BUILDPLATFORM golang:$GO_VERSION-alpine as go-builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY main.go .
COPY internal/ internal/
ARG TARGETPLATFORM
# Set Golang build envs based on Docker platform string
RUN set -x \
    && case "$TARGETPLATFORM" in \
        'linux/amd64') export GOARCH=amd64 ;; \
        'linux/arm/v6') export GOARCH=arm GOARM=6 ;; \
        'linux/arm/v7') export GOARCH=arm GOARM=7 ;; \
        'linux/arm64') export GOARCH=arm64 ;; \
        *) echo "Unsupported target: $TARGETPLATFORM" && exit 1 ;; \
    esac \
    && mkdir -p frontend/dist \
    && echo stub >frontend/dist/gitkeep \
    && go build -ldflags="-w -s"


FROM --platform=$BUILDPLATFORM node:$NODE_VERSION-alpine AS node-builder
WORKDIR /app

COPY frontend/package.json frontend/package-lock.json frontend/.npmrc ./
ARG FONTAWESOME_NPM_AUTH_TOKEN
RUN npm ci

COPY frontend/ ./
RUN npm run build

FROM alpine
WORKDIR /app
RUN apk add --no-cache lame
COPY --from=go-builder /app/matrimony ./
COPY --from=node-builder /app/dist dist/

ENV MATRIMONY_ADDRESS ":80"
ENV MATRIMONY_STATIC "dist"
ENV MATRIMONY_DATA "/data"
CMD ["./matrimony"]
