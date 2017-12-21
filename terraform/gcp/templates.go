// Code generated by go-bindata.
// sources:
// templates/bosh_director.tf
// templates/cf_dns.tf
// templates/cf_lb.tf
// templates/concourse_lb.tf
// templates/jumpbox.tf
// templates/vars.tf
// DO NOT EDIT!

package gcp

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

var _templatesBosh_directorTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x95\x51\x73\xa2\x30\x10\xc7\xdf\xf9\x14\x3b\x99\x7b\xac\x5e\xcb\x81\xfa\xd2\x4f\xd2\x71\x98\x00\x2b\x97\x6b\x24\x4c\x48\xb4\x37\x1d\xbf\xfb\x4d\x08\x01\xac\x48\x83\x57\xa7\xfa\x00\xe2\xee\x9f\xff\xfe\x76\x17\x84\x56\x95\x56\x40\x4a\x54\x47\x21\x5f\x93\x92\xee\x91\xc0\x7b\x00\x70\xa0\x5c\x23\x3c\x03\xf9\xf1\x5e\x08\x51\x70\x4c\x32\xb1\xaf\xb4\xc2\xa4\x8d\x5d\xa6\x29\x5f\xb8\x73\x93\x77\x22\xc1\x29\x08\x9c\x62\xad\xd3\x39\xa2\x7d\x78\xa3\x6b\x7f\x8e\xc8\xa6\xa2\xfe\x9d\x88\x0a\xcb\x44\xd1\xc2\x4b\x79\xc7\x24\x1e\x29\xe7\x4b\x93\xba\x30\xa9\xd7\x64\x73\x26\x31\x53\x42\xde\x28\xed\xd2\x47\xe4\xff\xe8\x7d\x95\x8a\xb7\xeb\xc2\x07\x2a\x97\x58\x1e\x12\x96\x9f\x16\x6d\xf0\x99\x00\x2b\x15\xca\x92\xf2\xf9\xd6\x5c\xe6\xc0\x95\xc4\x5a\x68\x99\x21\x90\xf1\xce\x12\x20\x83\xde\xda\x3b\x99\x6c\x18\xf9\x5c\xb8\x77\x59\x01\x00\xd5\x4a\x24\x99\x44\x7a\xd6\xdf\x1a\x9e\x61\x47\x79\x8d\xa3\xf5\x65\x2c\x97\xa3\x74\xac\x40\xf3\xbf\xad\xe2\x40\x25\xa3\x29\x47\x37\x6a\x83\x54\xf5\xb7\xc2\xd6\x5c\xad\x24\x2b\x0b\xe3\x26\xc7\x1d\xd5\x5c\x99\x8b\x4f\x8f\xcb\xe6\xfb\xf3\x69\x35\x0d\xa4\x77\xdd\x32\xb1\x17\xc6\x90\x5c\x80\x68\x43\x03\x00\x56\x35\xde\x12\x49\xcb\xe2\x6a\x45\x00\xed\x9d\x86\x7a\x1e\x9b\x57\x23\xdf\x25\x9c\x95\xaf\x9f\xf4\xd6\x0d\x04\x01\x82\x6f\x16\xf6\x79\x19\x17\x05\x74\x61\x03\x6f\x33\x9f\x07\x01\x80\xb5\x63\x6b\x37\xad\x7f\x21\x0e\xfe\x23\xd9\x9a\x00\xca\xb9\x38\x36\x4e\x00\x2a\x21\x55\x6d\xcd\xbc\x90\x30\x24\x0f\x40\x56\x9b\xd5\xc6\x1c\xc3\x38\x8e\x63\xb2\xb5\x61\x52\x28\x91\x09\x6e\xec\xa8\xac\x32\x06\x4f\x46\x4a\x51\x59\xa0\x32\x3b\x62\xef\x74\x5e\x4f\xb7\xff\x64\xeb\x4b\xaa\x4f\x99\x46\xd5\xc7\x7d\x05\x2b\x0f\xff\xfe\xdc\x36\x51\xf4\xab\x39\x6e\xa2\xe8\x0b\x39\xba\x87\xdd\x4c\x96\x5d\x9a\x07\xcf\x2e\xf6\xde\x4c\x07\xb5\x7c\xe4\x7a\x13\x20\xf7\x30\xf3\x67\xe3\x32\x16\x4a\xf8\x22\x1a\x4d\xb9\x23\xa9\x41\x51\x53\xc3\x17\x85\x76\xfc\xc2\x38\x8c\x1f\xed\xc9\x7a\xbd\xfe\x8e\x79\x6b\x5f\xa3\x86\x4f\x73\x61\x92\xe6\x87\xe0\x3b\x72\x74\x6f\xf7\x4f\x77\xf8\x7f\x90\x75\xcd\x7a\x98\xde\xac\xd9\x03\xea\x39\x94\xdf\x34\x88\x03\x56\x2c\xdb\xf7\xb0\x7c\x56\xfa\x5a\x8c\xce\x6f\x5a\xfb\x7f\x01\x00\x00\xff\xff\xc1\x20\x35\xc9\x59\x0b\x00\x00")

