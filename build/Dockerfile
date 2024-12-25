FROM golang:1.23-alpine3.21 AS build
RUN apk add --no-cache gcc build-base
COPY . /build
WORKDIR /build
RUN go build -o ./planning-poker-bot .

FROM alpine:3.21 AS prod
WORKDIR /app
COPY --from=build /build/planning-poker-bot .
RUN chmod +x ./planning-poker-bot

ENTRYPOINT ./planning-poker-bot
