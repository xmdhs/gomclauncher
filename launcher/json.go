package launcher

type launcherjson struct {
	ID string `json:"id"`
}

type LauncherjsonX115 struct {
	Hidden bool   `json:"hidden" example:"false"`
	ID     string `json:"id" example:"1.15.2"`
	PatchX115
	Root bool `json:"root" example:"true"`
}

type PatchX115 struct {
	MinecraftArguments     string           `json:"minecraftArguments"`
	AssetIndex             AssetIndexX115   `json:"assetIndex"`
	Arguments              ArgumentsX115    `json:"arguments"`
	Assets                 string           `json:"assets" example:"1.15"`
	Downloads              DownloadsX115jar `json:"downloads"`
	ID                     string           `json:"id" example:"game"`
	Libraries              []LibraryX115    `json:"libraries"`
	Logging                LoggingX115      `json:"logging"`
	MainClass              string           `json:"mainClass" example:"net.minecraft.client.main.Main"`
	MinimumLauncherVersion int              `json:"minimumLauncherVersion" example:"21"`
	Priority               int              `json:"priority" example:"0"`
	ReleaseTime            string           `json:"releaseTime" example:"2020-01-17T18:03:52+08:00"`
	Time                   string           `json:"time" example:"2020-01-17T18:03:52+08:00"`
	Type                   string           `json:"type" example:"release"`
	Version                string           `json:"version" example:"1.15.2"`
}

type ArgumentsX115 struct {
	Game []interface{} `json:"game"`
	Jvm  []interface{} `json:"jvm"`
}

type AssetIndexX115 struct {
	ID        string  `json:"id" example:"1.15"`
	Sha1      string  `json:"sha1" example:"7f84a500c46d0815e70b03f444a6171f0ab23f1d"`
	Size      float64 `json:"size" example:"234878"`
	TotalSize float64 `json:"totalSize" example:"2.16522528e+08"`
	URL       string  `json:"url" example:"https://launchermeta.mojang.com/v1/packages/7f84a500c46d0815e70b03f444a6171f0ab23f1d/1.15.json"`
}

type DownloadsX115jar struct {
	Client         ClientX115jar      `json:"client"`
	ClientMappings ClientMappingsX115 `json:"client_mappings"`
	Server         ServerX115         `json:"server"`
	ServerMappings ServerMappingsX115 `json:"server_mappings"`
}

type ClientX115jar struct {
	Sha1 string  `json:"sha1" example:"e3f78cd16f9eb9a52307ed96ebec64241cc5b32d"`
	Size float64 `json:"size" example:"1.5531492e+07"`
	URL  string  `json:"url" example:"https://launcher.mojang.com/v1/objects/e3f78cd16f9eb9a52307ed96ebec64241cc5b32d/client.jar"`
}

type ClientMappingsX115 struct {
	Sha1 string  `json:"sha1" example:"bd9efb5f556f0e44f04adde7aeeba219421585c2"`
	Size float64 `json:"size" example:"4.971989e+06"`
	URL  string  `json:"url" example:"https://launcher.mojang.com/v1/objects/bd9efb5f556f0e44f04adde7aeeba219421585c2/client.txt"`
}

type ServerX115 struct {
	Sha1 string  `json:"sha1" example:"bb2b6b1aefcd70dfd1892149ac3a215f6c636b07"`
	Size float64 `json:"size" example:"3.6175593e+07"`
	URL  string  `json:"url" example:"https://launcher.mojang.com/v1/objects/bb2b6b1aefcd70dfd1892149ac3a215f6c636b07/server.jar"`
}

type ServerMappingsX115 struct {
	Sha1 string  `json:"sha1" example:"59c55ae6c2a7c28c8ec449824d9194ff21dc7ff1"`
	Size float64 `json:"size" example:"3.737122e+06"`
	URL  string  `json:"url" example:"https://launcher.mojang.com/v1/objects/59c55ae6c2a7c28c8ec449824d9194ff21dc7ff1/server.txt"`
}

type LibraryX115 struct {
	Downloads DownloadsX115 `json:"downloads"`
	Extract   ExtractX115   `json:"extract"`
	Name      string        `json:"name" example:"com.mojang:patchy:1.1"`
	Natives   NativesX115   `json:"natives"`
	Rules     []RuleX115    `json:"rules"`
}

type DownloadsX115 struct {
	Artifact    ArtifactX115    `json:"artifact"`
	Classifiers ClassifiersX115 `json:"classifiers"`
}

type ArtifactX115 struct {
	Path string  `json:"path" example:"com/mojang/patchy/1.1/patchy-1.1.jar"`
	Sha1 string  `json:"sha1" example:"aef610b34a1be37fa851825f12372b78424d8903"`
	Size float64 `json:"size"`
	URL  string  `json:"url" example:"https://libraries.minecraft.net/com/mojang/patchy/1.1/patchy-1.1.jar"`
}