func templatesBosh_directorTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesBosh_directorTf,
		"templates/bosh_director.tf",
	)
}

func templatesBosh_directorTf() (*asset, error) {
	bytes, err := templatesBosh_directorTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/bosh_director.tf", size: 2905, mode: os.FileMode(480), modTime: time.Unix(1513818666, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesCf_dnsTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x95\x3f\xcf\x9b\x30\x10\xc6\x77\x3e\xc5\xc9\xea\x54\x09\xf4\x4a\x9d\x33\x54\xea\xdc\xa5\x63\xf5\x0a\x39\xf6\x85\x20\x19\xdb\xba\x3b\x48\xd3\x57\x7c\xf7\xca\x04\xf2\xa7\x0d\x6a\x18\x22\x31\x24\x4b\x00\x3f\xbe\xe7\xb9\x9f\x11\xd7\x69\xaa\xf5\xd6\x21\x28\x3e\xb2\x60\x53\xda\xd0\xe8\xda\x2b\xf8\xc8\x00\xe4\x18\x11\x36\xa0\x58\xa8\xf6\x95\xca\xfa\x2c\x23\xe4\xd0\x92\x41\x50\x55\x08\x95\xc3\xd2\x7a\x2e\x1b\xed\x75\x85\xb6\xfc\x1d\x3c\x2a\x50\xe8\xbb\xe1\xf1\xe9\x36\x15\xf2\xba\x41\x18\x7f\x1b\x50\x9f\x3e\x3a\x4d\x45\x92\xd5\xb6\xcf\x07\x59\x06\x90\xb6\x4c\xc2\xb3\xe8\x26\x55\x5f\x0c\x3a\x64\x43\x75\x94\x3a\xf8\xa4\xfb\xf6\xfd\x07\xa4\x12\xb0\x0b\x04\xb2\x47\xb8\xa9\x0e\xe8\xbb\x9a\x82\x6f\xd0\xcb\xd0\x40\x68\x25\xb6\xf2\x57\xbb\x43\x5c\x46\xea\x90\xf8\x94\xb8\xd3\xae\xc5\x53\x8c\x99\x46\x8b\xeb\x36\x8b\x14\x7c\xaa\xd0\xcf\x93\x22\x34\x81\x6c\xc9\x28\x0a\xd4\xa1\x76\xd6\x68\xb2\xb9\xf5\xfc\x0f\xa7\x0d\xa8\xcf\xc5\x83\xe6\x13\xb9\xfe\x84\x27\xa2\xb7\x5c\x0e\x74\x7e\x4e\xe6\x26\x34\xb1\x15\x2c\x2b\x17\xb6\xda\x95\xda\x5a\x42\xe6\xc2\xec\xf2\xf1\x52\xbd\x4f\x07\x7e\xf6\xff\x9a\xca\x89\xb8\xcb\xc9\x7d\x79\x7b\xcb\x32\x80\xeb\x24\x0b\x19\xf5\x2a\x15\x20\xb2\x5a\x34\x0f\x01\xcf\x9b\xff\x1b\xb1\x18\xff\x7b\xf5\xfe\x18\x60\xb3\xcb\x99\xf7\x79\xa4\xf0\xeb\x78\x0f\x30\xf3\xfe\x09\x88\xaf\x82\x5f\xdc\x57\x43\xf7\x5e\xba\xc5\x60\xc5\xc4\xb9\x97\x56\x4c\x7c\x2e\xd3\xe4\x4d\xa1\x15\xa4\x55\x42\xbd\xc4\x5b\x4c\xd5\x86\x18\x1d\xd2\x1c\xd9\x71\xf9\xb9\x74\x0f\x2b\xfa\x10\xdc\xc4\x5a\x4c\xd3\x85\xaa\x22\xac\xb4\x84\x59\xa2\x57\x92\x17\xd5\x85\x33\xeb\xc0\xf3\x63\xeb\xc0\x2f\x9c\x0f\x4e\x28\x42\xbb\x6f\xb7\xf7\x30\x8e\x4b\xcf\x24\x39\x9a\xaf\x8f\xe5\xd8\xfa\x0d\xcd\x3f\x01\x00\x00\xff\xff\xc5\xca\x78\x60\xaa\x0a\x00\x00")

