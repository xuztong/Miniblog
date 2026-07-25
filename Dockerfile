# ------- 阶段一：builder，专门负责编译 -------
FROM golang:1.25 AS builder
WORKDIR /app

ENV GOPROXY=https://goproxy.cn,direct

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp .

# ------- 阶段二：最终运行镜像，只需要跑起来 -------
FROM alpine:3.22
WORKDIR /app

COPY --from=builder /app/myapp .

EXPOSE 8080
CMD ["./myapp", "-c", "etc/application.yaml"]                                           