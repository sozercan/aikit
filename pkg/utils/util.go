package utils

import (
	"fmt"
	"net/url"
	"path"

	"github.com/moby/buildkit/client/llb"
)

func FileNameFromURL(urlString string) string {
	parsedURL, err := url.Parse(urlString)
	if err != nil {
		panic(err)
	}
	return path.Base(parsedURL.Path)
}

func Sh(cmd string) llb.RunOption {
	return llb.Args([]string{"/bin/sh", "-c", cmd})
}

func Shf(cmd string, v ...interface{}) llb.RunOption {
	return llb.Args([]string{"/bin/sh", "-c", fmt.Sprintf(cmd, v...)})
}

func Bashf(cmd string, v ...interface{}) llb.RunOption {
	return llb.Args([]string{"/bin/bash", "-c", fmt.Sprintf(cmd, v...)})
}
