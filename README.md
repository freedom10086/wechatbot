# wechatbot
最近chatGPT异常火爆，想到将其接入到个人微信是件比较有趣的事，所以有了这个项目。项目基于[openwechat](https://github.com/eatmoreapple/openwechat)
开发

### 目前实现了以下功能
 + 群聊@回复
 + 私聊回复
 + 自动通过回复
 
# 注册openai
chatGPT注册可以参考[这里](https://juejin.cn/post/7173447848292253704)

注册完成后申请API KEY [这里](https://platform.openai.com/account/api-keys)

# 安装使用
````
# 获取项目
git clone https://github.com/freedom10086/wechatbot.git

# 进入项目目录
cd wechatbot

# 复制配置文件
copy config.dev.json config.json

# 启动项目
go run main.go

启动前需替换config中的api_key
````

### 备注
go 设置代理 参考 https://goproxy.cn/

git设置代理
````
git config --global http.proxy 'http://127.0.0.1:6081'
git config --global https.proxy 'https://127.0.0.1:6081'


git config --global --unset http.proxy
git config --global --unset https.proxy
````