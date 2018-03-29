// Code generated by go-bindata.
// sources:
// assets/banner.txt
// assets/defGopkg.lock
// assets/defGopkg.toml
// schema/mashling_schema-0.2.json
// DO NOT EDIT!

package assets

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

var _assetsBannerTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x55\xcf\xee\xa4\x30\x08\x3e\x8f\x89\xef\xc0\x8d\xf9\xa5\x0b\xdc\x79\x96\x49\xb8\xec\xbd\x87\x5e\xfb\xf0\x1b\x68\xeb\x38\x4e\xd5\xdd\x45\x63\x94\xf2\x7d\x50\xfe\x54\xd8\x04\x11\x91\x55\x04\xd7\xe5\xad\x93\x26\xd4\x75\xec\x1f\xaa\x7e\x8b\xb8\x8e\x99\x49\xbb\xb8\x36\x0c\xb5\xa3\x5c\xa3\x4a\x24\xbc\x2e\x80\x44\xd4\x4c\x45\xd8\xb9\x44\xd7\x05\xa8\x19\xfa\x12\xe1\xd0\x36\xca\x6e\x06\x4d\x0b\xb7\x82\xb8\x2e\xa4\x42\x8c\x88\xf0\x4f\x48\x16\x8d\x4d\xb3\x0a\xee\xf5\x42\x7b\x7a\xed\x11\x1e\xbc\x6a\xec\x03\x00\x69\x17\x2f\xe0\x07\xd6\x17\x65\x82\x1e\x58\x77\xad\xf8\x0e\x87\xbe\x7c\xc8\x05\x16\x50\x14\xa3\x7c\x88\x74\xc4\x4e\xe8\xbc\xaa\xbc\x15\x19\x49\xc9\xf3\x4f\x44\xa2\x5f\x21\xa2\x7a\x3a\x99\x79\x64\x95\xb6\x5e\x38\x52\x7e\xe4\xce\x3b\x29\xbd\xa5\x75\x8d\xcc\x90\x90\x4a\xd9\xa0\xac\x29\xe5\x34\x11\x65\x9c\x82\xd1\xd1\x7d\x9b\x39\x89\xcc\xb0\xee\x9f\xe6\x70\xe0\x5c\xb2\x7a\xf8\x39\x27\xe2\xac\x07\x02\x25\x49\xd2\x08\xe6\x78\xe0\x52\x12\x61\xca\xd9\x8d\x89\x8e\x51\x73\x12\x4a\x97\x04\xa0\xa5\xa4\xc0\x4f\xf7\xed\x65\xe9\x29\xe4\x33\x06\x90\x92\xf3\x19\xc3\xbe\x08\xa7\x31\x00\xe6\x72\xcf\x20\x57\x31\x78\x26\xce\x19\x64\x30\xe8\x39\x01\x48\x29\x53\x02\x3f\x22\x06\xcd\x15\x1e\x67\x11\xc8\xbb\x29\xe4\x2a\x03\xbd\x12\x74\xda\x42\x72\xda\x44\x9b\xf7\xa4\x3e\xe9\xff\x81\xf5\x0a\x26\x8e\xb9\xa1\x5d\xc0\x31\x37\x17\x29\xef\xfd\x37\xc6\x16\x91\x49\xa5\xf7\xab\x0f\xec\x25\xd2\x4b\x1e\xfd\xc5\x1c\x03\xce\xa3\xd1\x94\xd9\x8f\xa3\x4b\xb0\x4f\x6d\xa2\xf6\xd7\x08\xb0\xf4\xc4\xfb\x28\xb3\xdc\x80\xb5\x94\xbc\x3f\xd1\xf4\x77\xdf\x2d\x21\xdf\xec\x37\xda\x2c\xed\x8e\x1a\x4c\x2c\xad\x45\x24\xc9\x5d\xd4\x90\x7b\xd8\x3b\xa1\xf1\xb7\x22\x46\x4f\xc6\x0d\x83\x84\x7b\xdf\x76\x3b\x70\x3d\x71\xcc\x8e\x26\x62\x65\xff\x91\x31\xe1\x90\x1b\x32\x14\xf7\x78\x14\x9d\xe8\x9a\xdc\x87\x37\x2a\x12\x14\x34\x63\xdf\xd4\xd4\x4e\x75\x5c\x97\x75\x79\x3c\x1e\x53\x32\xeb\x4f\x3b\xea\x4f\x21\x15\x6a\x3c\x9f\xf6\xf3\xa1\x7f\xda\x4f\xc3\x18\x98\xdf\xce\xe8\x2f\xfe\x5a\xa1\x9a\x39\xb2\x7d\x8f\x95\x70\x6f\xd6\x7c\x55\x40\x03\x04\x83\x17\x08\x18\xfa\xc3\x42\xf7\xaa\xe1\x13\xb7\x95\xf8\x16\xb7\xec\xc0\xf7\xf5\xb4\x0a\x2f\x73\xcb\x7a\xb8\x9e\x11\x44\xed\x81\xd7\x86\x1c\x3a\xab\x2f\xb3\x5f\x56\xcd\x4c\xda\x77\xbb\xc6\x8a\x63\xe2\xcd\xe4\x3c\x95\xdf\xb9\x35\xe9\x8e\xfe\xca\xbc\x36\xfa\x3f\x01\x00\x00\xff\xff\x0c\x3c\x9d\x92\xb1\x09\x00\x00")

