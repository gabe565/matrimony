ARG GO_VERSION=1.18
ARG NODE_VERSION=14

FROM golang:$GO_VERSION-alpine as go-builder
WORKDIR /app

RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download

COPY *.go .
COPY internal/ internal/
RUN go build -ldflags="-w -s"


FROM --platform=$BUILDPLATFORM node:$NODE_VERSION-alpine AS node-builder
WORKDIR /app

COPY frontend/package.json frontend/package-lock.json frontend/.npmrc ./
ARG FONTAWESOME_NPM_AUTH_TOKEN
RUN npm ci

COPY frontend/ ./
RUN npm run build

FROM alpine
LABEL org.opencontainers.image.source="https://github.com/gabe565/matrimony"
WORKDIR /app
COPY --from=go-builder /app/matrimony ./
COPY --from=node-builder /app/dist frontend/

ARG USERNAME=matrimony
ARG UID=1000
ARG GID=$UID
RUN addgroup -g "$GID" "$USERNAME" \
    && adduser -S -u "$UID" -G "$USERNAME" "$USERNAME"
USER $UID

VOLUME /data

ENV MATRIMONY_ADDRESS ":80"
ENV MATRIMONY_DATA "/data"
CMD ["./matrimony"]
