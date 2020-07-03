# gomclauncher
![Go](https://github.com/xmdhs/gomclauncher/workflows/Go/badge.svg)  
一个简单的命令行下的 minecraft 启动器。支持自动下载补全和验证 minecraft 游戏文件以及正版登录，支持通过安装程序安装后的 fabric 和 forge，支持 linux windows 和 mac。

## 使用方法
使用 -h 即可获得相关参数的使用说明。

例子 `./gml-linux -h`

启动游戏 `./gml-linux -run 1.16.1 -username xmdhs`

启动游戏并关闭检测启动器更新检测，游戏文件验证，版本隔离 `./gml-linux -run 1.16.1 -username xmdhs -test=f -independent=f -updata=f`

首次正版登录 `./gml-linux -run 1.16.1 -email example@example.com -passworld example`

第二次 `./gml-linux -run 1.16.1 -email example@example.com` 启动器不会保存你的密码，而是保存 accessToken 用于下次免密登录。

自定义启动 jvm 参数 `./gml-linux -run 1.16.1 -username xmdhs -flag "-XX:+AggressiveOpts -XX:+UseCompressedOops"`

下载游戏并使用镜像下载源并设置使用协程数为 32 `./gml-linux -downver 1.16.1 -type=mcbbs -int 32`

## 截图
![image.png](https://i.loli.net/2020/07/02/E7ZcBCGfo1v46kI.png)

## 使用项目
BMCLAPI https://bmclapidoc.bangbang93.com/
TSS Mirror https://www.mcbbs.net/thread-932755-1-1.html 