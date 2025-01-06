translation

En, [Link Text](/README.md)

# gomclauncher
![Go](https://github.com/xmdhs/gomclauncher/workflows/Go/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/xmdhs/gomclauncher)](https://goreportcard.com/report/github.com/xmdhs/gomclauncher)  
A simple Minecraft launcher for command line. It supports automatic download completion and verification of Minecraft game files, as well as official login. It supports launching through installed Fabric and Forge, and is compatible with Linux, Windows, and macOS (macOS has not been tested).

## Usage
Use `-h` to get relevant parameter usage instructions.

Example: `./gml-linux -h`

To start the game: `./gml-linux -run 1.16.1 -username xmdhs`

To start the game with a specific Java: `./gml-linux -run 1.16.1 -username xmdhs -javapath "./java"`

To start the game and disable launcher update checks, game file verification, and version isolation: `./gml-linux -run 1.16.1 -username xmdhs -test=f -independent=f -update=f`

First-time official login: `./gml-linux -run 1.16.1 -email example@example.com -password example`

Second time: `./gml-linux -run 1.16.1 -email example@example.com` The launcher will not save your password but will save the accessToken for next password-free login.

Login with Microsoft account: `./gml-linux -run 1.16.1 -email example@example.com -ms`

First-time external login: `./gml-linux -run 1.16.1 -email example@example.com -password example -yggdrasil example.com` No need to provide the full API address; the launcher will automatically complete it according to the protocol.

Second time: `./gml-linux -run 1.16.1 -email example@example.com -yggdrasil example.com`

To view all saved official/external login accounts: `./gml-linux -list`

To delete a saved external login account: `./gml-linux -email example@example.com -yggdrasil example.com -remove`

To delete a saved official login account: `./gml-linux -email example@example.com -remove`

To delete a saved official Microsoft login account: `./gml-linux -email example@example.com -ms -remove`

Customize JVM startup parameters: `./gml-linux -run 1.16.1 -username xmdhs -flag "-XX:+AggressiveOpts -XX:+UseCompressedOops"`

Download the game and specify the mirror download source, with 32 concurrent threads: `./gml-linux -downver 1.16.1 -type=bmclapi -int 32`

Download the game and use two download sources in a mixed way: `./gml-linux -downver 1.16.1 -type "bmclapi|vanilla"`

To view all available official versions: `./gml-linux -verlist release` (where `release` is the version type, which can be obtained with the command below).

To view other optional version types: `./gml-linux -verlist ?`

Remove unused files from the assets/objects folder: `./gml-linux -tidy`

To view the launcher version: `./gml-linux -v`

## Screenshots
![image.png](https://i.loli.net/2020/07/02/E7ZcBCGfo1v46kI.png)

## Projects used
BMCLAPI: https://bmclapidoc.bangbang93.com/  
authlib-injector: https://github.com/yushijinhun/authlib-injector
