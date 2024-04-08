# 使用官方 Go 镜像作为构建环境
FROM golang:1.20 AS builder

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum 文件到容器中，并下载依赖
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# 复制剩余的源代码
COPY . .

# 构建应用程序。启用 CGO 禁用，进行静态编译。
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp .

# 使用 scratch 作为运行时镜像的基础，这是一个空白镜像，不包含任何文件系统及内容
FROM scratch

# 从构建器中复制二进制文件到根目录
COPY --from=builder /app/myapp ./

# 暴露端口（如果应用程序使用端口的话）
EXPOSE 8080

# 启动应用程序
ENTRYPOINT ["./myapp"]