func assetsBannerTxtBytes() ([]byte, error) {
	return bindataRead(
		_assetsBannerTxt,
		"assets/banner.txt",
	)
}

func assetsBannerTxt() (*asset, error) {
	bytes, err := assetsBannerTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/banner.txt", size: 2481, mode: os.FileMode(438), modTime: time.Unix(1519905570, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsDefgopkgLock = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x59\x4b\x73\xe4\x48\x6e\xbe\x4f\xc4\xfc\x87\x8a\xf6\x61\x2f\x23\x31\xdf\x8f\x70\xec\xc5\xbe\xd8\x27\x1f\x66\x6f\x13\x1b\x1d\x00\x12\xc9\x62\x8b\x45\xd2\x49\x96\x5a\x9a\x5f\xef\x60\x95\x1e\xac\x2e\xd6\x88\x72\x6f\xe8\xa0\x4a\x00\x49\x7e\x99\x48\x00\x5f\x82\xff\xb6\xfb\xc7\xbe\x19\x77\xb9\x69\x79\xd7\x8c\x3b\x38\x4e\x7d\xcd\x1d\x17\x98\x38\xfd\xb6\x4b\xfd\xae\xeb\xa7\x1d\xa7\x66\xfa\xf7\x1d\xed\xa1\xab\x79\xdc\x1d\xe0\x79\x87\xbc\x3b\x76\xa9\xef\x78\x87\xcf\xbb\x69\xcf\xbb\x8e\x9f\xa6\xdd\xdf\x12\x0f\x3b\xee\xc6\x63\xe1\xbf\xdd\xff\xfa\xcb\xfc\xf7\xc7\x1f\x43\xe9\xbf\x31\x4d\xe3\x3f\xff\xf9\xeb\x2f\xbb\x1d\x16\xe8\x68\xbf\xfb\xfb\xee\xcb\x01\xc6\x89\xcb\x97\x59\xd8\xc1\x81\x67\x51\xdd\x4c\xfb\x23\xde\x53\x7f\xa8\xbe\x71\xce\x85\x9f\xb1\xff\x93\x1f\xaa\x6f\x63\xdf\x0d\x30\xed\x4f\xc6\x03\xd0\x03\xcc\x38\xfe\xbe\xfb\xe3\xcb\xfd\x97\xd3\x53\x0b\x3f\x36\x63\xd3\x77\xf3\x43\xd0\x69\x99\x9d\x54\xca\x42\x70\x49\x58\xed\x72\x56\x2e\x59\x2d\x30\x05\x29\x11\x82\x4e\xd6\xa1\xfb\xb2\x06\x6f\x05\xc9\xef\xfb\x7e\x68\xf2\x73\x35\x42\x81\x03\x6c\x82\x90\x65\x74\x5a\x0b\xa5\x25\x07\xe7\x29\x08\xaf\xbd\x09\x9e\x13\x19\xd2\x56\x2a\x94\x46\x4b\xab\xe2\x66\x08\x4d\x39\x0e\x23\x77\x55\xdb\xd7\xe5\x38\x6e\xc2\x90\x5c\x50\x4a\xea\x60\x02\x27\x17\x48\x80\x72\x82\x40\xfb\xe4\x52\xb2\x40\x1c\x8c\xd2\xd9\xe2\x79\x41\x8f\x5c\x5e\xe7\x3d\xca\x7b\x71\x6f\xb6\x22\xfb\xc7\x7f\xff\xc7\x7f\xfe\xcf\xef\x7d\x9e\xbe\x43\xe1\x2a\xb7\x7d\xdd\xdf\x51\xdb\x5c\x21\x9c\xc7\xbb\xdd\x17\xee\x1e\xbf\xfc\xf6\xf2\xfb\x38\x35\xed\xdb\x80\xfa\x2e\x37\xf5\x69\xda\xd5\x4a\x38\x07\x49\x2c\xb2\xf7\x56\x45\x00\xe9\x93\x0d\xde\xc4\xe8\x23\x46\x36\x29\xcb\x88\x89\x31\xd0\xd5\x4a\xc4\xbd\xbd\x57\x3f\xb5\x92\xbe\x9b\x4a\x83\xb7\x56\x03\x34\x35\x7d\x37\xdb\x7e\x7f\x5b\xc8\x42\x56\x25\xce\x4d\xd7\xcc\xe3\x55\x75\xd3\x8d\x13\x74\xc4\xab\xca\xb6\xe9\x1e\xb8\xac\xaa\x0e\x7d\xe2\xf6\xb6\xa6\x1a\x9b\xc3\xd0\xae\x3f\x75\xe4\xf2\xd8\xdc\x78\xe3\x78\x1c\x86\xbe\x4c\xab\xba\x89\x4f\xb1\xba\xa6\xba\x70\xe3\xac\x78\x6c\xa6\xe7\xf9\xa0\x5e\x0b\x0b\x0f\xed\xf3\x9a\x78\x7c\x7f\xeb\xea\x12\x96\xc2\x0a\x79\x0f\x8f\x4d\x5f\xc6\x37\xf5\x54\x9a\xba\xe6\x72\xf9\xa0\xa5\xb0\xa2\xd9\x7c\xf5\x70\x61\x04\xa9\xbc\x4d\xd9\x84\x08\xd6\x70\x0e\x36\x4b\xa7\x42\x34\x91\x98\x13\x67\xce\xc9\x59\x92\xff\xfa\xc3\xd5\xfe\xc5\xc1\x1a\x86\xf7\x6d\x1a\x86\x79\x0d\xfd\xb1\x2c\xfc\xf6\x12\x2e\xef\xc3\xc2\xd5\xd9\x29\xd7\xb2\x79\x8f\x2f\xa5\x09\x26\xb8\x94\x1c\x60\x18\x16\xfe\x5d\xc8\x2a\x7e\x1a\xca\x26\xf5\xe9\x27\x8f\xe3\x15\x88\xbf\xb2\xac\x52\x53\x78\x05\xf8\x5f\xce\x99\x7f\x7e\xc2\x3c\x1f\xbb\xcf\xbe\xa1\xee\x89\x2a\x2e\x65\x79\xca\x36\xce\x6a\xf9\xe9\x33\x5b\x75\x9e\x34\x40\x19\x3f\x3f\x6b\xea\x1f\xf8\xc3\x55\xcd\xab\x9f\x87\xd3\xf3\xc0\x1f\xd9\xce\x75\x76\x8b\x4d\x95\x1b\x6e\xd3\x47\x96\x85\xf3\xa5\xc9\x4b\x48\xbe\x09\xb9\xab\x9b\x8e\x7f\x18\x56\xe5\xd8\x75\x0b\xa3\xb6\xbf\x98\x73\x91\x6e\xe6\x41\x75\x80\x0e\x6a\x4e\xeb\xe1\x2d\x0d\x13\x29\x52\x3a\x88\x98\x7c\x90\x4c\x22\x58\x9b\xc1\x2b\xaf\x6c\x66\x63\xa5\xcf\xc1\xfc\x6b\xc3\xfb\x00\xe3\xbe\x6d\xba\xfa\x66\x11\x7c\x9a\xce\x29\xe0\x2d\x3c\xab\x61\x36\xff\xed\x4a\xff\x9a\xc3\xea\xbe\x34\x6d\x0b\x87\xe3\xd3\x8f\x5b\xd8\x36\x58\x51\xdf\xa5\x53\x95\x19\x2f\xc4\xa7\x9d\x5a\xdd\x14\x17\x5c\xf4\xce\x59\xd4\x46\x63\xd6\xd9\x5b\x8d\xc1\x85\x24\x9d\x10\xe8\x95\xd5\xce\x8a\xac\x5d\xde\xba\x7c\x18\x80\xf6\x5c\x4d\xfb\xd2\xe4\xe9\x9a\x9c\xcc\x60\xe6\xc5\x9c\xd5\x57\x68\x04\x04\xc3\xc0\x32\xa1\x0a\x80\x68\x59\x33\x7b\x2d\x38\x64\x30\x02\xa4\x35\xe4\x58\x8a\xe8\xb7\xa2\x49\xf0\xc8\x54\xef\xab\xba\xbf\x1b\x07\xfe\x7e\x8d\xe7\x24\xbd\x66\x19\x94\x18\x90\x9c\x35\xd1\xa6\xac\x12\x53\xb4\xc9\x93\x01\x43\x9a\x85\x92\x51\x68\xa1\x2d\xdb\x55\x18\x5b\x59\x2d\xbf\x6c\x55\xdd\xdf\x15\x1e\x9b\xb6\xe1\x8e\x9e\xaf\x11\x62\x61\x98\xab\xff\x0a\x15\x8a\x00\xe0\x19\x02\x26\x65\x4c\x30\x2a\x42\x96\xe0\x3d\x65\x23\x51\xa1\x46\x6d\x30\xa5\x75\x6e\xfb\xff\x00\xf9\xc4\xa5\x81\xf6\x6e\xec\x60\x18\x56\x70\xae\xb1\x6f\x8c\xd6\xb2\x90\x18\xb5\x71\x40\x32\x26\x52\x91\x51\x3a\x1b\x1c\x45\x91\x38\xc5\x08\xf1\x44\xd6\x36\x79\xf3\x15\xcc\xff\x1e\xf9\xc8\x9b\x10\x18\x43\x14\x84\xa5\x2c\xb5\x12\x16\xad\xcd\x2e\xb2\x34\x48\xe8\x62\x70\x3e\x49\xe0\xa8\x72\x0c\x2b\xc4\x57\xde\x8b\xad\xb8\xe6\x02\xde\xd7\xf9\x30\x55\xe7\x7f\x9b\xa0\xe9\x28\x00\x7d\xd4\x96\x59\x05\x26\x87\x2a\x38\xed\x0c\x22\x44\x34\x29\x39\x23\x05\xa1\x4e\x76\x25\x1b\xe9\xcf\x40\xab\xfb\x6a\x28\xfd\xd4\xe3\x31\xdf\x4a\x42\x4d\xff\x96\x2c\x4e\xa6\xeb\x99\x22\x99\x14\xd0\xc6\x0c\x81\x83\x54\x2a\xa5\x90\x5c\xd4\x42\x11\x71\xcc\xa8\x7c\x44\x72\x42\xbb\xb0\x1d\x5a\x0b\x5d\xfd\x21\xb8\x33\xa2\x0b\x7c\x74\x57\x73\x77\x57\xf7\x55\xe2\x91\x4a\x33\x4c\xfd\x7b\x12\x1c\xe6\xda\x36\xfe\x30\xac\xa0\x7b\xfe\x51\x94\x8e\x05\x2e\x98\xc0\x8b\x7c\x6a\x0e\x3c\x4e\x70\x18\x6e\x14\x11\xb6\x11\xbd\x47\xab\x30\x07\x36\x68\x4c\x04\xeb\xd9\x65\x1f\xb3\x92\x6a\xbe\x5c\x7a\x19\xcc\xcf\xa5\x86\x97\xad\xf9\x44\xa8\x59\xab\xc1\x19\x69\xbc\x30\xd1\xa1\xd2\xca\x03\x12\x90\x14\xa8\x9d\x8e\x0e\x53\x0c\x6c\x2c\x6d\xbe\x65\xbe\x54\x99\xb9\x9e\x4c\xfc\xb4\xed\x44\x4b\x06\x65\x75\xf0\x39\xbb\xec\x82\x09\x3a\xa6\xa0\xbc\xf3\x24\xbd\xd6\x39\x9b\x64\x92\xb4\x49\xac\xdd\x32\xe5\x67\x71\x1d\x8e\x4f\xdb\xf6\x05\xd0\x2a\x9b\x4d\x46\xe9\x7c\x60\x19\x3d\xb0\x8d\x46\x48\x8e\xc2\x8a\x0c\x02\x1d\xa1\xcd\xe9\xa7\xdc\xf5\x0d\x86\x19\xd9\xef\xa7\xb3\xb8\x09\x16\x01\x60\x14\xd2\x58\x14\x56\x7b\x67\xb5\x43\x04\xd4\x5a\x69\xa9\xb2\x09\x52\x49\xa5\x67\xb0\x5b\xb7\xe5\xdb\x71\x2e\x1f\x23\xed\x0f\x4d\x9a\xaa\xfd\x34\x0d\xa5\x3f\xbe\x22\xfe\x08\x0c\x4b\x8c\x41\x05\xa4\xc8\x36\x0a\x83\xc0\x24\xac\x3f\x55\x5b\x11\x09\x8c\xc8\xec\x33\xf0\x7a\xce\xd9\xba\x47\x0f\xe5\x33\xc9\x11\x83\x61\x2d\x80\x92\x95\x16\x20\x91\x41\x1f\x34\x24\x93\x73\xd0\x90\xb3\x56\x31\x62\xca\x37\x20\xad\xbc\xbd\x6d\xea\xfd\x34\x4e\x3c\xbc\xff\xba\x9b\x0a\x10\x97\xbb\xba\xbf\x95\x79\xee\x17\x8c\xb5\x6d\x99\xa6\xbe\x0c\xb8\xa0\x55\x2f\x0f\xfa\xfa\x42\x66\xae\x14\x0b\xe3\xb3\xc9\x57\xf1\x35\x7e\x55\xd5\x25\x07\x5a\xbf\x89\x82\x0e\x08\x2c\xb3\x60\x92\x4c\x1c\x45\x54\x49\x5b\x44\x26\x44\x13\x3d\xb3\x36\x86\x70\xad\xcd\x21\xed\x8d\xea\xb0\xd5\x53\xfd\xc0\xdd\xbc\x37\x4d\x57\xbf\x36\x3d\xe6\x0a\xd7\xe3\xc8\xe5\x71\xe3\x91\x02\xab\xb2\xd2\x46\x19\x13\x95\x71\xc9\x22\x91\xf2\x9a\x9d\x25\x4c\x94\xc1\x66\x9f\x1c\x39\xbd\xb9\xf2\x2f\x30\x55\x08\x63\x43\x9f\x70\xde\xf7\xa6\xf0\xfa\x2e\x7b\x64\x83\x48\x2c\x83\x4a\x5a\xb3\x15\x6c\x2c\x86\x94\x62\x94\x42\x1a\x8a\xe0\x21\x03\xa4\xcd\x0d\xb0\x25\xca\xe5\x2e\x6e\x41\x39\x27\xd9\xc5\x35\xe7\x46\xf1\xd1\x4e\x62\xa4\xe4\x04\xb2\x8f\x64\x48\x83\xcf\x10\x83\x91\xa8\x49\x6a\xc5\x46\x08\xe7\x60\x73\xda\x98\x41\xfe\xd9\x0c\x0f\x4d\x57\x9d\xff\xdd\xcd\x8e\x7e\x47\xfe\x31\xec\xdc\x42\xfd\xc3\x29\xaf\x5e\x8a\xf3\x5c\x9a\x91\x6f\x28\xcf\xaf\x9b\x6f\x82\xef\x06\x17\x85\xfb\xb6\xd7\xb4\x62\x6d\x6c\x42\x8d\x02\x41\x41\x12\x39\xa3\xcd\x42\x13\x1a\x13\xd8\x82\x8e\x09\x84\x88\x9b\x7b\xba\x43\xc3\xa5\x30\x55\xed\x9f\x66\xd3\xd9\x56\x99\x12\x18\x42\x2f\x64\x20\x16\x16\x94\x8d\x36\x26\x95\x59\x04\x0a\x9a\xb5\x56\x31\x4b\x17\x37\x97\xb9\xad\x81\xf9\x0a\xf4\xe9\xe9\xbf\x60\x5c\xe9\x7f\x9f\xe5\x5a\xad\x84\xa3\x10\xc2\xa1\xd4\xe4\x95\xca\x3e\x4b\xa5\x5d\x20\x21\x40\x27\x4d\x0a\x38\x40\x8c\x11\x04\x6d\xdf\xb2\x87\xfa\xb5\x23\xb2\xa9\xc0\x84\x20\x73\xb2\x21\xf9\xc0\xc2\x50\x76\x49\x24\x96\x6a\xbe\x6e\x7b\xe1\x29\x10\x29\x65\x22\x6e\xce\x06\x85\x4a\xff\xbd\xe5\xe7\x39\x2d\x1d\x78\x2a\x0d\x6d\xc4\x21\x83\x64\x11\x2d\x02\x47\x63\x83\xd2\x4e\x1b\x6d\xa4\x31\xe4\x24\x44\x67\xe7\xd2\xe2\xd8\x7e\x50\x55\x4e\xdc\xec\xbe\x2f\x75\xf5\x54\x51\x79\x1e\xa6\xeb\xe8\xfe\x32\x8e\xfb\x6a\xe2\x72\x68\x3a\x68\x57\x6e\xb9\x99\x08\xe6\x0b\x1b\x84\xe4\x0d\xe6\x94\x48\x51\x74\x20\xbc\xc6\xa4\xc0\x24\x05\x1e\x14\x7f\xc0\xa4\x97\x38\x3a\xbe\xae\xab\x6f\xb5\xeb\xcc\xe0\x5e\x03\x6b\x26\x07\xea\x72\x54\xed\xe7\x79\x6f\xb2\x26\x75\xef\x2d\xc1\xa6\x9b\xb8\x74\xd0\x9e\x79\x31\x97\x66\x11\xa3\x2d\x3f\x9d\xb8\x46\xcb\x4f\x8b\x56\x2b\xd0\x8d\xc0\x4d\xc1\x39\xca\xa4\x43\x24\xa6\x18\x6c\x72\x39\x81\x0a\x36\x46\xed\xc0\x7a\x0b\xd6\x82\x86\x8f\x3e\xc6\x2c\x97\x3d\x3e\x5f\x3b\xfe\x05\xc7\xb1\x6b\x9e\x16\xc9\xa4\x4b\xfd\xf7\x1b\x5d\xdf\xa0\x83\x90\x46\x06\x96\xf3\xed\x06\x65\x08\x82\xb5\x9b\x29\x86\x0d\x92\xd9\x1a\x0d\x59\x2b\x82\xed\xb0\x56\x19\xf3\x82\x4a\xc0\xc4\x17\xd4\x02\x26\xae\xf0\xd8\x2c\xba\x68\x6f\xbb\x3e\xab\x27\xc0\x6b\x45\xbd\xe8\xf6\xbd\xfb\x68\x91\x8c\xdf\x85\xa5\xe1\x55\xeb\x23\xbd\xbf\x70\x46\x7f\x84\xfa\x1d\xd7\xc8\x74\x2c\x5c\x61\x93\x9a\x72\x5c\xb4\xdf\xa7\x02\xdd\x98\xfb\x72\x78\x6f\xbc\x75\x0d\xf5\xe9\x6c\x7a\x25\xa4\x36\x95\x2b\x61\xb7\x36\xbd\x40\x57\xf3\x04\xd8\xde\x38\x3d\x2c\x23\xb0\x34\xd1\xc5\x60\x50\x92\xb3\x16\x83\x30\x06\x9c\x25\xa1\xe7\x5c\x46\xc1\x87\x94\xf4\x07\x6e\xea\xeb\x96\xef\x17\xde\xaa\xb9\x7b\xbf\xf3\xae\xf8\xeb\x3c\x03\x86\x66\xac\x60\x68\x2a\xe8\xba\x7e\x82\xcb\x06\xdb\xc2\xa4\x0c\x54\x8d\x13\x4c\xc7\x1b\x67\x0d\x82\x14\x32\x2b\x49\x39\x06\xcd\xde\xeb\x24\x48\x4a\xad\x19\xc9\x1a\x65\x7c\x54\x42\x68\x25\x3f\x60\x1c\xd7\x8b\x28\x03\x7d\x5c\xab\x11\x5a\xe8\x68\xd1\x30\x7c\x15\xcc\x94\x8a\xaf\xa5\xa5\x3f\x76\xa9\xf4\xd8\x2c\xfb\xc0\x89\x97\x5d\xf0\xae\xe3\xab\xaf\x0a\x85\x13\x77\x53\x03\xed\xb8\xe8\xe6\x52\x9f\x96\x6d\xcd\x19\x70\x8b\x27\xdc\x5f\x5b\xfc\xfa\x28\xab\x03\x8f\xe3\x8c\xfc\xd2\xa4\xbf\x3e\xcf\x6f\x82\x07\xe6\x01\xda\xe6\x71\xf1\x69\x88\x27\xb8\xf8\x9e\xd1\xc1\x61\xf9\xd6\x81\x17\x8b\x2f\x3c\xf6\xed\xe3\x8a\xa0\x4a\x0b\xdf\xbe\x09\x07\x18\xc7\x69\x5f\xfa\x63\xbd\x7f\x8f\x91\x09\xa6\xf1\x62\x74\x5c\x7c\x88\x82\xe1\x32\x68\x4e\x1f\xd4\xd6\x33\x90\x32\x4e\x49\xa1\x93\x14\xd9\x90\x71\xca\xb0\x45\x2b\x40\x66\x98\x59\x34\x6a\xa3\xd0\x39\xf5\xd7\x47\xfb\xfc\x55\xa8\x2e\x30\xec\x4f\x45\x72\x31\xae\x60\x18\xd2\x1a\x65\xb8\x3a\x22\x6f\xb9\xe1\x44\xbf\x5e\xa5\x3f\x32\xc2\xeb\x6f\xdc\x11\x8c\xd6\x8e\x84\xb2\x46\x31\x62\x06\x21\x49\xfb\x68\x22\x26\x2f\xad\xca\x28\xb2\xb3\xf1\x83\x02\xdf\x0f\x0f\xf5\x7d\xd3\x55\x30\x76\xf2\x0e\xb9\xdc\x3f\xca\x6d\x9d\x34\x1f\xa5\x09\x04\x42\x29\x9b\xb2\x07\xa3\x95\x90\x0a\x43\xca\x42\x5b\x4b\x0a\x94\x70\x1a\x48\xac\x10\xb2\x0f\xba\xfa\xaf\x88\xda\x04\xc3\xfd\xa3\xda\xd8\xf4\xf4\x10\x09\x1c\x9b\x8c\xa4\xa4\x8a\xac\x13\xda\x10\x40\x1b\xa4\xe8\x45\xce\x1c\xa4\x84\x6b\x7a\xa8\xee\xed\x47\x7d\x90\x57\x3c\xcf\x70\x68\xb7\xe2\xf1\x39\xfa\xe0\x02\x33\x79\x83\x5a\xa1\x88\x41\xa5\x24\x6d\x00\x2b\xc1\x18\x97\x64\x02\xcf\x78\xdd\x67\x54\xf7\x3f\x4b\x58\xbf\x4d\xd8\x36\xdd\xa9\x3d\x9a\x60\xb8\xa3\xb6\xe1\x6e\xe3\xfd\xdf\xeb\xec\x9c\x53\x0e\xb5\x26\x30\x40\x26\x27\x67\x6d\xe0\x8c\x28\x9c\x22\x93\xd9\xb0\x55\xe7\xa4\xf7\xeb\x2f\x7f\x9c\x82\x74\xe6\x82\x70\x7a\x14\x74\xd0\x3e\xff\xc9\xe5\xee\x15\x54\xe2\x73\x3b\xef\x4d\xf1\xbe\x50\x39\xcb\x9b\x6e\x38\x4e\xe3\x5d\x6a\x6a\x1e\xa7\x79\x42\xf4\x56\x38\xa9\x52\x40\x08\xa4\x50\xa5\x99\xd1\x7b\x0c\xde\x23\x90\x95\x48\x9e\xad\x97\x14\xc0\x65\x16\x19\x29\xca\xcc\x29\x05\x29\x53\x54\xda\xd0\xf9\x0b\xd4\x39\x75\xbc\x81\xa8\x87\xf1\x8e\x12\xb5\x4b\xdd\x05\x8e\xff\x0b\x00\x00\xff\xff\x14\x5f\xb4\x38\x80\x23\x00\x00")

