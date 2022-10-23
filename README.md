# HostLoc 每日刷分脚本

## 使用方法
1. 复制配置文件 example.json ``` cp example.json config.json```
2. 修改配置文文件 ```vim config.json ```
3. 你也可以通过 ```./hostloc -c /path/to/your/config 指定配置文件路径```

## 定时任务
- example (每日两点执行)
`0 2 * * * /root/HostLoc_CheckIn/hostloc`

## telegram 推送
```json
  "telegram": {
    "enable": true,
    "url": "https://api.telegram.org",
    "api": "这里填写 bot api",
    "chat_id": "这里填写对话 id"
  }
```