type ClassifiersX115 struct {
	Javadoc        JavadocX115        `json:"javadoc"`
	NativesLinux   NativesLinuxX115   `json:"natives-linux"`
	NativesMacos   NativesMacosX115   `json:"natives-macos"`
	NativesOsx     NativesOsxX115     `json:"natives-osx"`
	NativesWindows NativesWindowsX115 `json:"natives-windows"`
	Sources        SourcesX115        `json:"sources"`
}

type JavadocX115 struct {
	Path string  `json:"path" example:"org/lwjgl/lwjgl/3.2.1/lwjgl-3.2.1-javadoc.jar"`
	Sha1 string  `json:"sha1" example:"1f6b7050737559b775d797c0ea56612b8e373fd6"`
	Size float64 `json:"size"`
	URL  string  `json:"url" example:"https://libraries.minecraft.net/org/lwjgl/lwjgl/3.2.1/lwjgl-3.2.1-javadoc.jar"`
}

type NativesLinuxX115 struct {
	Path string  `json:"path" example:"org/lwjgl/lwjgl/3.2.1/lwjgl-3.2.1-natives-linux.jar"`
	Sha1 string  `json:"sha1" example:"9bdd47cd63ce102cec837a396c8ded597cb75a66"`
	Size float64 `json:"size"`
	URL  string  `json:"url" example:"https://libraries.minecraft.net/org/lwjgl/lwjgl/3.2.1/lwjgl-3.2.1-natives-linux.jar"`
}

type NativesMacosX115 struct {
	Path string  `json:"path" example:"org/lwjgl/lwjgl/3.2.1/lwjgl-3.2.1-natives-macos.jar"`
	Sha1 string  `json:"sha1" example:"5a4c271d150906858d475603dcb9479453c60555"`
	Size float64 `json:"size"`
	URL  string  `json:"url" example:"https://libraries.minecraft.net/org/lwjgl/lwjgl/3.2.1/lwjgl-3.2.1-natives-macos.jar"`
}

type NativesOsxX115 struct {
	Path string `json:"path" example:"ca/weblite/java-objc-bridge/1.0.0/java-objc-bridge-1.0.0-natives-osx.jar"`
	Sha1 string `json:"sha1" example:"08befab4894d55875f33c3d300f4f71e6e828f64"`
	Size int    `json:"size" example:"5629"`
	URL  string `json:"url" example:"https://libraries.minecraft.net/ca/weblite/java-objc-bridge/1.0.0/java-objc-bridge-1.0.0-natives-osx.jar"`
}

type NativesWindowsX115 struct {
	Path string  `json:"path" example:"org/lwjgl/lwjgl/3.2.1/lwjgl-3.2.1-natives-windows.jar"`
	Sha1 string  `json:"sha1" example:"e799d06b8969db0610e68776e0eff4b6191098bd"`
	Size float64 `json:"size"`
	URL  string  `json:"url" example:"https://libraries.minecraft.net/org/lwjgl/lwjgl/3.2.1/lwjgl-3.2.1-natives-windows.jar"`
}

type SourcesX115 struct {
	Path string  `json:"path" example:"org/lwjgl/lwjgl/3.2.1/lwjgl-3.2.1-sources.jar"`
	Sha1 string  `json:"sha1" example:"106f90ac41449004a969309488aa6e3a2f7d6731"`
	Size float64 `json:"size"`
	URL  string  `json:"url" example:"https://libraries.minecraft.net/org/lwjgl/lwjgl/3.2.1/lwjgl-3.2.1-sources.jar"`
}

type ExtractX115 struct {
	Exclude []string `json:"exclude" example:"META-INF/"`
}

type NativesX115 struct {
	Linux   string `json:"linux"`
	Osx     string `json:"osx" example:"natives-macos"`
	Windows string `json:"windows"`
}

type RuleX115 struct {
	Action string `json:"action" example:"allow"`
	Os     OsX115 `json:"os"`
}

type OsX115 struct {
	Name string `json:"name" example:"osx"`
}

type LoggingX115 struct {
	Client ClientX115 `json:"client"`
}

type ClientX115 struct {
	Argument string   `json:"argument" example:"-Dlog4j.configurationFile=${path}"`
	File     FileX115 `json:"file"`
	Type     string   `json:"type" example:"log4j2-xml"`
}

type FileX115 struct {
	ID   string `json:"id" example:"client-1.12.xml"`
	Sha1 string `json:"sha1" example:"ef4f57b922df243d0cef096efe808c72db042149"`
	Size int    `json:"size" example:"877"`
	URL  string `json:"url" example:"https://launcher.mojang.com/v1/objects/ef4f57b922df243d0cef096efe808c72db042149/client-1.12.xml"`
}
