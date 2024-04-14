FROM golang:1.22-alpine as builder

WORKDIR /app
RUN apk add --no-cache make nodejs npm

COPY . ./
RUN make install
RUN make build

FROM scratch
COPY --from=builder /test-app /test-app

EXPOSE 3000
ENTRYPOINT [ "./test-app" ]