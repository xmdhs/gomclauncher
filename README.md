# gomclauncher
![Go](https://github.com/xmdhs/gomclauncher/workflows/Go/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/xmdhs/gomclauncher)](https://goreportcard.com/report/github.com/xmdhs/gomclauncher)  
一个简单的命令行下的 minecraft 启动器。支持自动下载补全和验证 minecraft 游戏文件以及正版登录，支持启动通过安装程序安装后的 fabric 和 forge，支持 linux windows 和 mac（mac 暂未测试）。
## 使用方法
使用 `-h` 即可获得相关参数的使用说明。

例子 `./gml-linux -h`

启动游戏 `./gml-linux -run 1.16.1 -username xmdhs`

启动游戏并关闭检测启动器更新检测，游戏文件验证，版本隔离 `./gml-linux -run 1.16.1 -username xmdhs -test=f -independent=f -update=f`

首次正版登录 `./gml-linux -run 1.16.1 -email example@example.com -password example`

第二次 `./gml-linux -run 1.16.1 -email example@example.com` 启动器不会保存你的密码，而是保存 accessToken 用于下次免密登录。

首次外置登录 `./gml-linux -run 1.16.1 -email example@example.com -password example -yggdrasil example.com` 无需完整的 api 地址，启动器会按照协议自动补全。

自定义启动 jvm 参数 `./gml-linux -run 1.16.1 -username xmdhs -flag "-XX:+AggressiveOpts -XX:+UseCompressedOops"`

下载游戏并指定镜像下载源并设置使用的协程数为 32 `./gml-linux -downver 1.16.1 -type=mcbbs -int 32`

下载游戏并混合的使用两个下载源 `./gml-linux -downver 1.16.1 -type "mcbbs|vanilla"`

查看所有可以下载的正式版本 `./gml-linux -verlist release`， `release` 为版本类型，可通过下面的命令获取。

查看其他可选的版本类型 `./gml-linux -verlist ?`
## 截图
![image.png](https://i.loli.net/2020/07/02/E7ZcBCGfo1v46kI.png)

## 使用项目
BMCLAPI https://bmclapidoc.bangbang93.com/  
TSS Mirror https://www.mcbbs.net/thread-932755-1-1.html 
