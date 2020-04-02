// Code generated by go-bindata. DO NOT EDIT.
// sources:
// assets/cluster.tf (8.964kB)
// assets/outputs.tf (327B)
// assets/variables.tf (560B)

package assets

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _clusterTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\x6d\x8f\xdb\xb8\x11\xfe\xee\x5f\x31\x55\x02\x5c\x7b\x28\x25\xef\x4b\x36\xbb\x3d\x6c\x8b\xf4\x0d\x28\x70\xd7\x3b\xa0\x07\xf4\x43\xb0\x10\x28\x72\x2c\x13\x96\x48\x96\xa4\xbc\xab\xa4\xdb\xdf\x5e\x90\x92\x2c\x59\x96\xbd\x8e\x73\x9b\x97\xa2\x9b\x2f\xc6\x70\x38\x33\x7c\x66\xe6\xd1\x90\xd1\x46\xad\x05\x47\x03\x11\xbd\xb7\x11\xbc\x9f\x01\x68\xa3\x16\xa2\x40\xb8\x85\xa8\x2c\x49\xa1\x28\x77\x68\x5d\x34\x03\x30\x98\x0b\x25\xc1\xaf\x54\x96\x20\xb5\x8e\x9c\x7b\xf9\x1a\x8d\xf5\x0b\xb7\x10\xfd\xe7\xf7\x70\x1e\x5f\xbe\x8e\x66\x8f\xb3\x99\x41\xab\x2a\xc3\x30\xd8\x4e\x57\x58\xa7\x9a\x0a\x13\x41\xb4\xc2\xba\x71\xe5\x65\x92\x96\x08\xc1\xe6\xcb\xf7\x6b\x6a\x62\x56\x54\xd6\xa1\x09\xf2\x47\xb2\xc2\x3a\x6c\xf2\x71\x55\x59\x21\x98\xb7\x03\xb7\xe0\x23\xfc\xb5\x57\xb7\x76\x99\xf6\x2b\xbf\xd9\xf5\x2b\xa4\x75\x54\x32\x8c\x20\xa2\x5a\xa7\x16\xcd\x1a\x4d\xe3\xde\xd1\xdc\xc2\x6d\xf8\x09\xf0\x77\x1f\xc7\x9e\x28\xa8\xd6\xe4\xe5\x7b\xa6\x2a\xe9\x62\x21\x39\x3e\x3c\xfa\x80\x1e\x67\x33\x00\xa6\xa4\x44\xe6\xfc\xf1\x1b\x3b\x2f\xe0\xe7\x25\x02\xc7\x05\xad\x0a\x07\x95\x45\x13\x4e\xb8\x50\x06\x54\x65\xe0\xcd\x0f\x7f\x0b\x6a\xae\xd6\xc1\x9d\xb5\xcb\x28\x08\xbc\x66\x40\x36\xab\xa4\xab\x1a\xd9\x52\x59\x07\xb7\x60\xb1\x58\xc4\xed\x21\x85\xee\x3c\xd3\x52\x40\xff\x77\x0b\x11\x2d\x05\x99\x2f\xd8\xf9\x9c\xf3\x33\x4e\x2f\xe7\x57\xaf\xaf\xe7\x59\x04\x2f\xe0\xec\x3a\x9e\x5f\xc2\xf7\x3f\xff\x63\x06\xd0\xc1\x91\x76\x01\xb0\x57\xf1\x43\x41\x4d\x8e\x11\x0c\xff\x5e\xc0\x9b\xe2\x9e\xd6\xd6\xc7\x05\x1b\x9d\x70\x0c\xa6\xa4\x15\xd6\xa1\x64\xf5\x76\x0e\x43\x18\xc3\x5c\xc7\x2b\xac\x63\xc1\x03\x4c\x95\x74\x83\x60\x3d\xc6\x3e\x1d\x9b\x70\x82\x82\xaf\x25\xcd\x52\x8b\xac\x32\xc2\xd5\x69\x6e\x54\xa5\x53\xc1\x7d\x96\xde\x06\x44\xa2\x97\xef\xbd\x83\x6d\x0d\x6f\x29\x16\xfc\x31\xfa\xed\x61\x9d\x34\x57\xd6\x8a\x46\xd5\x23\x78\x37\x6b\xca\x7d\x2d\x7c\xf9\xfa\x26\xf0\x65\x15\xb5\x89\x6c\x8b\x68\x10\x70\x49\x9d\x43\x53\x2a\xeb\xd2\x42\x30\x94\x16\x53\xbf\x21\x68\x73\xb4\x4e\x48\xea\xda\x3e\x48\x96\xaa\xc4\xa4\xc9\x65\xd2\xef\x1b\x98\x20\xad\x89\x4d\x25\x6d\x05\x62\xb0\x54\x0e\x09\x3e\x20\xeb\xe2\x11\xb2\x10\x12\x37\x48\x00\x44\xf7\x4b\xdf\xa7\x6f\xe1\x57\x40\x16\x90\xac\xa9\x49\x0a\x91\x25\xac\x50\x15\x4f\x3a\x60\x93\x4c\x29\x47\x16\x42\x0a\xbb\x44\x0e\x77\xdf\x01\x57\x80\x6c\xa9\xe0\x9b\x7f\x52\xe1\x84\xcc\x9b\x9c\xfa\x4d\x44\x48\xe1\xe2\x38\xfe\xe6\x3b\xb0\x05\xa2\x86\x33\xaf\x2d\xb1\xc5\xd5\x7b\xcc\xd1\x01\x21\x52\x11\xb6\x44\xb6\x22\x0c\x8d\x13\x0b\xc1\xa8\x43\x20\xff\xfa\x11\x08\x2c\x9d\xd3\xf6\x77\x49\x62\x2f\x08\x56\xe4\x1e\xad\x23\x67\x31\x2d\xe9\x3b\x25\xe9\xbd\x8d\x99\x2a\x13\x8e\x59\x6c\x54\x56\x59\xa7\xd1\x30\xd4\x1e\xb3\x58\xa8\xe4\xf2\xec\x2f\x7f\xfd\xd3\xcd\xcd\x9f\xe3\x5c\xe7\xf0\x6f\xb0\x15\x57\x40\xb5\xf3\x1c\x00\x94\x73\x20\x7d\x1c\x9b\xb5\x10\x4f\x0d\x95\xe6\xd4\xe1\x9e\xf5\x00\x45\x51\x78\x3d\x6d\x54\x89\x6e\x89\x95\x25\x52\x71\x8f\xaf\x56\xc6\xa1\x19\x9d\xb0\xe1\x80\x41\xba\xb9\xba\x97\x9e\x09\xd3\xca\x14\x8f\xbd\xb2\xa3\x06\x1e\xde\x2d\x60\x90\xd5\x6f\x63\x47\x4d\x9c\xbf\x1b\xc5\x52\xae\x07\x4a\x90\x28\xed\x92\xb1\xc6\x8a\x0b\x03\x44\x37\x8b\xbd\x72\xc2\xa9\xa3\x0d\x1d\xdc\x85\x52\x39\xc4\x70\x25\x3a\x23\x98\x3d\x8d\xe5\xda\xcd\xff\x3b\xd4\xe6\xce\xe3\x86\xd9\x8e\xa4\xa9\x53\xe9\xa7\x05\x6e\x43\x41\xbb\xbc\xf2\xff\x76\x7e\xe6\x76\x1e\xe9\xda\xda\x3a\x2c\x99\x2b\x00\x25\xcd\x0a\xdc\xaf\x39\x61\x95\x72\x1e\x2a\xb6\x10\xd9\x42\x49\xc7\x94\x5c\x88\xfc\x6c\x84\x5a\x07\x0c\x2f\xe2\xdc\xd0\x05\x95\x34\x80\xa1\xac\x4d\x0c\x16\x48\x2d\x26\xad\x3c\xbd\x8a\xaf\xe2\xf3\x94\x96\xfc\xea\x32\xe6\x98\x8d\x02\xe0\x7a\x95\x03\x11\x70\x9c\x76\x7f\x30\x4e\xb1\x54\x92\x18\xf4\xbc\xf4\xd4\xf1\x5b\xe3\xa4\xe5\x85\x91\x36\x9a\xb5\x60\x63\x25\xb0\x8e\x1a\x77\x2c\xf1\x68\xa3\x1e\xea\xd3\x68\x27\x6c\x6d\x48\x67\xdc\xf9\xdb\x7f\x7b\x79\x60\xa7\xf7\x77\xf6\x95\x97\xdd\x90\xe3\x7d\x58\xab\x98\xa0\x0e\xd3\x0d\xf1\xa4\x94\x73\x83\xd6\x07\xec\x4c\x85\xa7\x12\x41\x38\x4a\x33\x5e\x34\xa0\x6d\xb3\xce\x41\xe6\xf9\x7c\x6c\xfb\x35\x30\xd5\xa9\x0c\x21\x73\x21\x1f\x7e\x99\x1e\x9a\x32\x65\xca\x80\x00\x3a\x96\x84\xe5\xc4\x0a\x87\x96\x34\x3b\x78\xd2\x26\x6f\xb4\xa9\x90\x40\x16\x76\x77\x17\x5d\x53\x51\xf8\x8d\xc9\x70\x54\xd8\x67\xbb\xd7\x19\xf5\xe8\xa8\x49\x0d\xb7\x69\xd7\x70\x83\x86\x1d\x8b\xda\xcb\xdf\xf6\x9c\xde\x56\xac\xef\x59\x9e\xed\xce\xea\x82\xa3\xf4\xdf\x0b\x34\xc3\x5e\x9b\xea\x71\x9e\x4d\xdc\x9f\x36\x21\xf4\x66\x9a\xee\x18\x84\xec\xfd\x76\x3f\xc3\x45\xa2\x0f\xa2\xa0\xd6\xee\x0b\xd0\xaf\xcd\x00\xd0\x03\x87\x4f\x9c\xa6\x51\xf2\xbc\xa0\x75\x51\xa7\xa2\x2c\x91\x7b\x76\x28\x6a\xe8\xe8\x60\x87\xf9\x06\x11\x46\x10\xf5\x31\xb6\x20\xee\x1e\x6c\x3f\x30\x1e\x09\x3f\xe2\x65\xd4\xe2\x90\x2b\xa6\x37\x34\xfa\x25\x0d\x92\x0d\x27\x6c\x1d\xac\x93\xf6\x7a\x9a\x5a\x7b\xaf\x0c\xdf\xd6\xeb\xa4\x33\x00\xbb\x12\x3a\x5d\x08\x49\x8b\xd4\x4a\xaa\xed\x52\xb9\x9e\x09\x27\x60\xe9\x17\x27\x10\x3e\x04\x71\xf3\x23\xed\x1e\x0c\xb6\xd5\xb7\x17\xdf\x4e\x1b\xb9\x3b\x3c\xa5\x4d\xf3\x32\xcf\x02\x29\xdf\x1d\xfc\x84\x75\x2f\x1d\x29\xcd\x51\xba\x0f\x7d\x21\xf0\x7b\x8e\x7e\x23\xf8\x64\x13\xf2\x9e\xa9\xd8\xa7\xb2\x2a\x8f\x1d\x8b\xa7\x6f\xef\xdb\x70\xb5\x9c\x70\x42\x6e\xc2\xf6\x36\x3d\x47\x7c\x8c\xbe\x8e\xcf\x51\x33\x58\xff\x08\xae\xd4\xed\x7d\xb0\xbd\x53\x86\x4b\x64\x80\xed\x89\x2b\x65\xbf\xb3\x5f\xdb\xba\x44\x86\x97\x39\xe2\x4d\x11\x99\x7f\xbb\x6f\xa1\xdf\x6c\xca\xa1\xcd\xe6\x93\x01\xd3\x73\xdd\x76\x86\x9a\x87\xb3\x06\xff\xe1\x2c\x73\xe0\xad\xac\x33\x40\x1a\x03\xb3\xf0\x3c\xc2\x8c\xd0\xdd\xf3\xc8\x1b\xad\xa1\x53\x82\xa0\x14\xd0\xef\x6a\xaa\x63\x50\x98\x70\x10\x5e\x6e\x84\xcc\xc3\xb8\xd6\xd4\xc4\xc2\xa8\x32\xf5\x17\xf9\x10\xd5\xf9\x79\xd3\x63\xaa\x13\x0d\x84\xda\x28\xa7\x98\x2a\xda\xf8\x1d\xd3\x0d\x12\x4c\x70\x93\x66\x85\x62\xab\xa6\x52\xe7\x71\xf8\x97\xcc\xa3\xbb\x76\x30\x3d\xe4\xf1\x7a\x7e\xf5\x6a\xc2\xe7\x46\xfc\xcb\x7b\x0d\xc6\x5f\x8f\x7c\x0e\x84\xbd\xc7\xa1\xbf\x17\xf0\x03\xad\x33\x04\x83\xd6\x5f\x58\x1d\x28\x59\xd4\xc1\x2a\xfc\xb4\xb9\x1f\x41\x3b\xfe\xff\xa1\xdd\xf2\xc7\xca\xc1\x92\x4a\x5e\x43\xd3\x66\x8e\xae\x7c\xb3\xb4\xaf\xc4\x16\xee\x85\x5b\xaa\xca\x41\x49\x65\x45\x8b\xa2\x06\x6b\x97\xc4\x6b\x08\xe9\x14\xb8\x25\xb6\x06\xe3\x8f\x06\xba\x81\xef\xe6\x6c\x3e\xdf\x01\x7b\xb4\x34\x04\x7c\x0c\xfa\x76\x71\x1f\x20\xa6\xe1\xad\xbe\x0b\x0e\xf7\x17\xc1\x6e\x50\x9d\x6c\x94\x7e\x72\x76\x5c\xf6\x8f\x6a\xcb\xf6\x49\xf3\x23\xba\x93\xb4\x16\x8e\x6e\xd2\x46\xff\xc8\x5e\x7d\x3a\x9f\xd7\xf3\xd7\x97\x7b\xf2\xb9\x59\x9a\xc8\x67\xc5\x3f\x34\x9f\xdd\x23\xf1\xb1\x85\x76\x62\x60\x1f\x5e\x68\x27\x04\xb6\x4b\x37\xa3\xa5\xcf\x85\xd8\x29\x81\x7d\x3c\x62\x9f\xb8\x2d\x9f\xea\x4b\x9e\x0d\xfa\x71\xff\x8d\x68\xfc\x95\x3c\x02\xe1\x8b\x8b\xf9\xd5\x1e\x84\x37\x4b\xcf\x80\xf0\x11\x91\xbd\xba\xbc\xd8\xfd\xee\x8e\x96\x9e\x21\xb2\x23\x48\xb2\x9f\xe6\x8f\xe1\xc7\x30\xc7\x3f\x39\xbf\x7c\xdf\x91\x5f\xd0\xff\xaa\x46\x99\x83\x53\xc5\xe5\x7c\xe7\xe3\x3a\x10\x4e\x4f\x15\xfe\x82\xd2\x61\xda\xde\x0a\x83\x97\x2f\xeb\x5b\xd9\xfd\xe7\xc6\x93\x8d\xd9\x2a\x7e\x58\x77\x7e\x8e\xa9\xf3\x66\x7e\x33\x85\xe3\x46\xfc\x3c\x5e\x2f\x76\x2b\x64\x4b\x7c\xba\xd7\x2f\xab\x5e\x9a\x57\xe9\x63\x69\x23\x68\x3f\x4d\x1b\x3f\x79\xb5\x67\x63\x8b\xeb\x29\x98\xae\x3f\x2e\x2d\x9f\x87\xa0\x3e\x6d\x29\xfc\x37\x00\x00\xff\xff\x9f\xb0\xf9\xeb\x04\x23\x00\x00")

