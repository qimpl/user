FROM golang:1.14-alpine AS build
WORKDIR /build
COPY . .
RUN go get .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o APP_NAME .

FROM scratch
WORKDIR /bin
COPY --from=build /build/APP_NAME .
EXPOSE 8000
CMD ["./APP_NAME"]
