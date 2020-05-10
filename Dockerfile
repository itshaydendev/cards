# Build Stage
FROM golang:1.14.2-alpine AS build
RUN apk add make
WORKDIR /build
COPY . .
RUN make

# Run stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=build /build/build/cards .

CMD [ "/app/cards" ]
EXPOSE 3000