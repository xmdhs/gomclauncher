package launcher

type launcherjson struct {
	ID string `json:"id"`
}

type LauncherjsonX115 struct {
	Hidden bool   `json:"hidden" example:"false"`
	ID     string `json:"id" example:"1.15.2"`
	patchX115
	Root bool `json:"root" example:"true"`
}

type patchX115 struct {
	MinecraftArguments     string           `json:"minecraftArguments"`
	AssetIndex             assetIndexX115   `json:"assetIndex"`
	Arguments              argumentsX115    `json:"arguments"`
	Assets                 string           `json:"assets" example:"1.15"`
	Downloads              downloadsX115jar `json:"downloads"`
	ID                     string           `json:"id" example:"game"`
	Libraries              []LibraryX115    `json:"libraries"`
	Logging                loggingX115      `json:"logging"`
	MainClass              string           `json:"mainClass" example:"net.minecraft.client.main.Main"`
	MinimumLauncherVersion int              `json:"minimumLauncherVersion" example:"21"`
	Priority               int              `json:"priority" example:"0"`
	ReleaseTime            string           `json:"releaseTime" example:"2020-01-17T18:03:52+08:00"`
	Time                   string           `json:"time" example:"2020-01-17T18:03:52+08:00"`
	Type                   string           `json:"type" example:"release"`
	Version                string           `json:"version" example:"1.15.2"`
}

type argumentsX115 struct {
	Game []interface{} `json:"game"`
	Jvm  []interface{} `json:"jvm"`
}

type assetIndexX115 struct {
	ID        string  `json:"id" example:"1.15"`
	Sha1      string  `json:"sha1" example:"7f84a500c46d0815e70b03f444a6171f0ab23f1d"`
	Size      float64 `json:"size" example:"234878"`
	TotalSize float64 `json:"totalSize" example:"2.16522528e+08"`
	URL       string  `json:"url" example:"https://launchermeta.mojang.com/v1/packages/7f84a500c46d0815e70b03f444a6171f0ab23f1d/1.15.json"`
}

type downloadsX115jar struct {
	Client         clientX115jar      `json:"client"`
	ClientMappings clientMappingsX115 `json:"client_mappings"`
	Server         serverX115         `json:"server"`
	ServerMappings serverMappingsX115 `json:"server_mappings"`
}

type clientX115jar struct {
	Sha1 string  `json:"sha1" example:"e3f78cd16f9eb9a52307ed96ebec64241cc5b32d"`
	Size float64 `json:"size" example:"1.5531492e+07"`
	URL  string  `json:"url" example:"https://launcher.mojang.com/v1/objects/e3f78cd16f9eb9a52307ed96ebec64241cc5b32d/client.jar"`
}

type clientMappingsX115 struct {
	Sha1 string  `json:"sha1" example:"bd9efb5f556f0e44f04adde7aeeba219421585c2"`
	Size float64 `json:"size" example:"4.971989e+06"`
	URL  string  `json:"url" example:"https://launcher.mojang.com/v1/objects/bd9efb5f556f0e44f04adde7aeeba219421585c2/client.txt"`
}

type serverX115 struct {
	Sha1 string  `json:"sha1" example:"bb2b6b1aefcd70dfd1892149ac3a215f6c636b07"`
	Size float64 `json:"size" example:"3.6175593e+07"`
	URL  string  `json:"url" example:"https://launcher.mojang.com/v1/objects/bb2b6b1aefcd70dfd1892149ac3a215f6c636b07/server.jar"`
}

type serverMappingsX115 struct {
	Sha1 string  `json:"sha1" example:"59c55ae6c2a7c28c8ec449824d9194ff21dc7ff1"`
	Size float64 `json:"size" example:"3.737122e+06"`
	URL  string  `json:"url" example:"https://launcher.mojang.com/v1/objects/59c55ae6c2a7c28c8ec449824d9194ff21dc7ff1/server.txt"`
}

type LibraryX115 struct {
	Downloads downloadsX115 `json:"downloads"`
	Extract   extractX115   `json:"extract"`
	Name      string        `json:"name" example:"com.mojang:patchy:1.1"`
	Natives   nativesX115   `json:"natives"`
	Rules     []ruleX115    `json:"rules"`
}

type downloadsX115 struct {
	Artifact    artifactX115    `json:"artifact"`
	Classifiers classifiersX115 `json:"classifiers"`
}

type artifactX115 struct {
	Path string  `json:"path" example:"com/mojang/patchy/1.1/patchy-1.1.jar"`
	Sha1 string  `json:"sha1" example:"aef610b34a1be37fa851825f12372b78424d8903"`
	Size float64 `json:"size"`
	URL  string  `json:"url" example:"https://libraries.minecraft.net/com/mojang/patchy/1.1/patchy-1.1.jar"`
}