func templatesCf_dnsTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesCf_dnsTf,
		"templates/cf_dns.tf",
	)
}

func templatesCf_dnsTf() (*asset, error) {
	bytes, err := templatesCf_dnsTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/cf_dns.tf", size: 2730, mode: os.FileMode(480), modTime: time.Unix(1513818665, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesCf_lbTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x59\x4b\x8f\xdb\x36\x17\xdd\xfb\x57\x10\xc2\xb7\xf8\x02\x44\x8e\xed\x99\xa6\xee\x22\xab\xa2\xdb\xb4\x8b\xee\x82\x80\xa0\xa8\x2b\x9b\x30\x47\x64\x49\xca\x8e\x11\xcc\x7f\x2f\x48\x4a\xb6\x9e\xd4\x63\x66\x50\x4c\xb2\x18\x8d\xc8\x7b\x2e\x79\xee\x83\x47\x9c\x33\x51\x8c\x24\x1c\x50\xa4\x35\xc7\x14\x94\x61\x19\xa3\xc4\x40\x84\x7e\xae\x10\x32\x57\x09\xe8\x0b\x8a\xb4\x51\x2c\x3f\x44\xab\xe7\xd5\x6a\xd0\x02\x4b\xc5\xce\xf6\xe7\x09\xae\x83\xd6\xa2\x30\xb2\x30\x28\x52\xa2\x30\xa0\x70\x42\xe8\x09\xf2\x14\x6b\x50\x67\x46\x4b\xa7\x67\xc2\x0b\x67\xf7\xbf\x9f\x07\x21\x0e\x1c\x30\x15\x4f\xb2\x30\xd0\x9e\xbe\xf6\x28\x31\x4f\xe2\x72\x24\xae\x46\x72\xf2\x04\xcf\x7d\x1e\x79\x82\x99\x1c\xf3\x73\xe0\x22\x21\x1c\x93\x34\x55\xa0\xf5\x9a\x66\x71\xf5\x58\xfe\x6c\x42\x6b\x7d\xc4\x52\x89\x1f\xd7\x69\xe8\x35\x58\xad\x8f\xb1\xb3\xec\x07\x36\x54\xe2\x39\xeb\xae\x21\x1b\x2a\x63\x6f\xda\x0f\x7d\xd1\xb3\x21\x2f\x03\xdb\xa7\x0a\xd2\x63\x91\xcc\xc4\xf3\x46\x4d\x44\x05\x5a\x14\x8a\x02\x8a\x5a\x56\x19\x53\x70\x21\x9c\x47\x28\xaa\x1e\x63\x9a\x79\x5f\x36\xd4\xc8\xff\x73\x0e\xcf\x44\xad\x21\x3f\x63\x96\x3e\xc7\x34\x8b\x85\x84\x3c\x5a\x21\x94\x82\x84\x3c\xd5\x58\xe4\xe8\x0b\xfa\xd6\x76\x90\x83\xb9\x08\x75\x5a\x27\x09\x8f\xcb\xe7\xe8\xbb\x05\xf7\xcf\x37\xf0\x71\xb3\x2a\xf5\x56\x08\x11\xce\xc5\xc5\xad\x11\x21\xa9\x84\x11\x54\x70\x0b\x63\xa8\x8c\xfc\x4b\xa1\x8c\xf6\xd8\xdf\xa2\xfd\x26\xfa\x88\xa2\xc7\xc7\x07\xe7\xf8\xd9\x02\x78\x36\xb0\x22\xf9\x01\xb4\x9b\xb4\x59\xbb\xff\x9f\x36\xd1\x77\x3b\xc1\x10\x75\x00\x83\x0d\x39\xf8\xe1\x17\x57\xcc\xf7\x60\x18\x9a\x75\x11\xa1\xe8\x5e\x19\xb5\x58\xf4\x44\x21\x1c\xdd\x12\x36\x13\xea\x42\x54\xca\xf2\x03\x56\x05\x07\x0f\x7f\x34\x46\xc6\xf7\x91\xd8\x8f\x4c\x88\xbb\x35\xb4\x2c\x33\x59\xad\x77\x71\xa9\x57\x3c\xa3\xa1\x34\x28\xc3\x60\x5d\xfa\x46\xb0\xae\x56\xce\x93\xb2\xbe\x35\xf0\x0c\x73\x96\x9f\x1c\x9e\x0d\xbc\x0f\xab\xc5\xdb\x6f\x5e\xc6\x8f\x5e\x4c\x90\xfe\x0f\x18\xd2\x4d\x8a\xf4\x34\x8e\x6c\x5d\x04\x49\xea\xc4\xa0\x96\x3f\x95\x87\x0e\x2f\x5d\x62\xdc\x7c\x3f\xd9\x35\x0d\x4d\x15\x93\x86\xb9\xae\x11\x29\x20\x9c\x5f\x11\x41\x5c\x90\x14\x25\x84\x93\x9c\x82\x42\x49\x61\x10\x67\xda\x40\x8a\x88\x46\x24\x47\x16\x04\xdd\x40\x0a\xc5\xf1\x13\x91\x83\xdc\x94\xe3\x0d\x42\x0a\xc5\x63\xfb\xae\x4e\xc9\xc4\xdd\xeb\xf6\xf6\x75\x60\xff\xc3\x24\xe8\x7e\x16\x2a\x83\x39\x54\xe8\x7e\x2e\x5e\x4c\x08\x42\x2d\x09\x32\xd0\x04\x5b\xb3\x2c\xae\xfd\xb5\x8e\x15\xee\x7b\x1d\x6d\x14\x95\x10\x77\x42\xb1\x54\x90\xb1\x1f\x1d\x2e\x7b\xb2\xa8\xd0\xa0\x2c\x23\x67\x96\x42\x6a\xb7\x80\x4a\xe5\x84\x4e\x70\x45\x9f\xdc\x9b\x9a\x37\x24\x09\x53\xae\x20\xee\xfa\xca\xbb\xc9\x18\x87\xff\x5b\x5f\x01\x25\xf6\xc1\xad\xa0\x0e\x17\x34\xfd\xe0\x4f\x2e\xce\x32\xa0\x57\xca\xa1\x3c\xbd\xa8\x02\x8b\x97\x40\x26\x14\xe0\x14\xb4\x51\xc2\x2e\xc2\xa8\x02\xdc\x61\x15\x62\xaf\x0c\x67\x2b\x21\xcb\x80\x86\xcf\x8d\xb2\x8b\x3b\x0e\x33\x52\x70\x53\x1d\x64\x2f\x94\x87\x53\xcb\xea\x08\x84\x9b\x23\xa6\x47\xa0\x27\xbf\x7e\x59\x24\x9c\xd1\xd8\x0f\xc4\xe5\x40\x70\x0b\xde\xc2\x6d\xc2\x35\xa7\x3a\x66\x25\x0e\x84\x32\xb5\x8a\xd8\x6f\xf6\x1b\xf7\x5e\xc1\x3f\x05\x68\x83\x25\x31\x47\x8b\xfd\xc9\xdb\x46\xa3\x94\x77\x1c\x4d\x59\xfc\x60\x3b\xb0\xe7\xf7\xd0\x22\x07\x97\x38\x51\xce\xd9\x18\x87\x96\xd3\x9b\x14\x75\x83\xf7\x21\xed\xbc\xb8\xdb\x6f\x42\xda\x6e\xfb\xb0\x59\xef\xb6\x5b\xa7\xef\x76\x3b\x3b\xff\xe1\x97\xf5\xf6\x37\xff\x62\xfb\xd9\x99\xd6\x05\x1f\x7a\x45\xc9\xd7\xfd\x94\x29\x3d\x49\x21\xf8\x98\xa6\xaf\x4d\x6d\x7e\xd4\xdc\xbf\xc2\x06\x53\xa1\xa1\x25\x6f\x96\x23\x25\x75\x9f\x37\x23\xcd\xfa\xc0\x87\x73\xec\x36\xfb\xfd\x7c\x40\xec\x76\xbb\xdd\x3d\xbf\x46\x3f\x0d\x46\xa2\x16\x3e\x11\x1b\xd9\xb1\x30\x74\xb6\x08\x40\x6b\x26\x72\x4c\xb2\x8c\xe5\xcc\xb8\x73\xed\xeb\x9f\x5f\xff\x18\x89\x6b\x9f\x10\x1e\x0e\xef\xd8\x3a\x1a\xe2\x75\x5e\x82\x0f\x2a\x56\x0b\xe3\xe2\xe1\xf5\x75\x3d\x78\x7f\xff\xfe\x57\x4b\x75\xbf\xde\x2d\xc1\xf2\xa2\xad\xdd\x17\x4c\xa8\xda\x66\x65\xdd\x6d\x27\x95\x56\x6d\xfa\x7b\x28\xab\xed\x66\xf7\x18\x3f\xec\x7e\xfd\xbc\x5f\x5e\x5c\x1d\x76\xc3\xd5\xd5\x68\x8a\xbd\xec\x8e\xf1\xba\x40\x1c\x04\xa2\x38\x25\x8e\x1d\x79\xb0\x54\x1c\x74\x5a\xcb\x22\x02\x82\xcd\xc5\x4a\xb1\xda\xfe\x5d\x0c\x5d\xe0\xbb\x81\xec\x90\xd5\x1b\xce\x8f\x2b\x84\xc2\x21\xed\xed\x59\x21\xca\xc7\x19\x9f\xd9\xb5\x6a\x8b\x0e\xb6\xad\x5a\xbe\xbf\x46\xf3\x9a\x70\x11\xb9\xbc\x6b\x5d\xf4\x6c\x8d\x71\x19\xb9\xaa\xb2\x13\xe6\xe5\xe7\x24\xc4\xd9\xf9\x38\x31\x15\x7b\x24\xfd\xa4\x16\xd3\x9b\x8f\x17\x5d\xde\x0a\x4d\xca\xc6\xdb\xec\xf9\xb9\x78\xd1\xe1\x1c\x74\xb7\x3d\xaf\x90\x7c\xed\x2b\xeb\x45\x74\xcc\x62\xe3\x0d\xc8\xd8\x6f\xde\x84\x8b\xf6\xf5\xfd\xd2\x2a\x2c\x2f\xf2\xbb\x7f\x71\x69\x01\xdb\xc3\x72\x0c\xb8\xd2\x15\x37\xd4\x9a\xed\x0c\x49\xe2\x8d\x9b\x71\xeb\xc6\xcc\xcf\xba\xe9\xfb\x4a\x52\xbc\xf9\xc7\xe0\xfe\xf1\x71\xa9\x92\x68\x70\x3d\x59\x43\x74\xe8\x18\xe2\x62\x56\xe3\xeb\x23\x39\xa0\x17\x2a\x0f\xaf\x28\xfa\x87\x56\x10\x72\x3e\xab\x36\x4b\xba\xc3\xd5\x69\xc3\xf9\xf2\xfa\xec\xf9\x73\xd8\xbf\x01\x00\x00\xff\xff\x16\xc9\x3a\x1c\x9c\x1d\x00\x00")

