FROM golang:1.14-alpine AS build
WORKDIR /build
COPY . .
RUN go get .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o user .

FROM scratch
WORKDIR /bin
COPY --from=build /build/user .
EXPOSE 8000
CMD ["./user"]