type classifiersX115 struct {
	Javadoc        javadocX115        `json:"javadoc"`
	NativesLinux   nativesLinuxX115   `json:"natives-linux"`
	NativesMacos   nativesMacosX115   `json:"natives-macos"`
	NativesOsx     nativesOsxX115     `json:"natives-osx"`
	NativesWindows nativesWindowsX115 `json:"natives-windows"`
	Sources        sourcesX115        `json:"sources"`
}

type javadocX115 struct {
	Path string  `json:"path" example:"org/lwjgl/lwjgl/3.2.1/lwjgl-3.2.1-javadoc.jar"`
	Sha1 string  `json:"sha1" example:"1f6b7050737559b775d797c0ea56612b8e373fd6"`
	Size float64 `json:"size"`
	URL  string  `json:"url" example:"https://libraries.minecraft.net/org/lwjgl/lwjgl/3.2.1/lwjgl-3.2.1-javadoc.jar"`
}

type nativesLinuxX115 struct {
	Path string  `json:"path" example:"org/lwjgl/lwjgl/3.2.1/lwjgl-3.2.1-natives-linux.jar"`
	Sha1 string  `json:"sha1" example:"9bdd47cd63ce102cec837a396c8ded597cb75a66"`
	Size float64 `json:"size"`
	URL  string  `json:"url" example:"https://libraries.minecraft.net/org/lwjgl/lwjgl/3.2.1/lwjgl-3.2.1-natives-linux.jar"`
}

type nativesMacosX115 struct {
	Path string  `json:"path" example:"org/lwjgl/lwjgl/3.2.1/lwjgl-3.2.1-natives-macos.jar"`
	Sha1 string  `json:"sha1" example:"5a4c271d150906858d475603dcb9479453c60555"`
	Size float64 `json:"size"`
	URL  string  `json:"url" example:"https://libraries.minecraft.net/org/lwjgl/lwjgl/3.2.1/lwjgl-3.2.1-natives-macos.jar"`
}

type nativesOsxX115 struct {
	Path string `json:"path" example:"ca/weblite/java-objc-bridge/1.0.0/java-objc-bridge-1.0.0-natives-osx.jar"`
	Sha1 string `json:"sha1" example:"08befab4894d55875f33c3d300f4f71e6e828f64"`
	Size int    `json:"size" example:"5629"`
	URL  string `json:"url" example:"https://libraries.minecraft.net/ca/weblite/java-objc-bridge/1.0.0/java-objc-bridge-1.0.0-natives-osx.jar"`
}

type nativesWindowsX115 struct {
	Path string  `json:"path" example:"org/lwjgl/lwjgl/3.2.1/lwjgl-3.2.1-natives-windows.jar"`
	Sha1 string  `json:"sha1" example:"e799d06b8969db0610e68776e0eff4b6191098bd"`
	Size float64 `json:"size"`
	URL  string  `json:"url" example:"https://libraries.minecraft.net/org/lwjgl/lwjgl/3.2.1/lwjgl-3.2.1-natives-windows.jar"`
}

type sourcesX115 struct {
	Path string  `json:"path" example:"org/lwjgl/lwjgl/3.2.1/lwjgl-3.2.1-sources.jar"`
	Sha1 string  `json:"sha1" example:"106f90ac41449004a969309488aa6e3a2f7d6731"`
	Size float64 `json:"size"`
	URL  string  `json:"url" example:"https://libraries.minecraft.net/org/lwjgl/lwjgl/3.2.1/lwjgl-3.2.1-sources.jar"`
}

type extractX115 struct {
	Exclude []string `json:"exclude" example:"META-INF/"`
}

type nativesX115 struct {
	Linux   string `json:"linux"`
	Osx     string `json:"osx" example:"natives-macos"`
	Windows string `json:"windows"`
}

type ruleX115 struct {
	Action string `json:"action" example:"allow"`
	Os     osX115 `json:"os"`
}

type osX115 struct {
	Name string `json:"name" example:"osx"`
}

type loggingX115 struct {
	Client clientX115 `json:"client"`
}

type clientX115 struct {
	Argument string   `json:"argument" example:"-Dlog4j.configurationFile=${path}"`
	File     fileX115 `json:"file"`
	Type     string   `json:"type" example:"log4j2-xml"`
}

type fileX115 struct {
	ID   string `json:"id" example:"client-1.12.xml"`
	Sha1 string `json:"sha1" example:"ef4f57b922df243d0cef096efe808c72db042149"`
	Size int    `json:"size" example:"877"`
	URL  string `json:"url" example:"https://launcher.mojang.com/v1/objects/ef4f57b922df243d0cef096efe808c72db042149/client-1.12.xml"`
}
