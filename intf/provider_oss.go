package intf

type OssProvider interface {
	Upload(rootDir, fileName string, data []byte) (string, error)
}
