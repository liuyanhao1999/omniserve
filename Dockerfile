# -- 第一个构建阶段：使用 Go 官方镜像来构建你的 Go 程序 --
# 将基础镜像版本从 1.22 改为 1.23.8
FROM golang:1.23.8 AS builder

# 设置工作目录
WORKDIR /app

# 将本地项目的代码复制到容器中
COPY . .

# 下载 Go 模块依赖，这一步可以缓存
RUN go mod download

# 编译你的 Go 程序，生成一个名为 elegance-gateway 的可执行文件
# -o 参数指定了输出文件名
# CGO_ENABLED=0 是为了创建静态链接的可执行文件，这样就不依赖于系统库
# -a -installsuffix cgo 也是为了减小体积
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o elegance-gateway .

# --- 第二个运行阶段：使用一个非常小的基础镜像来运行程序 ---
# alpine 是一个非常轻量级的 Linux 发行版，非常适合作为基础镜像
FROM alpine:latest

# 设置工作目录
WORKDIR /root/

# 暴露你的应用程序所使用的端口。假设你的 Go 程序监听在 8080 端口
EXPOSE 8080

# 复制第一阶段构建好的可执行文件
COPY --from=builder /app/elegance-gateway .

# 启动你的程序
CMD ["./elegance-gateway"]