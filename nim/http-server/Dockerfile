FROM nimlang/nim:alpine as builder
WORKDIR /app
COPY . /app
RUN nim c --passL:"-static -no-pie" app.nim

FROM scratch
WORKDIR /app
COPY --from=builder /app/app /app/app
EXPOSE 8080
CMD ["./app"]
