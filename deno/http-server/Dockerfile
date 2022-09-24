FROM hayd/deno:alpine-1.8.2 as builder
WORKDIR /app
COPY . /app
EXPOSE 8080
CMD ["deno", "run", "--allow-net", "main.ts"]