func templatesCf_lbTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesCf_lbTf,
		"templates/cf_lb.tf",
	)
}

func templatesCf_lbTf() (*asset, error) {
	bytes, err := templatesCf_lbTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/cf_lb.tf", size: 7580, mode: os.FileMode(480), modTime: time.Unix(1513818663, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesConcourse_lbTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x93\xc1\x6e\xdb\x30\x0c\x86\xef\x7a\x0a\x82\xd8\x71\x32\x0a\xaf\x87\x5e\x76\x1a\x76\xed\x76\xd8\x6d\x28\x04\xc5\xa6\x1d\xa1\xaa\x28\x48\x72\x8c\xa1\xf0\xbb\x0f\xb2\x1d\xc7\x5d\xb3\x24\x40\x10\xa0\x27\x13\x34\xf5\x93\xfa\x7e\x8a\xbb\xe4\xbb\x04\x58\xb1\xab\xb8\x0b\x91\x54\xd2\xa1\xa5\xa4\x3c\xb3\x45\x78\x15\x00\x3b\x6d\x3b\x82\xaf\x80\x9f\x5e\x5b\xe6\xd6\x92\xaa\xf8\xc5\x77\xe9\x4d\x69\x31\xc5\x72\x8c\x9d\x7e\xa1\x01\xc5\x20\xc4\x7b\x79\xbb\x51\xc6\x9f\x13\xd6\x75\x1d\x28\xc6\x62\x39\x26\xf7\x99\xf9\x3b\xa9\x07\x8a\xdc\x85\x8a\x00\xff\x39\xdf\x98\x40\xbd\xb6\x16\x01\xf7\xa1\x5c\xb4\xa6\xe6\x79\x46\x00\x98\xda\xef\x74\x28\xc8\xed\x94\xa9\x87\x43\x9d\x64\x4f\x0e\x73\x29\xa5\x9e\xc3\xf3\xd1\x49\xe7\x7f\xc5\x66\x63\xe5\x3e\x9e\xaf\x2f\x00\xb4\xb5\xdc\x8f\xed\x00\x7c\xe0\xc4\x15\xdb\x2c\x93\x2a\x8f\x53\x92\x43\x8a\xd3\x18\xbf\xf1\xe1\x0e\x3f\x03\xde\xdf\x7f\xc9\x9f\xb2\x2c\x4b\x7c\x12\x00\x43\x16\x9a\x49\x27\xdd\xc6\xb1\xf4\x70\x99\xa7\x93\x20\x66\x5c\xb8\x72\x40\x2e\xb9\x05\xc3\xff\x19\x9c\xc6\xfc\x66\x55\x70\xb5\x01\x17\x6a\x0b\x80\x48\x31\x1a\x76\x4a\x37\x8d\x71\x26\xfd\xc9\xf5\x8f\x3f\x1e\xbf\x9f\xf1\x97\x43\xaf\x43\x6d\x5c\xab\x42\x67\x09\x01\x63\xdc\xca\x43\x56\x4e\xd9\xb5\xcf\x67\xbc\x8e\x71\x8b\x0b\xe7\x55\xf5\x85\x1b\x1f\xc9\x36\xca\x1a\xf7\x3c\x64\x95\xec\xaa\x0a\xda\xb5\x34\xaa\x8c\x56\x0a\x00\xe3\xd5\x7a\x09\x7e\x7d\xfb\x39\x67\x67\x47\x8e\xb7\xbc\xfa\x2d\xbc\x63\xb5\x4d\xc9\xc7\xab\x68\x8d\x0a\x37\xe3\x95\x5f\xc0\x07\xc3\x75\x35\xad\x9b\xc1\x7a\xb8\xbb\x35\xab\xbf\x01\x00\x00\xff\xff\x5f\xcd\x56\x41\x23\x06\x00\x00")

