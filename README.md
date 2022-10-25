# HostLoc 每日刷分脚本

## 获取二进制文件

### 编译 (需要拥有 go 环境)
```
cd HostLoc_CheckIn
make default
```

## 使用方法

1. 复制配置文件 example.json `cp example.json config.json`
2. 修改配置文文件 `vim config.json`
3. (可选) 你也可以通过 `./hostloc -c /path/to/your/config` 指定配置文件路径
4. (可选) 默认在 /root/HostLoc_CheckIn/下运行, 你可以修改目录下的 `run.sh` 脚本
5. 添加定时任务

## 定时任务
- 需要先安装 cron 有关的软件包

### Debian / Ubuntu
```shell
apt install cron
systemctl enable cron
systemctl start cron
```

### Archlinux
```shell
pacman -S cronie

# 默认 editor 为 vi, 可能需要在 /etc/environment 下配置环境变量
EDITOR=vim

# systemd
systemctl enable cronie
systemctl start cronie

```

* 使用 `crontab -e` 添加定时任务
* example (每日两点执行) `0 2 * * * /root/HostLoc_CheckIn/run.sh`

## 多帐号配置

```
"accounts": [
  {"username": "第一个帐号名", "password": "密码"},
  {"username": "第二个帐号名", "password": "密码"},
  {"username": "cuper", "password": "114514"}
]
```

## telegram 推送

```
  "telegram": {
    "enable": true,
    "url": "https://api.telegram.org",
    "api": "这里填写 bot api",
    "chat_id": "这里填写对话 id"
  }
```