// Code generated by go-bindata.
// sources:
// config/config.toml
// DO NOT EDIT!

package config

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _configConfigToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x53\x4b\x6f\xdb\x3c\x10\xbc\xf3\x57\x2c\xe4\xcb\xf7\x01\x8d\x2c\x4b\x4e\xe2\x0a\xf0\x21\x08\x72\x48\xd1\xb4\x40\x72\x34\x82\x62\x25\xad\x25\xc2\x7c\x08\x24\xe5\x3c\x7e\x7d\xb1\xb4\xe5\x44\x6d\x2e\x05\x22\x1f\x68\x72\xf6\x31\x33\x5c\xce\xe0\xda\xf6\x2f\x4e\xb6\x5d\x80\xff\xea\xff\x21\xcf\x16\x05\x9c\xf1\xb2\x82\x4a\x61\xbd\x0b\xb6\x87\x6f\xd6\x77\x03\xc2\x1d\x4a\x43\x5f\xe0\x4a\x29\xb8\xe7\x04\x0f\xf7\xe4\xc9\xed\xa9\x49\xc5\x0c\x1e\x88\xe0\xfb\xed\xf5\xcd\x8f\x87\x1b\xd8\x5a\x07\x4a\xd6\x64\x3c\x81\x34\x5b\xeb\x34\x06\x69\x4d\x2a\xc4\xec\x73\x3e\x31\x83\xbb\x2b\xee\x06\xd7\xd6\x6c\x65\x3b\xb8\xd8\x00\xfe\xbd\xce\x27\xf1\x11\x41\x06\x45\xb0\x86\xe4\x0e\x59\x39\xdc\x0f\x26\x48\x4d\x53\x7e\x89\xd8\x93\xf3\x4c\x74\x0d\xc9\x3e\x4b\x8b\x34\xcf\x12\x21\x36\x38\x84\xce\xba\x47\x01\x60\x50\xc7\x2a\xa3\xf7\x89\x00\xb0\xae\x45\x23\x5f\x0f\x0a\x4f\x1d\x6e\x7f\x72\xe6\x13\x55\x9c\x36\x38\xc5\x48\x96\xc6\x5f\xb9\xca\x38\x0f\x1b\x2d\xcd\xaf\x23\xb4\xc8\x2f\x23\xb8\x28\x8b\xa2\x28\x38\x95\x34\x4a\xc5\xc9\x9d\xf5\x81\x43\xbc\x0e\x7d\x4a\xcf\xa8\x7b\x45\x69\x6d\x35\xd7\xe8\xad\x63\x2c\x3f\xe7\x26\x9e\x1c\xc7\xf1\xca\x3c\x23\x8e\xde\xf3\x19\xaf\x4f\xd6\x35\x5c\xb8\xc1\x80\x15\x7a\x7a\xaf\x47\x47\xce\x67\xa4\xd0\x07\x59\x73\xa6\xd4\xd8\xbe\x83\xe6\x47\xc8\x13\xba\xba\x2b\x2f\xd2\x65\xf2\xa6\xab\x0b\xa1\x2f\xe7\x73\x65\x6b\x54\xcc\xb6\xfc\x9a\x67\x51\xe2\xec\x8f\x88\x69\x91\x31\x6a\x24\xcc\x81\x23\x69\x26\x7b\xda\x5b\x17\x58\xc5\x86\x13\x98\x35\xdf\x9c\x1d\xa2\xf0\x4c\x00\x90\xc1\x4a\x11\x87\x07\x37\x90\x10\x9b\x41\x7e\xa0\x6d\x27\x2b\x34\xf8\x91\xb4\x03\x32\x6a\x8a\x4f\x26\x1a\x79\xd2\x33\x21\xb1\x5c\x16\x8f\x1f\x35\x25\xb3\x97\xce\x1a\x4d\x26\x30\xee\x86\x38\x0c\x0d\xed\x49\xd9\x9e\x4f\xa3\xf7\xb6\xde\x51\x9c\x24\x8d\x75\x27\x0d\x9d\x4d\x59\x26\xb1\x72\xd3\x5b\x69\xe2\x9d\x87\x7a\x6a\x6c\x5e\x5c\x5e\x24\x13\x07\x16\xd1\x82\x4a\x9a\xc6\xbf\x95\x29\xe7\x1a\xd5\x13\x3a\x2a\x9d\xe5\x70\x25\xcd\xce\xff\x7d\xcf\xe5\xe4\x3e\x38\xb0\xee\x07\x58\xc3\x79\x76\xfc\x98\x27\x69\xeb\x5e\xf8\x30\x5f\xe6\xab\x15\x1f\x8a\x8d\xb2\x6d\x7b\x90\xb1\x95\x8a\xa6\x12\x52\x65\xdb\x24\x0a\x7c\xf6\xf2\x95\x81\x45\x76\xd8\x1e\x5c\x2f\x8e\xbb\x0a\xeb\xdd\xd0\x33\xab\x4b\x66\xc8\x12\xe3\x8b\x5c\xc3\x16\x95\x67\x47\x7b\x67\x9f\x5f\xde\xbc\x3e\x21\x00\x3c\x4e\xe3\x74\xf0\x7f\x7f\xd8\xfc\x0e\x00\x00\xff\xff\xcb\xa9\xce\xd5\x2e\x05\x00\x00")

func configConfigTomlBytes() ([]byte, error) {
	return bindataRead(
		_configConfigToml,
		"config/config.toml",
	)
}

func configConfigToml() (*asset, error) {
	bytes, err := configConfigTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/config.toml", size: 1326, mode: os.FileMode(420), modTime: time.Unix(1536001378, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"config/config.toml": configConfigToml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"config": &bintree{nil, map[string]*bintree{
		"config.toml": &bintree{configConfigToml, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

