# HostLoc 每日刷分脚本

## 获取二进制文件

### (1) 从 Github Release 中下载
```shell
# 以 amd64 架构的系统为例
mkdir HostLoc_CheckIn
cd HostLoc_CheckIn
wget https://github.com/LordPenguin666/HostLoc_CheckIn/releases/download/V1.0/hostloc-check-in-linux-amd64.tar.gz
tar xvf hostloc-check-in-linux-amd64.tar.gz
```

### (2) 编译 (需要拥有 go 环境)
```shell
git clone https://github.com/LordPenguin666/HostLoc_CheckIn.git
cd HostLoc_CheckIn
make default
```

## 使用方法

1. 复制配置文件 example.json `cp example.json config.json`；
2. 修改配置文件 `vim config.json`；
3. (可选) 你也可以通过 `./hostloc -c /path/to/your/config` 指定配置文件路径；
4. (可选) 默认在 `/root/HostLoc_CheckIn/` 下运行，你可以修改目录下的 `run.sh` 脚本；
5. `chmod +x run.sh`
6. 添加定时任务。

## 定时任务
- 需要先安装 cron 有关的软件包

### Debian / Ubuntu
```shell
apt install cron
systemctl enable cron
systemctl start cron
```

### Arch Linux
```shell
pacman -S cronie

# 默认 editor 为 vi, 可能需要在 /etc/environment 下配置环境变量
EDITOR=vim

# systemd
systemctl enable cronie
systemctl start cronie

```

* 使用 `crontab -e` 添加定时任务
* (每日两点执行) `0 2 * * * /root/HostLoc_CheckIn/run.sh`

## 基础配置
```json
  "time": 5,   // 访问空间间隔 (单位 s), 最低为 5s
```

## 多帐号配置

```json
"accounts": [
  {"username": "第一个帐号名", "password": "密码"},
  {"username": "第二个帐号名", "password": "密码"},
  {"username": "cuper", "password": "114514"}
]
```

## Telegram 推送

```json
  "telegram": {
    "enable": true,                    // 开启推送
    "api": "这里填写 bot api",
    "chat_id": "这里填写对话 id"
  }
```
