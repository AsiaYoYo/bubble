# bubble清单
一个基于gin+gorm开发的清单小项目。

# 使用指南
## 下载
```
git clone https://github.com/AsiaYoYo/bubble.git
```

## 配置MySQL
1. 在你的数据库中执行以下命令，创建本项目所用的数据库：
```
CREATE DATABASE bubble DEFAULT CHARSET=utf8mb4;
```
2. 在bubble/conf/config.ini文件中按如下提示配置数据库连接信息。
```
port = 9000

[mysql]
user = 你的数据库用户名
password = 你的数据库密码
host = 你的数据库host地址
port = 你的数据库端口
dbname = bubble
```

## 编译
```
go build
```

## 执行
Mac/Unix：
```
./bubble conf/config.ini
```
Windows:
```
bubble.exe conf/config.ini
```

# 使用docker部署
## 构建docker image
```
docker build . -t bubble_app
```

## 启动容器
```
docker run -itd -v /YOUR_DIR/config.ini:/conf/config.ini -p 9090:9090 --name bubble bubble_app
```