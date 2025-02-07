# 可执行文件名称
SERVER=userSvr

# 配置文件路径
CONFIG_FILE=./conf/trpc_go.yaml

# 日志文件路径
LOG_FILE=start.log

# 默认目标
all: stop build run clean

# 构建目标
build:
	@echo "Building the binary..."
	go build -o $(SERVER) .

# 运行目标
run:
	@echo "Running the server..."
	@nohup ./$(SERVER) --conf=$(CONFIG_FILE) > $(LOG_FILE) 2>&1 &
	@echo "Checking if the service started successfully..."
	@sleep 4   # 用 nohup 命令在后台运行服务器, 添加短暂的延迟，确保进程有足够的时间启动
	@if pgrep -f "$(SERVER)" > /dev/null; then \
		echo "Service started successfully."; \
	else \
		echo "Failed to start service."; \
		exit 1; \
	fi

# 清理目标
clean:
	@echo "Cleaning up..."
	rm -f $(SERVER)

# 停止服务目标
stop:
	@echo "Stopping the service..."
	pkill -f "$(SERVER)" || true

# 重新启动目标
restart: stop all

.PHONY: all build run clean stop restart