// Code generated by go-bindata.
// sources:
// test-data/copy-dir-basic/a.txt
// test-data/copy-dir-basic/b.txt
// test-data/copy-dir-basic/dir/c.txt
// DO NOT EDIT!

package bindata

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

var _testDataCopyDirBasicATxt = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\xe4\x02\x04\x00\x00\xff\xff\xa5\x85\x6e\x48\x02\x00\x00\x00"

func testDataCopyDirBasicATxtBytes() ([]byte, error) {
	return bindataRead(
		_testDataCopyDirBasicATxt,
		"test-data/copy-dir-basic/a.txt",
	)
}

func testDataCopyDirBasicATxt() (*asset, error) {
	bytes, err := testDataCopyDirBasicATxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test-data/copy-dir-basic/a.txt", size: 2, mode: os.FileMode(420), modTime: time.Unix(1441732290, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _testDataCopyDirBasicBTxt = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\xe2\x02\x04\x00\x00\xff\xff\x66\xd6\x43\x63\x02\x00\x00\x00"

func testDataCopyDirBasicBTxtBytes() ([]byte, error) {
	return bindataRead(
		_testDataCopyDirBasicBTxt,
		"test-data/copy-dir-basic/b.txt",
	)
}

func testDataCopyDirBasicBTxt() (*asset, error) {
	bytes, err := testDataCopyDirBasicBTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test-data/copy-dir-basic/b.txt", size: 2, mode: os.FileMode(420), modTime: time.Unix(1441732290, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _testDataCopyDirBasicDirCTxt = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\xe6\x02\x04\x00\x00\xff\xff\x27\xe7\x58\x7a\x02\x00\x00\x00"

func testDataCopyDirBasicDirCTxtBytes() ([]byte, error) {
	return bindataRead(
		_testDataCopyDirBasicDirCTxt,
		"test-data/copy-dir-basic/dir/c.txt",
	)
}

func testDataCopyDirBasicDirCTxt() (*asset, error) {
	bytes, err := testDataCopyDirBasicDirCTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test-data/copy-dir-basic/dir/c.txt", size: 2, mode: os.FileMode(420), modTime: time.Unix(1441732290, 0)}
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
	"test-data/copy-dir-basic/a.txt": testDataCopyDirBasicATxt,
	"test-data/copy-dir-basic/b.txt": testDataCopyDirBasicBTxt,
	"test-data/copy-dir-basic/dir/c.txt": testDataCopyDirBasicDirCTxt,
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
	"test-data": &bintree{nil, map[string]*bintree{
		"copy-dir-basic": &bintree{nil, map[string]*bintree{
			"a.txt": &bintree{testDataCopyDirBasicATxt, map[string]*bintree{
			}},
			"b.txt": &bintree{testDataCopyDirBasicBTxt, map[string]*bintree{
			}},
			"dir": &bintree{nil, map[string]*bintree{
				"c.txt": &bintree{testDataCopyDirBasicDirCTxt, map[string]*bintree{
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

