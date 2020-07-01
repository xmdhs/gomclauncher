# gomclauncher
一个简单的命令行下的 minecraft 启动器。支持自动下载补全和验证 minecraft 游戏文件以及正版登录，支持 linux windows 和 mac 。

## 使用方法
使用 -h 即可获得相关参数的使用说明。

例子 `./gml-linux -h`

启动游戏 `./gml-linux -run 1.16.1 -username xmdhs`

启动游戏并关闭检测启动器更新，游戏文件验证，版本隔离 `./gml-linux -run 1.16.1 -username xmdhs -test=f -independent=f -updata=f`

自定义启动 jvm 参数 `./gml-linux -run 1.16.1 -username xmdhs -flag "-XX:+AggressiveOpts -XX:+UseCompressedOops"`

下载游戏并使用镜像下载源 `./gml-linux -downver 1.16.1 -type=mcbbs`

## 截图
![image.png](https://i.loli.net/2020/07/02/E7ZcBCGfo1v46kI.png)

## 使用项目
https://github.com/google/uuid

BMCLAPI https://bmclapidoc.bangbang93.com/