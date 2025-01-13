# 协议文件名
PROTOCOL_FILE_NAME=video.proto

# 根据协议生成对应的桩代码
build:
	trpc create \
		-p PROTOCOL_FILE_NAME \
		-o ./pb \
		--protocol trpc \
		--rpconly \
		--nogomod \
		--mock=false