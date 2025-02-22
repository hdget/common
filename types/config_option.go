package types

type SdkConfig struct {
	App            string
	Debug          bool // debug mode
	ConfigFilePath string
	ConfigRootDirs []string
}
