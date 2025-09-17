# Stage 1: Build frontend
FROM node:current-alpine AS fe-build
RUN apk update && \
    apk upgrade && \
    apk add git && \
    rm -rf /var/cache/apk/*

ENV NODE_ENV=production

WORKDIR /build/shared/fe/api-client/
COPY shared/fe/api-client/package.json shared/fe/api-client/yarn.lock ./
RUN yarn install --frozen-lockfile

WORKDIR /build/fe/
COPY fe/package.json fe/yarn.lock ./
RUN yarn install --frozen-lockfile

WORKDIR /build/shared/fe/api-client/
COPY shared/fe/api-client/ ./
RUN yarn build

WORKDIR /build/fe/
COPY fe/ ./

COPY .git/ ./
RUN sed -i -e "s/__GIT_TAG__/$(git rev-parse --verify HEAD)/g" .env.production && \
    rm -rf .git/

RUN yarn build

# Stage 2: Prepare backend
FROM golang:alpine AS api-build
RUN apk update && \
    apk upgrade && \
    rm -rf /var/cache/apk/*

WORKDIR /build

COPY api/go.mod api/go.sum ./
RUN go mod download
COPY api/ ./
ENV CGO_ENABLED=0 GOOS=linux GOARCH=arm64 GOFLAGS=-buildvcs=false
RUN go build -trimpath -ldflags="-s -w" -o /build/api

# Stage 3: Final container
FROM alpine:latest
WORKDIR /app
RUN apk update && \
    apk upgrade && \
    apk add dumb-init && \
    rm -rf /var/cache/apk/*

RUN addgroup -S app_perms && adduser -S -G app_perms app_perms 
RUN mkdir -p /app/images && chown -R app_perms:app_perms /app/images
USER app_perms:app_perms

COPY --from=fe-build /build/fe/dist ./public
COPY --from=api-build /build/api ./api

ARG app_port
EXPOSE $app_port

ENTRYPOINT ["/usr/bin/dumb-init", "--"]
CMD ["./api"]