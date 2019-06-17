package main


func CutOff(s string, length int) string {
	if len(s) < length {
		return s
	}
	return s[:length] + "..."
}

type FileInfo struct {
	Name string
	Path string
}

