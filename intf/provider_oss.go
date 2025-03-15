package intf

type OssProvider interface {
	Provider
	Upload(rootDir, fileName string, data []byte) (string, error)
}
