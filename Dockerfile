FROM golang:1.22-alpine as builder

WORKDIR /app
RUN apk add --no-cache make nodejs npm

COPY . ./
RUN make install
RUN make build

FROM scratch
COPY --from=builder /tester /tester

EXPOSE 3001
ENTRYPOINT [ "./tester" ]