func assetsDefgopkgLockBytes() ([]byte, error) {
	return bindataRead(
		_assetsDefgopkgLock,
		"assets/defGopkg.lock",
	)
}

func assetsDefgopkgLock() (*asset, error) {
	bytes, err := assetsDefgopkgLockBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/defGopkg.lock", size: 9088, mode: os.FileMode(438), modTime: time.Unix(1522306830, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsDefgopkgToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x8f\x31\xcf\xd3\x30\x10\x86\xf7\x4a\xfd\x0f\xa7\x74\xa5\xf6\xc7\x27\x26\xa4\x0c\x4c\xec\x4c\xa0\xaa\x83\x63\x5f\x1c\x93\xd8\x67\xce\xe7\xaa\xfd\xf7\x28\x4d\xa0\x85\x52\x86\x6f\x89\xa2\xd7\xef\xf3\xbc\xba\x1d\x7c\xa6\x3c\x7a\x25\x14\x27\xc0\xb3\x89\x79\xc2\xed\x66\xb7\xdd\xec\xe0\x0b\xf6\xc8\x20\x04\x83\x48\x2e\x1f\xb5\xf6\x34\x99\xe4\x95\x0f\x32\xd4\x4e\x05\xd2\x0e\xb3\x76\x64\x8b\xbe\x39\xd4\x20\x71\x9a\xe9\x9e\x18\x1c\x8a\x09\x13\xba\xfb\x0d\x47\xb6\x46\x4c\x62\x24\x50\x52\xeb\x14\xe3\x8f\x1a\x18\x1d\xb4\x70\x68\x56\xbf\xa5\xa8\x6b\x41\xd6\x32\x84\xe4\xb5\x8d\x6e\xf9\x6b\x8e\x33\x11\x7c\xa2\x67\x40\x66\xfa\x8e\x56\x74\x1e\xfd\xd7\xe6\x1d\x34\x5d\x90\xae\xda\x11\x45\x11\xfb\x87\xca\xa7\xf9\xf3\xed\x6a\x9d\xc5\x87\x83\xa5\x54\x84\x4d\x48\x72\xbc\x4e\x01\x24\x13\x11\x5a\x78\x36\xd4\x2c\xad\x13\x72\x09\x94\xe6\xe2\x7b\xf5\xa2\x5e\x9a\xb7\x1b\x5f\x57\x65\xc7\x26\xd9\x61\x2e\x3a\x3c\xad\x59\xa1\xca\xf6\x6f\x38\x5e\x7a\xe2\xf1\x0f\x7c\x99\xa6\x13\x32\x07\x87\xff\x19\x3e\xeb\xcb\xe3\x05\xaf\xea\xc3\xdd\x05\x99\x6b\xc2\x5f\x06\x4a\x7b\x4f\xd0\x42\x6f\xa6\x82\x4b\xe6\x69\x2f\x58\xa4\x40\x0b\xc2\x75\x0d\x6b\xaa\x05\xdd\x3e\x1b\x3b\x1a\x8f\xb7\xb7\xed\xe6\xb7\xef\x1f\xe4\x73\xee\x67\x00\x00\x00\xff\xff\xfc\xcc\xd9\x95\xab\x02\x00\x00")

func assetsDefgopkgTomlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDefgopkgToml,
		"assets/defGopkg.toml",
	)
}

func assetsDefgopkgToml() (*asset, error) {
	bytes, err := assetsDefgopkgTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/defGopkg.toml", size: 683, mode: os.FileMode(438), modTime: time.Unix(1521189328, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaMashling_schema02Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\xbd\x6e\xdb\x30\x10\xde\x03\xe4\x1d\x08\x25\x63\x12\x76\xe8\xe4\xb1\x9d\x3a\x25\x40\xbb\x15\x41\x40\x4b\x27\x89\xa9\x44\xaa\xe4\x39\x81\x11\xf8\xdd\x0b\xfd\x5a\x54\x49\x4a\xae\x25\xd7\x01\xe2\xc9\xe0\xfd\x1f\xbf\xfb\x48\xea\xed\xf2\x82\x10\x42\x82\x6b\x1d\xa6\x90\xb3\x60\x45\x82\x14\xb1\x58\x51\xfa\xac\xa5\xb8\xad\x57\xef\xa4\x4a\x68\xa4\x58\x8c\xb7\x9f\x3e\xd3\x7a\xed\x2a\xb8\x69\x4c\x71\x5b\x40\x69\x27\xd7\xcf\x10\x62\xb7\xac\xe0\xf7\x86\x2b\x88\x82\x15\xf9\x59\x2f\x55\xcb\x39\xd3\x69\xc6\x45\xf2\xd4\x04\xbc\xe9\xc9\x12\x86\xf0\xca\xb6\x41\xbd\xf4\xd8\x7a\x2a\x94\x2c\x40\x21\x07\x1d\xac\xc8\x9b\xc7\x97\x21\x35\x72\xd3\xa8\xb8\x48\x82\xbd\x74\x67\x8b\xeb\xb6\x37\x6b\xeb\xc4\xf6\x1a\x3b\xb1\x60\x39\x0c\x6d\x2a\xc1\x0b\x28\xcd\xa5\xb0\xca\x42\x29\x62\x9e\x6c\x14\x43\x2e\x85\xb6\xaa\xa0\xe2\x49\x02\xca\x2e\x84\x17\x10\xf8\x94\x32\x11\x65\x7e\x95\x8c\x8b\x5f\x3a\x30\xc5\x8f\xc3\x0a\x5d\xbd\x37\x6b\xb4\x8a\x88\x77\x03\xda\xdf\xce\xd7\x9f\xb9\xfd\x46\x5c\x17\x19\xdb\x3e\x2d\x92\x74\xeb\x9c\xe7\x2c\x59\xc0\x3b\xe8\x50\xf1\x02\x97\x68\xcb\x00\x72\x6e\xf7\x1c\x21\xf7\xc8\x2b\x9d\x6b\x05\x71\x99\xc2\x15\x8d\x20\xe6\x82\x57\x3e\xa9\x11\xc2\x92\x97\x2b\x37\xa3\x2a\xa6\x54\xc7\x0e\xa3\x45\x0d\xe6\x60\x89\xa2\x8c\x10\x27\x2c\xaa\x9e\xdc\xe5\x2a\x2a\xfd\x9f\xa4\x9c\x8e\xc8\x96\xa8\xa5\x71\x3e\x7b\x21\xe6\xd2\xd0\x4f\xc0\xa2\xa8\x4a\x80\x65\x0f\x7d\xee\x8c\x59\xa6\xa1\x77\xfe\xd4\x7f\x5b\xeb\xa0\x97\xf8\xe0\x90\x33\x27\xe7\x64\x47\x54\xe5\xd6\x26\xd0\x80\xc8\x45\x72\xa6\x27\x47\x63\xf5\x7e\xb8\xb7\x6b\xe7\xb8\x63\xfb\x0e\x77\x6a\x0e\xe0\xf9\x06\x67\x34\x65\xf2\x37\xe0\x2d\x4b\xff\x3a\x03\xfd\x3b\x58\x3b\xad\x1f\x00\x27\xb6\xfe\x0c\xad\x3e\x00\xfe\xde\x00\x6e\x5e\x16\x16\x81\xf9\x08\x60\xa5\x80\xfb\xd8\x6e\xef\xda\x19\x7f\xd4\x9e\x5a\x0c\x0a\x44\x38\x4c\xa1\x4b\x65\x1a\x56\x8e\xcc\x62\x7f\x86\x1e\x90\xc6\x59\x0e\xf9\x92\xf3\x58\x30\xc5\x7c\x37\xaa\x73\x9c\x46\x5b\x1d\x7b\xcc\xcd\xdf\xfd\x0e\x47\x47\x77\x69\x2a\x74\x19\x32\x97\x8f\x4a\x81\x47\x5e\x71\x79\x01\x76\x60\xde\x95\xd9\x18\xb2\xcd\xdc\xbc\x2a\xc4\xd2\x10\xb7\xb6\xeb\xe6\x4d\x9a\x3a\xa7\x87\xf2\xe2\x68\x34\x54\xfd\x68\x38\x30\x96\xc7\x21\x69\xb0\xc3\x36\x19\xb6\xdf\xd1\xf4\x8a\xd2\x84\x63\xba\x59\xdf\x85\x32\xa7\x3f\xbe\x7d\xf9\x7a\xff\x5d\xc6\xf8\xca\x14\xd0\x38\x93\x89\xbc\x0d\xa5\x40\xc5\xd7\x74\x9d\xc9\x35\xcd\x99\x46\x50\x94\x85\x25\xfa\x4a\x85\xd7\xde\x93\xa6\xf9\xf2\x76\xf7\xac\x9d\x0c\x47\xec\x03\xe4\xeb\xc5\xc4\xb3\xcc\xe1\x7d\xbe\x93\xb1\x7a\x74\xce\x7d\x2c\x7a\xbf\x94\x45\x5c\x17\x0c\xc3\x14\x8e\xbf\xea\x4d\x78\xc8\x9a\xef\x4a\xd7\x66\x4c\x79\xee\xce\x4a\xa3\xbd\x26\x9c\x32\x77\x3f\x6d\x92\xe9\xd4\x59\xa9\xfa\x3f\xc1\x10\x0f\x0d\x92\x43\xa8\xb0\x52\xe6\x13\x58\x83\x1c\xc4\x52\x64\x84\xa9\x48\xbf\xc4\xff\x11\x9b\x8b\x62\x83\x0f\x23\xf7\x06\x4b\xfc\x09\xbb\xdc\x59\x1c\x7e\x9d\xb0\x05\x9c\x54\x2f\x71\x93\xe4\x04\xb1\x8b\x5e\x17\xa3\xca\xfa\xef\xee\xf2\x62\x77\x79\xf1\x27\x00\x00\xff\xff\x36\x63\x4c\x69\xaf\x19\x00\x00")

func schemaMashling_schema02JsonBytes() ([]byte, error) {
	return bindataRead(
		_schemaMashling_schema02Json,
		"schema/mashling_schema-0.2.json",
	)
}

func schemaMashling_schema02Json() (*asset, error) {
	bytes, err := schemaMashling_schema02JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/mashling_schema-0.2.json", size: 6575, mode: os.FileMode(438), modTime: time.Unix(1519905570, 0)}
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
	"assets/banner.txt": assetsBannerTxt,
	"assets/defGopkg.lock": assetsDefgopkgLock,
	"assets/defGopkg.toml": assetsDefgopkgToml,
	"schema/mashling_schema-0.2.json": schemaMashling_schema02Json,
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
	"assets": &bintree{nil, map[string]*bintree{
		"banner.txt": &bintree{assetsBannerTxt, map[string]*bintree{}},
		"defGopkg.lock": &bintree{assetsDefgopkgLock, map[string]*bintree{}},
		"defGopkg.toml": &bintree{assetsDefgopkgToml, map[string]*bintree{}},
	}},
	"schema": &bintree{nil, map[string]*bintree{
		"mashling_schema-0.2.json": &bintree{schemaMashling_schema02Json, map[string]*bintree{}},
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

