package launcher

type LauncherjsonX115 struct {
	Hidden bool   `json:"hidden"`
	ID     string `json:"id"`
	patchX115
	Root bool `json:"root"`
}

type patchX115 struct {
	MinecraftArguments     string           `json:"minecraftArguments"`
	AssetIndex             assetIndexX115   `json:"assetIndex"`
	Arguments              argumentsX115    `json:"arguments"`
	Assets                 string           `json:"assets"`
	Downloads              downloadsX115jar `json:"downloads"`
	ID                     string           `json:"id"`
	Libraries              []LibraryX115    `json:"libraries"`
	Logging                loggingX115      `json:"logging"`
	MainClass              string           `json:"mainClass"`
	MinimumLauncherVersion int              `json:"minimumLauncherVersion"`
	Priority               int              `json:"priority"`
	ReleaseTime            string           `json:"releaseTime"`
	Time                   string           `json:"time"`
	Type                   string           `json:"type"`
	Version                string           `json:"version"`
}

type argumentsX115 struct {
	Game []interface{} `json:"game"`
	Jvm  []interface{} `json:"jvm"`
}

type assetIndexX115 struct {
	ID        string  `json:"id"`
	Sha1      string  `json:"sha1"`
	Size      float64 `json:"size"`
	TotalSize float64 `json:"totalSize"`
	URL       string  `json:"url"`
}

type downloadsX115jar struct {
	Client         clientX115jar      `json:"client"`
	ClientMappings clientMappingsX115 `json:"client_mappings"`
	Server         serverX115         `json:"server"`
	ServerMappings serverMappingsX115 `json:"server_mappings"`
}

type clientX115jar struct {
	Sha1 string  `json:"sha1"`
	Size float64 `json:"size"`
	URL  string  `json:"url"`
}

type clientMappingsX115 struct {
	Sha1 string  `json:"sha1"`
	Size float64 `json:"size"`
	URL  string  `json:"url"`
}

type serverX115 struct {
	Sha1 string  `json:"sha1"`
	Size float64 `json:"size"`
	URL  string  `json:"url"`
}

type serverMappingsX115 struct {
	Sha1 string  `json:"sha1"`
	Size float64 `json:"size"`
	URL  string  `json:"url"`
}

type LibraryX115 struct {
	Downloads downloadsX115 `json:"downloads"`
	Extract   extractX115   `json:"extract"`
	Name      string        `json:"name"`
	Natives   nativesX115   `json:"natives"`
	Rules     []ruleX115    `json:"rules"`
}

type downloadsX115 struct {
	Artifact    artifactX115    `json:"artifact"`
	Classifiers classifiersX115 `json:"classifiers"`
}

type artifactX115 struct {
	Path string  `json:"path"`
	Sha1 string  `json:"sha1"`
	Size float64 `json:"size"`
	URL  string  `json:"url"`
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
	Path string  `json:"path"`
	Sha1 string  `json:"sha1"`
	Size float64 `json:"size"`
	URL  string  `json:"url"`
}

type nativesLinuxX115 struct {
	Path string  `json:"path"`
	Sha1 string  `json:"sha1"`
	Size float64 `json:"size"`
	URL  string  `json:"url"`
}

type nativesMacosX115 struct {
	Path string  `json:"path"`
	Sha1 string  `json:"sha1"`
	Size float64 `json:"size"`
	URL  string  `json:"url"`
}

type nativesOsxX115 struct {
	Path string `json:"path"`
	Sha1 string `json:"sha1"`
	Size int    `json:"size"`
	URL  string `json:"url"`
}

type nativesWindowsX115 struct {
	Path string  `json:"path"`
	Sha1 string  `json:"sha1"`
	Size float64 `json:"size"`
	URL  string  `json:"url"`
}

type sourcesX115 struct {
	Path string  `json:"path"`
	Sha1 string  `json:"sha1"`
	Size float64 `json:"size"`
	URL  string  `json:"url"`
}

type extractX115 struct {
	Exclude []string `json:"exclude"`
}

type nativesX115 struct {
	Linux   string `json:"linux"`
	Osx     string `json:"osx"`
	Windows string `json:"windows"`
}

type ruleX115 struct {
	Action string `json:"action"`
	Os     osX115 `json:"os"`
}

type osX115 struct {
	Name string `json:"name"`
}

type loggingX115 struct {
	Client clientX115 `json:"client"`
}

type clientX115 struct {
	Argument string   `json:"argument"`
	File     fileX115 `json:"file"`
	Type     string   `json:"type"`
}

type fileX115 struct {
	ID   string `json:"id"`
	Sha1 string `json:"sha1"`
	Size int    `json:"size"`
	URL  string `json:"url"`
}