func clusterTfBytes() ([]byte, error) {
	return bindataRead(
		_clusterTf,
		"cluster.tf",
	)
}

func clusterTf() (*asset, error) {
	bytes, err := clusterTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cluster.tf", size: 0, mode: os.FileMode(0644), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc7, 0x2d, 0x85, 0x1c, 0xaf, 0xd2, 0x42, 0x9d, 0x35, 0x27, 0x16, 0xdc, 0x6b, 0x2a, 0xd4, 0x4f, 0xd4, 0x4e, 0x29, 0xc4, 0x7f, 0x93, 0x56, 0xdc, 0x6, 0x46, 0x41, 0x1, 0x9a, 0xbb, 0x31, 0x2a}}
	return a, nil
}

var _outputsTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x8f\x41\x0a\xc2\x30\x10\x45\xf7\x39\xc5\x10\x5c\xb9\xe8\x0d\x5c\x79\x04\x0f\x10\xa6\xcd\x20\x85\x98\x84\x99\x49\x55\x4a\xef\x2e\xc4\x06\xac\x48\xdd\x05\xfe\xfb\x2f\x7f\x52\xd1\x5c\x14\xec\x18\x45\x31\x0e\x24\x16\x66\x03\x30\x61\x28\x04\x27\xb0\x87\x19\xef\xe2\x5a\xda\x61\xce\x4e\x88\x27\xe2\xee\xb8\x58\xb3\x18\xd3\x04\xbe\x3f\x87\x22\x4a\xfc\x53\xc0\x5e\xdc\xf0\xce\x3b\xdf\xb7\xe7\xd6\x80\x57\x8a\xfa\xe7\xff\x90\xd0\x2b\x89\xba\x0a\x7f\x6f\xb8\x91\xf2\x38\xc8\xa5\x0e\xdc\x17\xad\xe8\x7a\xcc\x56\x93\x39\x3d\x9e\xfb\xf5\x8a\x7c\x94\x5f\x01\x00\x00\xff\xff\x6f\xd3\xe1\x13\x47\x01\x00\x00")

