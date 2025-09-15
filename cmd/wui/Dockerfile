#syntax=docker/dockerfile:1
FROM oven/bun:1.2.22-slim AS web

WORKDIR /build

COPY cmd/wui/package.json cmd/wui/bun.lock ./
RUN bun install --frozen-lockfile

COPY cmd/wui/ .
RUN bun run build

FROM golang:1.25.1-bookworm AS api

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/wui/main.go ./
COPY pkg/ ./pkg
COPY --from=web /build/dist ./dist

RUN go build -o /usr/bin/wui

EXPOSE 3333

CMD [ "/usr/bin/wui" ]
