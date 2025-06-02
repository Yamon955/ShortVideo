
Install trpc-cmdline
``` shell
go install trpc.group/trpc-go/trpc-cmdline/trpc@latest
```
Using trpc-cmdline to generate rpc stub instead of a full project:
```shell
trpc create -p helloworld.proto -o out --rpconly
```

Redis docker启动命令     带布隆过滤器
https://github.com/RedisBloom/RedisBloom
``` shell
docker run -p 6379:6379 -it --rm redis/redis-stack-server:latest
```

MinIO 启动
``` shell
minio server path
```

MySQL 启动
``` shell
mysql -u admin -p
```