func outputsTfBytes() ([]byte, error) {
	return bindataRead(
		_outputsTf,
		"outputs.tf",
	)
}

func outputsTf() (*asset, error) {
	bytes, err := outputsTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "outputs.tf", size: 0, mode: os.FileMode(0644), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd5, 0xd7, 0xa7, 0xff, 0x49, 0x7d, 0xf7, 0xf2, 0xf0, 0xad, 0x3d, 0x63, 0x4d, 0x52, 0x40, 0x3f, 0x63, 0xb2, 0x30, 0x39, 0xfe, 0xb0, 0xbb, 0x6b, 0x82, 0xfe, 0xd3, 0xf7, 0xce, 0xb2, 0x30, 0x67}}
	return a, nil
}

var _variablesTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x91\x41\x4e\xc3\x30\x10\x45\xf7\x39\xc5\xc8\x7b\x2c\xd4\x0a\x10\x8b\x9c\x65\x34\x89\xa7\xc5\xc2\xb1\x8d\x67\xdc\x2a\x42\xbd\x3b\x8a\xbb\x80\xca\x05\xb2\xcc\x7b\xff\xeb\xcb\x73\xa2\xe2\x69\x0a\x0c\x66\x0e\x55\x94\x0b\x46\x5a\xd8\xc0\xe7\x70\x19\x86\x6f\x48\x39\xa3\x8f\xa2\x14\x67\xc6\x39\xd5\xa8\x9d\x12\x12\x39\x65\x51\xa4\x23\x47\xfd\x45\x72\xd3\x7f\x35\x3f\x0d\x8e\x47\x1f\xfb\x31\x37\x25\x81\x44\xee\x19\xd7\x2c\x9e\xb8\x88\x4f\x71\x33\x00\x74\xcd\x0c\x23\x2c\x94\x07\x00\xc7\x07\xaa\x41\x61\x6c\x08\xc0\x50\x2d\xa9\xd0\xc3\xb2\xca\x47\x30\xd0\xbe\x11\xcc\x93\x7d\xb1\xed\x17\x5e\xb9\xdd\xd9\xc7\xbd\xdd\x99\x9b\x4c\x4e\xa2\xc7\xc2\x2d\x38\x82\x79\xb5\xcf\x76\xbf\x19\x97\x6e\x56\x15\x2e\x77\x5f\xd8\x4d\x98\x49\xe4\x9c\x8a\xeb\x98\xc8\x1b\xe6\x3a\x05\x3f\xe3\x3b\xaf\x1d\x5e\x48\x95\xcb\x92\x44\xd1\xa5\x73\xdc\x0e\x81\xb5\x84\xbf\xbc\xe0\x67\x8e\xc2\x78\xf0\xa1\x9f\xd2\x0a\xda\x29\xfb\xba\xaf\x00\x00\x00\xff\xff\xb4\x82\x23\x39\x30\x02\x00\x00")

func variablesTfBytes() ([]byte, error) {
	return bindataRead(
		_variablesTf,
		"variables.tf",
	)
}

func variablesTf() (*asset, error) {
	bytes, err := variablesTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "variables.tf", size: 0, mode: os.FileMode(0644), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xbe, 0xa1, 0xc8, 0x99, 0x32, 0x37, 0x1c, 0xdc, 0xb1, 0x38, 0xcb, 0x56, 0x9d, 0x23, 0xd7, 0xdd, 0x79, 0xb1, 0xd5, 0xda, 0x30, 0x2a, 0x2d, 0xb9, 0x4b, 0x8c, 0xef, 0x78, 0xdc, 0x9c, 0xc3, 0x2b}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"cluster.tf":   clusterTf,
	"outputs.tf":   outputsTf,
	"variables.tf": variablesTf,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"cluster.tf":   &bintree{clusterTf, map[string]*bintree{}},
	"outputs.tf":   &bintree{outputsTf, map[string]*bintree{}},
	"variables.tf": &bintree{variablesTf, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}