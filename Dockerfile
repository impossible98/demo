# one
FROM golang:1.18 as one
WORKDIR /src/
COPY . .
RUN apt-get update; \
    apt-get install --yes --no-install-recommends make; \
    make install; \
    make build
# two
FROM debian:9-slim
WORKDIR /app
COPY --from=one /src/dist/dragonfly ./
ENV PORT=8080
EXPOSE 8080
ENTRYPOINT ["/app/dragonfly", "serve"]
