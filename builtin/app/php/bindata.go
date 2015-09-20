// Code generated by go-bindata.
// sources:
// data/common/dev/Vagrantfile.tpl
// DO NOT EDIT!

package phpapp

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path/filepath"
)

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
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
	name string
	size int64
	mode os.FileMode
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

var _dataCommonDevVagrantfileTpl = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x54\xdd\x6e\xdc\x46\x0f\xbd\xd7\x53\x30\xda\xc4\xb0\x01\x4b\xc2\x17\x7c\xc9\x85\x1b\x1b\x4d\x6d\xa7\x36\xd0\x3a\x46\xd6\xed\x4d\x51\x6c\x66\x35\x94\x34\x8d\x34\x1c\xcc\x8c\xd6\x5e\x78\xf7\xdd\x4b\x4a\x5a\xff\x34\x41\x5a\x20\x31\x66\x29\xf2\x90\x3c\x3c\xe4\x0c\x7e\x46\x8b\x5e\x45\xd4\xb0\x5c\xc3\xc7\x18\xe9\x10\x34\x81\xa5\x08\xa8\x4d\x7c\x91\xcc\x92\x19\xdc\x34\x26\x00\xff\x8b\x0d\xc2\xef\xaa\xf6\xca\xc6\xca\xb4\x08\xf5\x3f\x63\xa1\x22\x3f\x78\x69\x5c\x61\x4b\xae\x43\x1b\x81\x2a\x86\x88\x02\xa1\x9c\x6b\x4d\xa9\xa2\x21\x5b\x04\xf4\x2b\x53\x62\x0e\x97\x11\x42\x43\x7d\xab\x87\xa4\x4b\x84\x46\x59\x9d\x49\x72\xd4\x39\xdc\x10\x74\xa4\x4d\xb5\x16\x58\xc6\x79\x92\xfe\x10\xfa\x80\x43\xb6\xf7\xce\x89\x21\x4f\x92\xe9\x73\x5e\x92\xad\x4c\xdd\x7b\xdc\x4f\x5f\xa7\x07\xd2\xd1\x66\x34\x6d\x12\x80\xf1\x95\xaf\xba\x7c\x49\x77\x70\x0c\x69\xa3\x42\x63\x4a\xf2\xae\x70\x1e\x4b\x13\xf0\xed\xff\xd3\x84\x1d\x67\x70\x41\x81\x1b\xb0\xed\x1a\x2c\xc6\x5b\xf2\x5f\x9e\x85\x4f\x36\x48\x9d\x37\x2b\xe6\x61\x31\x19\xd2\x43\x30\xee\x08\xd2\xfb\x7b\x21\x62\x61\xdc\x42\x69\xed\x31\x04\xd8\x6e\x27\xe0\x39\xc6\xde\x81\x82\xb0\xb6\x25\xf3\x57\x51\xab\xd1\x43\xe5\xa9\x03\xea\x3d\x08\x8a\xb1\x35\x68\xc3\x05\x45\xf2\xdc\x3e\x41\xb1\x1a\xbb\x7b\x56\xc3\x08\xb0\x98\x00\x24\xa5\x53\xb1\xc9\x77\x00\x9c\xf0\x10\xd2\x5d\x64\x7a\xc8\xb1\x00\x74\xcb\x73\xe3\xfa\x1e\xac\x50\x7b\xea\xdd\x13\xcb\x58\xe4\xb9\x55\x4b\x1e\xf3\x7c\x7e\x01\xaa\x96\x51\xf2\x78\x6f\x95\xd7\x02\x1c\x88\xc7\x1f\xa3\x3c\xa7\xee\xb9\x57\x87\x56\xa3\x2d\x0d\x86\xa1\x83\xf0\x58\x69\x08\x4d\x3e\x45\x2f\x46\xac\x63\x88\xbe\xc7\x31\xd1\x07\xea\xad\x1e\x74\x01\xbb\xc9\x8d\xbf\xf6\x4d\x05\xca\xae\x0f\xd8\xeb\xfe\xd5\xa0\x2e\x66\x04\x8c\xe5\xe7\x2e\x62\xc1\x96\x90\x33\xcf\xf0\x6a\xcb\x6e\xf2\x9d\x47\x5a\x10\xcb\xb1\x78\xf4\xca\x98\x18\x0e\x6f\x89\x5c\x7e\xca\xd6\xc8\x64\xc9\x30\xbe\x4f\xa5\x80\x0d\x0c\xf2\xe3\x99\xab\xf3\xb4\x32\x41\x2a\x4c\x43\x83\x6d\x2b\x13\xb7\xad\xb1\xc8\x1c\x96\x1a\x66\xf7\x1c\xb0\x85\xbd\x3d\x58\xb2\xb4\xa6\x9f\x45\xa7\x8c\xcd\x43\x93\x8e\xcd\x30\x55\xd2\x0f\x17\x3d\x50\xf0\x0b\x29\x0d\xaa\x6d\x87\xf1\x57\x5e\xd5\xb2\x3b\x01\x1a\xf4\x38\xf4\xcd\x2c\x3c\x23\x38\x7f\xa4\x64\xe7\x2d\xbc\x88\xde\x1e\xa3\x07\x46\xa4\xf3\xc9\xb2\xf1\xc8\x59\xb6\xdb\x6f\x56\x70\x69\x43\x94\x02\x96\xbd\xe1\x65\x44\xbb\x32\x9e\xac\x44\xfd\xd7\xce\x5f\x86\xd2\x1b\x17\x17\xbc\xe6\x09\x63\x27\xc9\x13\x03\xcf\xe4\xdd\xbb\xf9\xe9\xa7\xcb\xeb\x9b\x24\x60\x84\x8c\x27\x3f\x03\x19\x52\x86\x77\x58\x1e\x81\xfc\xed\x59\x44\x25\x75\x1d\x1f\x00\xb8\x35\xb1\x61\x2e\xa2\xeb\x23\xb4\x54\xd7\x72\x64\xf8\x29\x37\x42\x9b\xe0\x5a\xb5\x46\x9d\x10\xee\x1f\xc0\x3d\xbc\xfc\x11\x5e\x9f\xec\xfd\x0f\x36\xa3\xa7\x87\x2c\x0e\xd0\x70\x02\x05\x13\x52\xd8\xbe\x6d\x7f\x80\xed\x43\x46\xf6\x3a\xda\x61\x2b\x96\x2f\x56\xe6\x8e\xf1\x3b\xde\x50\xd6\x66\x42\xed\x80\x8a\x65\x43\x90\xfe\x21\x11\x7f\x72\x8a\x74\x42\xf8\x55\x7d\x41\x30\x51\x16\x20\x36\x2a\xc2\xe7\x69\x67\x80\x25\xfe\x19\x6a\x62\xed\x8f\x5b\xdb\x0e\x4b\x2b\xf7\x89\x4f\x8b\x18\x06\x15\x8d\xa8\xac\x91\x87\x9d\x84\x13\x2e\xb3\xa1\x0e\x77\x96\x22\x17\xd5\xf8\x52\xb2\x9d\x4e\xeb\x20\x7b\x26\x7b\x38\xcc\x5b\x05\x91\x2f\x77\x61\x6c\xc2\x0b\xf2\x82\xd7\x17\x1d\xa4\xbf\x05\x3c\xbb\x9a\x33\x45\x29\x14\x18\xcb\x82\x0b\x92\xff\x7a\x31\x4e\x0f\x4e\x9e\x90\xc1\x65\x59\x9e\xeb\x58\xcd\x93\xc0\x0d\x84\x9e\xaf\x65\x44\x84\x4c\xfd\x1b\x0c\x03\x10\x8e\x01\xd3\x39\x17\x12\x80\x2f\x5d\x54\x3e\x26\x95\x49\x98\x4b\x48\x27\x61\x49\x0b\xd7\x17\xd7\x79\x9e\xa7\x09\xde\x39\xf2\x11\xce\xce\x7f\xba\x7c\x7f\xb5\xf8\xf0\xe9\xe3\xd5\xcd\xf9\xd5\xd9\xb1\x25\x6b\x64\x35\x55\x19\xcd\x8a\x07\x31\x81\x2b\x17\x33\x3e\x36\xd0\x3b\x2d\x77\x26\x5b\x7f\xf5\xc5\x4c\xda\xcd\xd6\xe0\xd6\xb1\xe1\x7d\x0f\x54\x45\xbe\x36\x98\xb1\x60\x1d\xfa\x28\x37\xe9\x1b\xb6\x4c\xf4\x46\xf6\x11\x50\xeb\x4c\x40\x99\x50\x0a\x66\xb8\xbc\x82\xe9\xd4\x11\x59\x3e\xe1\x7f\x15\xae\x71\x6f\xb2\x37\xf9\xdb\xef\x95\x50\xf6\xbe\x05\x71\xfc\xaa\xff\x53\xea\x18\x17\xfd\x40\x82\x88\x20\x76\x4e\x90\x86\x88\x2c\xcc\xa1\x89\xd1\x85\xa3\xa2\x60\xc4\x72\xe7\x4b\xbe\x2e\x26\x74\x1e\xfb\x46\x90\x1f\xb2\x77\x2b\x78\xf0\x73\x8d\xf2\x50\xf4\xc1\x17\x2d\x95\xaa\x2d\x96\xc6\x16\xbb\x8f\xc9\xb4\x7b\x7f\x07\x00\x00\xff\xff\x46\x23\x40\x80\xf1\x07\x00\x00"

func dataCommonDevVagrantfileTplBytes() ([]byte, error) {
	return bindataRead(
		_dataCommonDevVagrantfileTpl,
		"data/common/dev/Vagrantfile.tpl",
	)
}

func dataCommonDevVagrantfileTpl() (*asset, error) {
	bytes, err := dataCommonDevVagrantfileTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/common/dev/Vagrantfile.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"data/common/dev/Vagrantfile.tpl": dataCommonDevVagrantfileTpl,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"data": &bintree{nil, map[string]*bintree{
		"common": &bintree{nil, map[string]*bintree{
			"dev": &bintree{nil, map[string]*bintree{
				"Vagrantfile.tpl": &bintree{dataCommonDevVagrantfileTpl, map[string]*bintree{
				}},
			}},
		}},
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

