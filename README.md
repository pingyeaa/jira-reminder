### 使用说明
在根目录创建config.ini配置文件，根据以下格式自行配置
```ini
[mysql]
host = 192.168.1.1
port = 4455
username = abc
password = 123456
database = jira

[jira]
project_code = YKJYGYDD,YKJYGYFK
status = 处理中,开发中,重新打开,已受理
type = 子任务

[wechat]
url = https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=[你的企业微信机器人key]
```

- project_code为项目编号
- status为在什么状态先才进行企业微信推送提醒
- type什么类型的JIRA才进行推送
- wechat中的url为企业微信机器人的请求地址

### 编译方式
```bash
go build -mod=vendor
```

### 备注
- 默认推送姓名为JIRA的经办人，数据表对应jiraissues中的ASSIGNEE字段
- 定时任务请自行设置