func templatesConcourse_lbTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesConcourse_lbTf,
		"templates/concourse_lb.tf",
	)
}

func templatesConcourse_lbTf() (*asset, error) {
	bytes, err := templatesConcourse_lbTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/concourse_lb.tf", size: 1571, mode: os.FileMode(480), modTime: time.Unix(1513818659, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesJumpboxTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x8f\x41\x0a\x83\x30\x10\x45\xf7\x39\x45\x08\xdd\xaa\x20\x64\x23\xf4\x2c\x21\x35\x83\xb5\x44\x27\x4c\x66\x44\x10\xef\x5e\x4a\x6d\xad\xd0\x4d\xbb\x1d\xe6\xbd\xff\x3f\x41\x46\xa1\x16\xb4\xe9\x10\xbb\x08\xae\xc5\x21\x09\x83\xf3\x21\x10\xe4\x6c\xb4\xb9\xc9\x90\x2e\x38\x17\x7d\x32\x7a\x51\x5a\x8f\x7e\x00\x7d\xd6\xe6\xb4\x4c\x9e\x4a\x18\x27\xd7\x87\xb5\xf8\xf8\x52\xab\x52\x28\x9c\x84\xdf\xb0\x13\x8a\x4f\x7a\xf2\x51\x36\xfc\x7b\x62\xb9\x9b\xca\xed\xb4\x36\x75\x7d\xb0\xc2\xcc\x40\xa3\x8f\xee\xd5\xe9\x2f\xeb\x41\x19\x7a\x82\x96\x91\xf6\xe1\x07\xef\x95\x39\xe5\xa6\xaa\x7e\x6b\x6d\xad\xb5\x8f\x94\x7b\x00\x00\x00\xff\xff\x10\xf1\x6b\x1b\x66\x01\x00\x00")

func templatesJumpboxTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesJumpboxTf,
		"templates/jumpbox.tf",
	)
}

func templatesJumpboxTf() (*asset, error) {
	bytes, err := templatesJumpboxTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/jumpbox.tf", size: 358, mode: os.FileMode(480), modTime: time.Unix(1513818657, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesVarsTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xd0\x41\xca\xc3\x20\x10\x05\xe0\xbd\xa7\x78\x0c\xff\xfa\xbf\x41\xcf\x52\x6c\x32\xc8\x14\x71\x64\x22\x42\x1b\xbc\x7b\xb1\x09\x4d\xba\x28\xe8\xf2\xf1\xc9\x63\x5e\xf5\x26\xfe\x16\x19\x94\x4d\xef\x3c\x95\xab\xcc\x84\xd5\x01\xe5\x91\x19\x17\xd0\x52\x4c\x52\x20\xd7\x9c\x3b\xb0\x71\x10\x4d\x03\xf0\xa9\x89\x07\x18\xa7\x3a\x56\x3c\x19\xcf\x9c\x8a\xf8\xb8\xfc\xd4\xd9\xb4\xca\xcc\x06\x0a\xaa\x21\xee\xfd\xa7\x9f\xdd\xff\xad\xd5\xdb\xff\x29\x6c\xe4\x80\x7d\x04\xf4\xf7\x41\xc7\x32\x6f\xb3\xdd\x8e\x6f\xb3\x85\xad\xf7\xbf\x02\x00\x00\xff\xff\xb3\x6b\x5f\x21\x53\x01\x00\x00")

func templatesVarsTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesVarsTf,
		"templates/vars.tf",
	)
}

func templatesVarsTf() (*asset, error) {
	bytes, err := templatesVarsTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/vars.tf", size: 339, mode: os.FileMode(480), modTime: time.Unix(1513818756, 0)}
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
	"templates/bosh_director.tf": templatesBosh_directorTf,
	"templates/cf_dns.tf": templatesCf_dnsTf,
	"templates/cf_lb.tf": templatesCf_lbTf,
	"templates/concourse_lb.tf": templatesConcourse_lbTf,
	"templates/jumpbox.tf": templatesJumpboxTf,
	"templates/vars.tf": templatesVarsTf,
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
	"templates": &bintree{nil, map[string]*bintree{
		"bosh_director.tf": &bintree{templatesBosh_directorTf, map[string]*bintree{}},
		"cf_dns.tf": &bintree{templatesCf_dnsTf, map[string]*bintree{}},
		"cf_lb.tf": &bintree{templatesCf_lbTf, map[string]*bintree{}},
		"concourse_lb.tf": &bintree{templatesConcourse_lbTf, map[string]*bintree{}},
		"jumpbox.tf": &bintree{templatesJumpboxTf, map[string]*bintree{}},
		"vars.tf": &bintree{templatesVarsTf, map[string]*bintree{}},
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

