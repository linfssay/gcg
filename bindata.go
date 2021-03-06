// Code generated by go-bindata.
// sources:
// template/controller/controller.go.tmpl
// template/controller/gen_controller.go.tmpl
// template/controller/gen_es_controller.go.tmpl
// template/main.go.tmpl
// template/model/database/db.go.tmpl
// template/model/database/es.go.tmpl
// template/model/database/esquery.go.tmpl
// template/model/database/table.go.tmpl
// template/model/model.go.tmpl
// template/router/duplicateQuest.go.tmpl
// template/router/gen_router.go.tmpl
// template/router/router.go.tmpl
// DO NOT EDIT!

package main

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

var _controllerControllerGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x90\xcd\x4a\x3b\x31\x14\xc5\xd7\x13\xc8\x3b\x5c\x66\xd5\xc2\x40\xfa\xff\x2f\x85\x59\xa9\xcb\x6e\xf4\x09\x62\x7a\x1b\x8a\xd3\x4c\x49\x32\x48\x29\x05\x07\x2c\x48\x57\x2e\xaa\xa0\x82\x1f\xa0\x3b\x61\x28\x8a\xe0\xaa\x2f\xe3\x7c\xf8\x16\x12\x67\x2a\xad\xab\x5c\xce\xcd\x39\x27\xbf\x30\x06\x12\x15\x6a\x6e\x11\x8e\xc6\x20\x85\xa4\x64\xc4\xc5\x31\x97\x08\x22\x56\x56\xc7\x51\x84\x9a\x12\x4a\x18\xcb\x67\x2f\xf9\x6c\x59\x5c\x7d\x14\xaf\x97\xe5\xcd\x59\xb5\x5a\xe4\xb7\x77\x9f\xab\xc7\x22\xcd\x28\xb1\xe3\x11\xc2\x01\x9a\x24\xb2\x60\xac\x4e\x84\x85\x09\x25\xde\x6e\xdc\x43\x18\x28\x4b\x89\xd7\x35\xd2\x6d\x06\x4a\x52\xe2\xed\x71\xcb\x9d\x8e\xba\xcf\x05\x4e\xa6\x94\x4c\xeb\x92\xaf\xc5\x75\x95\x65\xe5\xfc\xbd\x38\x4d\xcb\x87\x94\x12\x11\x2b\x63\xe1\x30\x11\x02\x8d\x01\x08\xe1\x7f\xa7\xb3\x56\xf7\xb5\x8e\x75\xad\xfd\xdb\xb4\xaf\x1f\xb5\xe5\x75\xfd\x21\xf8\xc5\xf9\x45\x3e\xbf\xf7\xb7\x22\x9a\x55\xfe\xb4\xac\xde\x9e\x7d\x97\xd4\x4f\x94\x68\x70\xba\x46\xb6\x44\x83\x01\xc1\xf0\x97\x02\x82\xde\x1f\x88\x76\xe3\x70\xe0\x8c\x69\x34\xb0\x13\x82\xc2\x93\x56\x2d\xb7\x29\xf1\x34\xda\x44\xab\x8d\x7b\x9e\x8b\x0e\xdc\x30\x34\xf2\xe7\x74\xa9\x6e\x70\x7f\xf2\x1d\x00\x00\xff\xff\x43\x87\x6a\x86\x9f\x01\x00\x00")

func controllerControllerGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_controllerControllerGoTmpl,
		"controller/controller.go.tmpl",
	)
}

func controllerControllerGoTmpl() (*asset, error) {
	bytes, err := controllerControllerGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "controller/controller.go.tmpl", size: 415, mode: os.FileMode(438), modTime: time.Unix(1562726307, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _controllerGen_controllerGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x93\x4f\x6f\xd3\x40\x10\xc5\xcf\x44\xca\x77\x98\xee\x01\xd9\x92\xb5\x8e\x10\xa7\x4a\x3e\xd0\x16\x90\x40\xb4\x85\x20\x71\x5e\xaf\x27\x1b\x27\xeb\x5d\x6b\x3c\xae\xa8\xac\x7c\x77\xb4\xb1\x0b\x2d\xca\x1f\xa3\x14\x89\x03\x47\xef\x78\xdf\xdb\xf9\xcd\xbc\x34\x05\x83\x0e\x49\x31\x42\x7e\x0f\x46\x9b\xe9\xa4\x56\x7a\xad\x0c\x82\xf6\x8e\xc9\x5b\x8b\x34\x9d\x4c\x27\x65\x55\x7b\x62\x88\xa6\x13\x00\xe1\x90\xd3\x25\x73\x2d\xb6\x5f\xa6\xe4\x65\x9b\x4b\xed\xab\xd4\xaa\xbc\x61\xa5\xd7\x29\xea\xa5\xef\xab\x5d\x27\x6f\xc9\xaf\x50\xf3\xb5\xaa\x70\xb3\x49\x2b\x5f\xa0\x4d\xbb\xae\x75\x05\x92\x2d\x1d\x82\xbc\xba\xe8\x6b\x62\x3a\x89\x83\xd9\xa2\x75\x1a\xba\xae\xad\x6b\xa4\x4b\x55\xa1\x05\xf9\x55\xe5\x16\xfb\xbf\xe6\xea\x0e\x23\x0d\xc1\x42\x5e\x7a\xc7\xf8\x9d\x63\x40\x22\x4f\xd0\x05\x4b\x80\x3b\x45\xe0\xf3\x15\xec\x34\x91\x7b\x85\xfb\xcb\x48\x04\xe7\x99\x96\x17\xa5\x2b\xa2\x97\x3e\x5f\xc5\xfd\x79\xb9\x08\xa5\xb3\xcc\x95\x76\xb0\x79\x41\xc8\x2d\x39\xd0\xf2\xc3\xfc\xe6\x3a\x0a\x44\xe4\x9c\x15\xb7\xcd\xcd\xc7\x04\xbe\x60\xd3\x5a\xfe\xd4\x98\xe8\xf5\x6c\x96\x84\xbb\xf2\x6d\x78\x64\x14\x27\x20\x44\x3c\xa8\x0e\xa6\x87\x94\x7e\x09\xbd\x9a\xcd\x12\xe1\xd7\x22\xf1\xf9\x4a\x6e\x31\xc4\x41\x67\x73\x9c\xd9\x15\x5a\xe4\x7f\x83\xda\xb1\x76\xff\x3a\xb8\x81\xc5\x6f\xe8\xde\x23\xef\x6d\xf1\xdb\x12\xe9\x08\x3c\xed\x5d\x01\xe7\x19\x68\xf9\xb9\x45\xba\xbf\x55\xa4\xaa\x48\x84\x53\x31\x3c\x58\x91\xd9\x12\x7a\x5c\x57\x64\x1e\xca\x81\x3e\x61\x03\xa5\x63\xa4\x85\xd2\xd8\x6d\x7e\x12\x54\x64\xce\x32\x21\x9e\x00\x6c\x20\xdb\x3d\xa8\x11\x8d\x78\x57\x24\x8a\xcc\x03\x49\xb4\x0d\xc2\x73\x8a\x9f\x34\x22\xc2\x66\xfc\x64\xde\x58\xfb\x7f\xa9\xff\x7c\x54\x01\x5b\x68\x67\x3c\xe8\x77\x25\x35\x7c\x6a\x04\x46\xc4\xe0\x60\x14\xf6\xc6\xe1\x84\xad\x1d\x1a\x7b\x1a\x89\x9d\xb1\x78\x16\x93\xc7\x06\xe3\x56\xe7\x60\x44\x7e\x04\x00\x00\xff\xff\x18\x37\x23\xe3\xc1\x07\x00\x00")

func controllerGen_controllerGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_controllerGen_controllerGoTmpl,
		"controller/gen_controller.go.tmpl",
	)
}

func controllerGen_controllerGoTmpl() (*asset, error) {
	bytes, err := controllerGen_controllerGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "controller/gen_controller.go.tmpl", size: 1985, mode: os.FileMode(438), modTime: time.Unix(1595298856, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _controllerGen_es_controllerGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x92\xcd\x8a\xdb\x30\x10\xc7\xcf\x35\xf8\x1d\xa6\x3a\x14\x1b\x8c\xbc\xf4\xb8\x90\xcb\x66\x43\xa1\xa5\xbb\xdb\xba\x7d\x00\x59\x9e\x75\xe4\xc8\x92\x19\x8f\x4a\x17\x93\x77\x2f\x8a\xdd\x42\x4b\x92\x86\x1e\x4a\x8e\xf6\x68\xfe\x1f\x3f\xa9\x2c\xa1\x45\x87\xa4\x18\xa1\x7e\x81\x56\xb7\x69\x32\x28\xbd\x53\x2d\x82\xf6\x8e\xc9\x5b\x8b\x94\x26\x69\x62\xfa\xc1\x13\x43\x96\x26\x00\xc2\x21\x97\x5b\xe6\x41\x1c\xbe\x5a\xc3\xdb\x50\x4b\xed\xfb\xd2\xaa\x7a\x64\xa5\x77\x25\xea\xad\x9f\xa7\xd3\x24\x9f\xc8\x77\xa8\xf9\x41\xf5\xb8\xdf\x97\xbd\x6f\xd0\x96\xd3\x14\x5c\x83\x64\x8d\x43\x90\xf7\x77\xf3\x4c\xa4\x49\x1e\xcd\x9e\x83\xd3\x30\x4d\x61\x18\x90\xd6\xaa\x47\x0b\xf2\x8b\xaa\x2d\xce\xa7\x36\xd5\x9a\x50\x31\x66\x1a\xa2\x8d\x5c\x7b\xc7\xf8\x9d\x73\x40\x22\x4f\x30\x45\x5b\x80\x6f\x8a\xc0\xd7\x1d\x1c\x35\x92\x27\xc5\xe7\x65\x24\x82\xdb\x95\x96\x77\xc6\x35\xd9\x1b\x5f\x77\xf9\xfc\xdf\x3c\xc7\xd1\xeb\x95\x33\x76\xb1\x79\x45\xc8\x81\x1c\x68\xf9\xbe\x7a\x7c\xc8\x22\x15\x59\xb1\xe2\x30\x3e\x7e\x28\xe0\x33\x8e\xc1\xf2\xc7\xb1\xcd\xde\xde\xdc\x14\x71\x57\x6e\x62\xc8\x2c\x2f\x40\x88\x7c\x51\x5d\x4c\xcf\x29\xfd\x2e\x24\xfc\x4e\x14\xbe\xee\xe4\x2f\x14\x79\xd4\xda\x5f\xc2\xee\x1e\x2d\x5e\x09\xbb\xbf\x95\xfe\x0f\xf8\x16\x1a\x97\xe3\xfb\x3a\x34\xd7\xf2\xf4\xae\x00\xdf\x42\xe3\x0f\x7c\x9b\xea\x1d\xf2\x78\xb2\xe7\x79\x78\xa6\x81\xdb\x15\x68\xf9\x29\x20\xbd\x3c\x29\x52\x7d\x26\x4c\x23\xf2\x7f\x4a\x79\xfc\x06\x0e\xf9\x4e\xc7\x33\xcd\xcf\x3e\x3f\x02\x00\x00\xff\xff\x09\x8b\x07\x59\x1f\x05\x00\x00")

func controllerGen_es_controllerGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_controllerGen_es_controllerGoTmpl,
		"controller/gen_es_controller.go.tmpl",
	)
}

func controllerGen_es_controllerGoTmpl() (*asset, error) {
	bytes, err := controllerGen_es_controllerGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "controller/gen_es_controller.go.tmpl", size: 1311, mode: os.FileMode(438), modTime: time.Unix(1562726307, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mainGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\xc8\x4d\xcc\xcc\xe3\xe5\xe2\xe5\xca\xcc\x2d\xc8\x2f\x2a\x51\xd0\xe0\xe5\x52\x50\x50\xaa\xae\xd6\x0b\x28\xca\xcf\x4a\x4d\x2e\xf1\x4b\xcc\x4d\xad\xad\xd5\x2f\xca\x2f\x2d\x49\x2d\x52\xe2\xe5\xd2\x04\xa9\x4d\x2b\xcd\x4b\x06\x6b\xd4\xd0\xac\xe6\xe5\xe2\x84\x48\xea\x05\x81\x29\x0d\x4d\x5e\xae\x5a\x5e\x2e\x40\x00\x00\x00\xff\xff\xf4\x0b\xd4\xcc\x5f\x00\x00\x00")

func mainGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_mainGoTmpl,
		"main.go.tmpl",
	)
}

func mainGoTmpl() (*asset, error) {
	bytes, err := mainGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "main.go.tmpl", size: 95, mode: os.FileMode(438), modTime: time.Unix(1562726307, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _modelDatabaseDbGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x90\x5f\x6b\x83\x30\x14\xc5\x9f\x0d\xe4\x3b\xdc\xc9\x28\x3a\x5a\xed\x5e\x07\xbe\xb4\x85\x51\xd8\x3f\x68\xdf\x47\x8c\x57\x17\x16\x13\x1b\x63\x69\x11\xbf\xfb\x48\x94\xd2\x16\xe6\x93\xf7\x9e\x93\xdf\xb9\xf7\xa6\x29\x54\xa8\xd0\x30\x8b\x05\xe4\x67\xa8\x78\x45\x49\xc3\xf8\x2f\xab\x10\xfa\xbe\x53\x05\x1a\x29\x14\x42\xf2\xc1\x6a\x1c\x06\x4a\xfa\x1e\x1e\x2d\xcb\x25\xc2\x4b\x06\xc9\xde\xff\xb9\x3e\x25\xa2\x6e\xb4\xb1\x10\x51\x12\x84\x95\xb0\x3f\x5d\x9e\x70\x5d\xa7\x55\xa7\x99\x2a\x59\x2d\xe4\x39\x6d\x0f\xf2\x14\x52\x12\x7c\xc3\x8d\x43\x2f\xda\x83\x5c\x14\x46\x1c\xd1\xa4\xf5\xb9\x3d\xc8\x90\x92\xd8\x31\x8f\xcc\x78\x60\x91\x6f\x55\xa9\x61\xfc\x9e\x1c\x26\xd9\xac\x26\x4f\xd9\x29\x0e\xaf\x68\x37\xab\x28\xbe\x68\xd0\x53\x12\x88\x12\xa6\x87\x0f\x19\x28\x21\x7d\x33\x30\x68\x3b\xa3\x26\x85\x92\x60\xa0\xc4\x51\xc7\x7a\x0e\x68\x8c\x5b\xcd\x73\xd6\x5a\x29\xe4\x36\x0a\xc7\xa1\xe6\x10\xf6\x7d\xb2\xd3\x9d\xe1\x38\x0c\xb3\x86\x99\x16\xf7\xa2\xc6\x6c\x6f\x3a\x9c\x49\xcd\xb3\x37\xcd\x99\x0c\xe3\x31\xdb\x91\xae\x83\x1b\xa6\x04\x8f\xd0\x98\x78\x4c\x9d\xb6\x4a\x76\x68\xdf\xd9\x69\x5b\x48\x74\x79\x6d\xb4\x8c\xef\xb5\xcf\x06\xd5\xa8\x3d\x2f\xaf\xd5\x2f\xa1\xaa\xc8\xd5\x77\x3b\x0d\x97\xbb\xac\xa5\x6e\xd1\x5f\xe6\xff\x83\x4c\xb0\xcd\x2a\xf1\xee\xe8\x76\x3c\xf0\x4e\xcf\xfc\x0b\x00\x00\xff\xff\xb2\x70\xa4\x74\x2e\x02\x00\x00")

func modelDatabaseDbGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_modelDatabaseDbGoTmpl,
		"model/database/db.go.tmpl",
	)
}

func modelDatabaseDbGoTmpl() (*asset, error) {
	bytes, err := modelDatabaseDbGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "model/database/db.go.tmpl", size: 558, mode: os.FileMode(438), modTime: time.Unix(1562726307, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _modelDatabaseEsGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x91\x4f\x8b\xdb\x30\x10\xc5\xcf\x16\xe8\x3b\x4c\x0d\x0b\x4a\x31\xca\x7d\xc1\x87\xfe\x49\x4f\x21\x84\x86\xf6\xd4\x8b\x22\x8f\x15\xb1\xb6\x26\x8c\xe4\xa4\x25\xe4\xbb\x17\xc9\xf1\xd2\xee\xa9\xd0\x83\x91\x3d\x6f\xe6\xbd\xdf\x58\xeb\x35\x38\x0c\xc8\x26\x61\x07\xc7\x5f\xe0\xac\x93\xe2\x6c\xec\x8b\x71\x08\x11\xf9\xe2\x2d\x4a\x21\x85\x1f\xcf\xc4\x09\x94\x14\x55\x6d\x29\x24\xfc\x99\xea\xfc\xde\x8f\xf3\xe9\x7c\x3a\x4d\x47\x6d\x69\x5c\xd3\xe0\x2f\xc8\xb8\xc6\xc1\xc4\xe4\x6d\x91\x07\x72\xe5\xa4\x58\x4b\xb1\xca\x86\x17\xc3\xc5\xcd\x0e\x1e\x43\x82\xf7\x8f\x6e\xfd\xa9\x7c\x4b\x51\x9d\x28\x26\x68\xa1\xbe\xdd\xf4\xe6\x70\xa0\x89\x2d\xde\xef\xcb\x74\x3f\x05\x0b\x3e\xf8\xa4\x56\x00\x37\x29\x2a\x64\x26\x1e\xc8\xc1\x73\x0b\x03\x39\xbd\xc3\xab\xa2\xa8\x0f\xa9\xa3\x29\x35\x50\x7f\xd8\xef\xeb\xa6\x28\xdb\x98\xba\x2f\x83\x71\x71\x25\x45\x95\x29\x90\xcb\x43\xfc\x4a\xd3\x94\x5a\x0b\x0b\xd3\x0e\xaf\x33\x96\x5a\x2a\x07\x4c\x9b\x3c\xb2\x25\xa7\x96\xe8\x55\x03\x7f\xc8\xdf\xbe\x6e\x55\x5e\x61\x95\x63\x7c\x5f\x1c\xdf\xb5\x10\xfc\x50\x78\xab\xb3\x09\xde\xe6\xd9\xac\xdf\x73\x4f\xe8\xa9\x01\x4b\x1d\xce\xf1\xcf\x2d\xcc\x34\x7a\xef\x83\x9b\xbd\xf4\x67\x52\x8f\xbf\xaf\x3f\x1a\xfb\xe2\x98\xa6\xd0\xa9\x7f\xce\xe8\xc7\xa4\xf7\xec\x43\xea\x55\xbd\x99\x59\x23\x1a\xb6\x27\x60\x4c\x13\x07\xec\xe0\xea\xd3\xa9\x50\xc0\x53\x07\x26\x74\x70\x41\x8e\x9e\x02\x3c\xc5\x1f\xa1\x5e\x00\x33\xac\xfe\x3e\x2b\x7a\x37\x8d\x47\xe4\x72\x2f\x15\xc6\x47\xff\xdb\x25\xfe\x8a\x7b\x4c\xce\x4b\xfd\x37\xfb\x1b\xc2\x57\x84\x42\x74\x97\xe2\x77\x00\x00\x00\xff\xff\xcb\xa4\xfa\xc5\xe4\x02\x00\x00")

func modelDatabaseEsGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_modelDatabaseEsGoTmpl,
		"model/database/es.go.tmpl",
	)
}

func modelDatabaseEsGoTmpl() (*asset, error) {
	bytes, err := modelDatabaseEsGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "model/database/es.go.tmpl", size: 740, mode: os.FileMode(438), modTime: time.Unix(1562726307, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _modelDatabaseEsqueryGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x95\x4d\x6e\xdb\x3a\x10\xc7\xd7\x16\xa0\x3b\xcc\x13\x10\x40\x0a\xf2\x28\xbc\xed\x03\xbc\x49\xec\x1a\x6e\x81\xa0\xa8\xd3\xae\x0a\x14\xb2\x38\x32\x98\x48\xa4\x3a\xa2\xdc\x18\x84\x81\xee\x8a\x5e\xa1\x9b\xae\x7a\x8d\xf6\x36\x41\x7b\x8b\x82\xa2\xe4\x8f\xd8\xf9\x70\x53\xae\xa8\x21\x67\x86\x33\xbf\x3f\xa9\x38\x86\x19\x4a\xa4\x44\x23\x87\xe9\x02\x66\xe9\xcc\xf7\xca\x24\xbd\x4a\x66\x08\xc6\xd4\x92\x23\xe5\x42\x22\xb0\xc1\xe9\x79\x52\xe0\x72\xe9\x7b\xbe\x27\x8a\x52\x91\x86\xd0\xf7\x00\x02\x2d\x0a\x0c\x9a\x59\x56\xe8\x66\x62\x8c\xc8\x00\xdf\x03\x7b\x49\xa2\x48\x68\xf1\x02\x17\x17\x8b\x12\x21\x10\x52\x07\x36\x02\x40\x50\x69\x4a\x95\x9c\x07\x60\xcc\xbf\x80\x92\xb7\xe6\x54\x49\x8d\xd7\x36\x4c\x64\x13\xcd\x13\x82\x77\xd0\x07\x9b\x83\x9d\xab\x0f\xd6\x16\x1f\xfb\xde\x70\x02\x67\x84\x89\x46\xdf\x3b\x8e\x7d\x2f\xab\x65\x0a\xa1\x9a\x5e\xc2\xb1\x31\x75\x59\x22\x9d\x25\x05\xe6\xc0\x2e\x92\x69\x8e\xee\xdc\x11\x0c\x27\xce\x27\x8c\x40\x48\x8d\x94\x25\x29\x9a\xa5\xb1\x31\x01\xca\x5a\x9f\x00\x12\xc1\xff\x7d\x48\x73\x81\x52\xb3\xb1\xe4\x78\x1d\x46\xcc\x2e\xf7\x7a\xee\x2b\x30\x66\x33\x68\xe0\x56\xbb\xf1\x88\xc2\x57\x63\xcc\xc3\xb6\x07\x6c\xac\x55\x62\x4f\xcf\xb6\x0f\xbf\x0e\xb2\x5c\x46\xdb\x89\x5c\xd7\xf2\x0a\x77\x83\x3e\x10\x67\x4f\x98\xae\xf9\xbd\xde\xa9\xe2\x8b\xe7\x95\x92\x36\x48\x57\xf7\x40\x85\x2d\x14\x76\x9a\xa4\x57\x33\x52\xb5\xe4\x61\x14\x35\xab\xb6\x5c\x22\xf8\xa7\x0f\x52\xe4\x60\x9c\x47\x56\x68\x9b\x53\xea\x5c\x86\x48\xe4\x76\x36\x19\x08\x75\x4d\xd2\xf6\x9a\x8d\xb9\xef\x2d\x37\x70\x72\xcc\xf1\x50\x9c\x83\xc6\x27\x8c\x60\xaa\x54\x6e\x5c\x61\x84\xd5\x6d\x8e\xdd\x36\xf6\x77\x11\x3e\x05\xe0\x5e\x7c\x87\xc1\xdb\x42\xb7\x1e\xf7\xe1\x02\xd8\xc3\xcb\x8d\x72\xcd\x8b\x0d\x89\x14\xad\x5c\x5c\x53\x1b\x6e\x59\x92\x57\xe8\xac\x6d\xd6\x15\xea\x2c\x0c\x1c\x41\x0b\xa0\xce\x35\x1c\x55\x6f\x65\x70\x62\xbf\xd8\xab\xc6\x12\x75\x7c\x9a\x50\x9a\x6a\x6c\x05\xb0\x92\xc0\xeb\x92\x1f\x7c\xa3\x9d\x4f\x2b\x01\x57\xd0\x1e\x05\x74\xbb\x56\xfd\x7b\x40\x09\x8f\x16\xc1\x13\x34\xb0\x4f\x02\x07\x29\x60\x47\x00\x03\x95\xae\x2f\xee\x43\x5a\xb8\x43\x09\x77\xe9\xa0\x49\xb2\x49\xbb\x6e\x5a\x0a\xf6\x37\xb1\x1f\xf5\x2e\xe8\x16\xf3\x08\x75\xb5\x01\x79\x38\x19\xa1\xbe\x93\x71\x28\x38\x54\x9a\x84\x9c\x6d\x3d\xda\xee\xc0\xf6\xef\x60\x25\xb2\xb1\x60\xcd\x71\xfc\xeb\xe3\x97\x9f\x3f\x3e\x09\x7e\xf3\xf5\xdb\xcd\xe7\xef\xd6\x36\x43\xfd\xdf\x6d\x55\x8c\x50\xdf\xf3\x28\x8c\x79\x28\x78\xc4\xfe\xa0\x87\xfb\xde\xbf\x65\xbb\xdf\x9e\x83\x3d\xb3\x31\xba\xfd\x71\x5c\xa1\x86\x79\x92\xd7\xed\xd5\xb2\x15\xf5\x9b\x7d\x13\x55\x53\x8a\xbb\x37\x6d\xa4\x34\x70\x95\xd6\x05\x4a\x7b\xd1\x40\x48\x98\x23\x55\x42\x49\x38\xe2\x90\x91\x2a\x40\xd8\xaa\xe0\xa8\x3a\x01\x6d\x85\xdb\x22\x6a\xa2\x8e\x79\x3b\x79\xe3\x7c\x3a\xb3\xf5\x68\xe7\x56\xec\xeb\x63\xb7\x24\xd5\xf4\xb2\x01\xf9\x3b\x00\x00\xff\xff\x95\x49\xcf\x8c\x33\x08\x00\x00")

func modelDatabaseEsqueryGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_modelDatabaseEsqueryGoTmpl,
		"model/database/esquery.go.tmpl",
	)
}

func modelDatabaseEsqueryGoTmpl() (*asset, error) {
	bytes, err := modelDatabaseEsqueryGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "model/database/esquery.go.tmpl", size: 2099, mode: os.FileMode(438), modTime: time.Unix(1565687364, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _modelDatabaseTableGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x54\x5d\x6f\xdb\x36\x14\x7d\x9e\x00\xfd\x87\x5b\x22\x6d\xa5\xd8\x60\x96\xd7\x02\x42\xd1\x26\x4e\x66\xcc\x71\x36\xdb\x43\x1f\x86\xa1\xa1\xa5\x6b\x8d\x29\x45\xb9\x24\xe5\xd4\x20\xf4\xdf\x07\x52\xd6\xac\xd4\xb1\xe3\x6e\x03\x86\xbe\xd8\x12\x75\xbf\xce\xb9\x87\xe7\xec\x0c\x72\x94\xa8\x98\xc1\x0c\xe6\x6b\xc8\xd3\x3c\x0c\x96\x2c\xfd\xc4\x72\x04\x6b\x2b\x99\xa1\x12\x5c\x22\xd0\xcb\xf7\x63\x56\x60\x5d\x87\x41\x18\xf0\x62\x59\x2a\x03\x51\x18\x00\x10\xc3\x0b\x24\xfe\x69\x51\x18\x12\x06\xb1\x8b\x58\x31\x05\x1f\x21\x01\xf7\x91\x8e\xcb\x07\x77\x76\x76\x1a\x06\x9a\xad\x30\x0c\x4e\xcf\xc2\x60\x51\xc9\x14\xa2\x72\x7e\x0f\xa7\xd6\x56\xcb\x25\xaa\x0b\x56\xa0\x00\x3a\x63\x73\x81\x4d\xaf\x18\xa6\x6c\x85\x51\x0c\x5c\x1a\xeb\x5a\xb8\xb2\xfa\xb3\x80\x04\x08\x69\xdf\x79\xe6\x3e\xbb\xb7\x8c\x19\x36\x67\x1a\xe1\x4d\x02\xd7\x68\x2e\xdf\x47\xb1\x3b\xe6\x0b\xdf\x87\x3e\x6e\xf3\x8b\xe2\x05\x53\xeb\x9f\x71\x5d\xd7\x90\x24\x60\x6d\xe7\xe8\x12\x17\xac\x12\xc6\x4d\xe0\xfb\x42\xd3\x95\x0c\xc7\xd3\xc1\x64\x06\xc3\xf1\xec\x16\xee\xac\xed\xce\x7a\x07\x91\xb5\x39\x9a\xa1\xd4\xa8\xcc\x45\x29\xaa\x42\x02\x6d\xfe\xb5\x2b\xb4\x62\xa2\x42\xbd\x89\xf2\x89\x57\x1c\x45\x76\x51\x56\xd2\xe8\x6e\x24\x69\x3a\x2a\xd4\x0e\x48\x0b\x8a\xde\x54\xda\x0c\xbe\x60\x1a\xe9\xcf\xa2\xef\x8b\xdc\xce\xef\x77\xfa\x40\xdc\x66\x1b\x9e\xf5\xe1\xa3\x2b\xa1\x50\xd3\x11\xd3\x9b\xd1\x86\x59\xb4\x89\xe1\x19\x24\x8e\xbb\xc8\x07\xfb\xc3\x1a\x85\xc6\x0d\xe4\x16\x74\xb5\xcc\x98\xc1\x5d\xbc\x1a\x0d\xf8\x39\x7e\xf3\x01\xbb\xa3\x3c\xfc\x89\x0a\x81\x67\xc9\x5b\xd2\x56\xfc\x67\xa0\xfa\xcf\xac\x2f\xde\x96\x7f\x0e\xf5\x0f\x4f\xc3\x76\x3f\x0a\x4d\xa5\x24\xf0\x2c\x0c\xea\x56\xaf\x19\x0a\x34\xdf\xa4\xd8\x4b\x9f\x11\xc5\x30\x2f\x4b\xb1\x61\x72\x8f\x30\x1b\x82\xdf\x24\x70\x90\xe2\x66\x04\x37\xf3\x39\xec\x52\xfa\x34\x93\x70\x98\xb0\xad\x46\x1c\x60\xa3\x2a\xf4\x90\x1d\xe0\x6b\x34\x30\xe2\xda\x38\x33\xf8\xe0\x9a\x75\xb0\x5f\xa3\xd9\x0b\xdc\xc7\x46\x69\x29\x33\xd0\x46\x71\x99\xf7\x81\xa9\x5c\x03\xa5\x94\x4b\x83\x6a\xc1\x52\xb4\x75\x0c\xbf\xff\xb1\x9f\xbc\xe6\xae\x95\xf3\x7b\xaf\x91\x43\x91\xb6\x3e\x70\xdf\x1b\x52\xc9\x74\x30\x1a\x5c\xcc\xe0\xab\xeb\xe6\xf2\xbb\xb7\x0d\xae\x26\xb7\x37\xbb\xc4\x7f\xf8\x69\x30\x19\x74\xa8\xff\x11\xc8\xc6\x4a\x3c\xc4\x17\xce\x80\xb6\xce\xd0\x4b\x08\x30\x99\x01\x81\x5e\xf3\xbd\x47\x92\xb7\x4d\x86\x1f\x14\x95\x7a\xa4\xfa\x29\x0a\x4c\x4d\xf4\xca\x41\xed\x83\x5f\x98\x23\x8b\x52\xda\x1a\x96\xcb\x78\x91\x80\xe4\xa2\x35\xa0\x45\x61\xdc\x0e\xa5\x11\x32\x42\xa5\x76\x64\xeb\x6a\x75\x84\xdb\xdd\xe3\x3b\x21\xfc\x58\xdc\xf0\x52\x1e\xb9\xcf\x77\x42\x78\xa9\x1f\x50\xfa\x77\xb5\xca\xf3\xe4\xbc\x59\x88\xb5\xa0\x98\xcc\x11\x4e\x78\x1f\x4e\x56\xae\x61\x5b\x02\xea\xba\x89\x68\xd6\x6c\x18\x97\x1a\x4e\x56\xf4\xba\x9c\xad\x97\x08\xa4\x11\x36\x69\xa2\xf8\x62\xf7\x92\x9d\xac\xa8\x9f\xac\xae\xf7\x29\xc4\x8d\xb7\x8d\xba\x03\xc1\x3f\x21\xbc\x7e\x49\x7a\x87\x8a\xf5\xc8\xcb\xd7\x1d\x35\x59\x8b\x32\xab\xbf\x7e\x7c\x56\x63\xff\xa9\xb4\xae\xb8\xfa\x76\x8f\xf0\x49\x47\x78\xc4\x01\x7f\xfd\x5b\x57\x0e\xeb\xab\xff\x5b\x54\x47\xf9\xc3\x51\x0e\xb1\x09\x19\x0d\x6f\x86\x33\x38\x27\x4f\x2d\xf4\xd7\x0a\xd5\x7a\x52\x3e\x7c\x89\x1e\x19\x06\x9d\x1a\x55\xa5\x66\x9a\x32\xe9\x2e\xec\xbf\xda\xb2\x5f\xf2\x5f\x01\x00\x00\xff\xff\xe2\xf4\x38\xce\x13\x0a\x00\x00")

func modelDatabaseTableGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_modelDatabaseTableGoTmpl,
		"model/database/table.go.tmpl",
	)
}

func modelDatabaseTableGoTmpl() (*asset, error) {
	bytes, err := modelDatabaseTableGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "model/database/table.go.tmpl", size: 2579, mode: os.FileMode(438), modTime: time.Unix(1595309263, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _modelModelGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\xcd\xbd\x6a\xc3\x30\x14\xc5\xf1\x5d\xa0\x77\x38\x18\x0f\x2d\x14\x79\x0f\x64\xa9\x4b\xbb\x65\xca\xde\xc8\xd6\x45\xa8\xd5\x17\xb2\xec\x62\x84\xde\xbd\xc8\xe9\xd2\xac\x7f\xe9\xfe\xce\x30\x40\x93\xa7\x24\x33\x29\x4c\x3b\xf4\xac\x39\x8b\x72\xfe\x96\x9a\x50\xca\xea\x15\x25\x6b\x3c\x41\xbc\xbd\x5e\xa4\xa3\x5a\x39\xe3\xcc\xb8\x18\x52\xc6\x13\x67\x40\x97\x8d\xa3\x8e\xb3\x67\xce\x36\x99\xf0\x89\x33\x5a\x11\x97\xf0\xd3\xbe\xe6\x3d\x1e\x50\x8c\x94\x46\xe9\xc8\x42\x5c\xe5\x64\xe9\x8e\x61\xc9\x69\x9d\x33\x4a\x93\x4a\x41\x92\x5e\x13\x7a\xf3\x82\x7e\xc3\xe9\x0c\x31\x06\xbb\x3a\xbf\xa0\xed\xe2\xbf\xd3\x6f\xe2\xdd\x90\x55\xb5\xa2\x94\x7e\x13\x1f\xe1\xba\xc7\x66\xde\xd4\x74\xea\x8e\xf4\xf7\xde\xe1\x6b\x09\xfe\xa1\xdd\x80\x61\xb8\x5f\x8e\xc1\x39\xf2\xf9\x80\xc8\xab\xb6\x55\x39\xfb\x0d\x00\x00\xff\xff\x2a\xcf\x78\xea\x1b\x01\x00\x00")

func modelModelGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_modelModelGoTmpl,
		"model/model.go.tmpl",
	)
}

func modelModelGoTmpl() (*asset, error) {
	bytes, err := modelModelGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "model/model.go.tmpl", size: 283, mode: os.FileMode(438), modTime: time.Unix(1562730696, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _routerDuplicatequestGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x54\x4d\x8f\xe3\x44\x10\x3d\x3b\x52\xfe\x43\x6d\x24\x50\x1b\xa2\x36\x20\x86\x43\x50\xb8\x2c\x0b\x83\x04\xa3\x11\x99\x15\x12\xa3\x11\xea\xb4\x2b\x71\xcb\x76\xdb\x54\x97\xe3\x31\x28\x37\x18\xed\x0d\xae\x70\xe2\xc2\x0d\xa4\x39\x72\x58\xfe\x0d\x59\xf8\x19\xa8\xfd\x91\x49\x66\xe6\xb0\xb9\x38\x5d\xae\x7a\xef\xd5\xab\x72\x47\x11\xac\xd1\x22\x29\x46\x58\x36\xb0\xd6\xeb\xf1\xa8\x54\x3a\x55\x6b\x04\x2a\x2a\x46\x1a\x8f\xc6\x23\x93\x97\x05\x31\x88\xf1\x28\x98\xb0\xc9\x71\xe2\xff\xac\x0d\x27\xd5\x52\xea\x22\x8f\x32\xb5\x74\xac\x74\x1a\xa1\x4e\x8a\xf6\xa5\x45\x8e\x12\xe6\xb2\x3d\xac\x72\x6e\x9f\x9a\x9a\x92\x8b\x28\x8f\x4f\xda\xa3\xf1\xa9\xa1\xc7\xe7\xa6\xc4\x16\xfd\xe3\xaa\xcc\x8c\xf6\x62\x1c\x53\xa5\x19\xbe\x1f\x8f\x82\xe0\x4b\xfc\xb6\x42\xc7\x9f\x1b\xc7\xe0\x7f\xb9\x2a\x2f\x1d\x93\xb1\xeb\x2b\xaf\x46\x5e\x98\x1c\x7d\xde\x53\xa5\x13\xfc\x62\x81\xba\xb0\x71\x60\x2c\x7f\xf0\x7e\x1b\x4d\x50\xa7\xbe\x36\xb8\xbc\xea\xaa\xc6\xa3\x60\xdb\x33\xaf\x2a\xab\xe1\x0c\xeb\x3d\xb1\x13\xda\xa3\xe4\x2e\x86\x16\x61\xaa\x7d\x79\xe6\xa9\x87\xf2\x10\xde\xba\xd3\xe9\x05\x12\x72\x45\x16\xde\xdc\x47\xef\xab\x9e\xe5\x2a\x45\xf1\x98\xec\x70\x7a\x5f\xf8\x6c\xe0\x9f\x1e\x89\x9f\xed\x75\xb4\xf1\xed\x78\xb4\x1d\x8f\xa2\x08\xce\xa9\xd0\xe8\x1c\x18\x07\x9c\x20\xe4\x26\x8e\x33\xac\x15\x21\xf8\xd6\xd8\x14\x56\xf6\x5d\x0a\x77\xa0\x3b\x1c\x0a\x85\xc5\x6b\x06\x3f\x37\x79\xaa\x6c\x9c\x21\x7d\x52\x59\x1d\x3e\x88\x1c\x36\xea\xe1\x84\xee\x52\x9e\x16\x96\xf1\x9a\x43\x40\xa2\x82\xba\x79\xad\x0a\x82\x6f\xa6\x1b\x98\xcd\x49\xd9\x35\x82\x93\xfb\x36\xda\xf7\x81\x59\xc1\x66\x3e\xd7\xf2\x5c\x71\x22\xc2\x2e\x16\x44\x91\xce\x50\x11\x54\x25\xbc\x07\x0e\x35\xda\x18\x94\x8d\xc1\xa9\x0d\x02\x75\x66\x76\x99\x1e\x3f\x9d\xf2\x01\xfe\x81\xd7\x3d\x5a\xb0\xca\x59\x9e\x93\xb1\x9c\x59\xd1\xda\x7d\x56\xd4\x22\x94\x8b\x6a\x29\x38\x0c\xfb\xa4\xb8\x22\xe5\x3d\x82\xd9\x1c\x1e\x24\xf5\x39\x65\x0f\x32\xe4\x0e\x71\xb3\xea\x36\xe4\x21\x7a\xf4\xee\x3b\xed\x2f\xfc\xc8\xc9\xc3\xd1\x0e\xd2\x82\x18\x33\x64\x14\x47\xba\xa7\x90\x0e\xc8\xdb\xee\xb9\x1d\x7c\x31\x0e\x34\x29\x97\x74\xe7\x14\x1b\x2f\x57\x0f\xc5\x22\x94\xa7\xa8\x62\x24\xf9\x29\xb2\x98\xf4\xd1\x14\x9b\x49\x8f\x67\x56\x90\x62\x33\x9f\x4f\x26\x83\x80\x7e\x8e\xed\x00\x9f\x11\x3d\xb7\xaa\xe2\xa4\x20\xf3\x1d\xc6\x47\xd4\xb5\x27\xca\xe3\x13\x79\x86\xb5\xe8\xd1\x0e\x7d\x1d\x26\x38\x10\x15\xf2\x2b\x32\x8c\x8b\x76\xc9\x45\x3d\x85\x21\xe1\xed\x14\x9b\x10\x00\xa2\x68\x77\x7b\xe3\x98\x76\x37\xbf\xec\x7e\xfc\x7d\xf7\xe2\xb6\xfe\xe7\xaf\x3f\xbb\xe2\x3c\x3e\x49\xb1\x99\xcd\xc1\x13\x2c\x5a\xd3\x57\x62\xf2\xc6\xf5\x64\x0a\xb5\x5c\x54\xb9\xb0\x26\x0b\x43\x0f\x71\x77\xdc\xdd\xde\xd4\xff\xfe\xfa\x43\xa2\x5c\xf2\xdf\xdf\x7f\xbc\x7a\xf1\xf3\xe5\xd5\xb2\x61\x7c\xf5\xdb\xcb\xdd\xcb\x9f\xfc\xe7\xdd\xb7\xff\xe4\xc8\xe9\xcb\x8e\xea\x4a\x7e\xe6\xbe\x46\x2a\x44\x08\x8f\xf9\x72\x86\xf5\xe9\xc5\xc5\xf9\x33\xbf\xd6\xc2\x5f\x65\x72\xc1\x8a\x2b\xb7\x40\xda\x18\x8d\xcf\xad\xda\x28\x93\xa9\x65\x86\x7d\xf7\x5b\xc0\xcc\xe1\x1e\xeb\x51\x4a\x38\x5c\xb3\xbb\x3d\x42\x22\xef\xb4\xff\x14\x85\x0e\x3f\x6c\xcf\x4f\xe6\x60\x4d\xb6\x87\x0b\xb4\xec\xa4\x20\xd1\xbd\x3d\x19\x64\x5b\x93\x1d\x4e\x6f\xdb\xdf\x13\xaf\xc7\x70\x1f\xbe\x2d\x3c\x02\xee\xae\x9c\xff\x03\x00\x00\xff\xff\x22\xee\x6a\xc6\x31\x06\x00\x00")

func routerDuplicatequestGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_routerDuplicatequestGoTmpl,
		"router/duplicateQuest.go.tmpl",
	)
}

func routerDuplicatequestGoTmpl() (*asset, error) {
	bytes, err := routerDuplicatequestGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "router/duplicateQuest.go.tmpl", size: 1585, mode: os.FileMode(438), modTime: time.Unix(1562726307, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _routerGen_routerGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\xd2\xd1\x4a\xfb\x30\x14\xc7\xf1\xeb\x05\xf2\x0e\x87\x5e\x6d\x7f\xfe\x26\xcf\x30\x66\x1d\xde\xe8\xa0\x13\xaf\xd3\xf4\xac\xad\x4b\x9b\x9a\xa4\xc2\x08\x7d\x77\xc9\x02\x95\x81\x9b\x01\xbd\xff\xf5\xd3\x03\xdf\x70\x0e\x35\xf6\x68\x84\x43\x28\x4f\x50\xcb\x9a\x92\x41\xc8\xa3\xa8\x11\x8c\x1e\x1d\x1a\x4a\x28\x69\xbb\x41\x1b\x07\x4b\x4a\x00\x00\x32\xef\xd9\xce\xe8\x37\x94\xee\x49\x74\x38\x4d\x5c\xea\xde\x19\xad\x14\x9a\x2c\x4c\x16\x59\xdd\xba\x66\x2c\x99\xd4\x1d\x57\xa2\xb4\x4e\xc8\x23\x47\xd9\xe8\x8c\x92\x55\x00\x0f\x63\x2f\xc1\xfb\x71\x18\xd0\x6c\x44\x87\x0a\xd8\x5e\x94\x0a\xa3\xb7\xae\xaa\x25\xc2\xbf\xf0\x01\xcb\x65\xa3\x57\x3e\xfe\x98\xf3\xee\x64\xdf\xd5\x7c\xd8\x02\xd9\xee\xb9\xd8\x2f\x33\x7e\x95\x2a\xc4\x07\x66\xff\xbf\xee\x63\x37\x97\xab\x24\xf3\x1e\x15\xba\x44\x35\x6e\xa3\xbb\xcd\x03\xbb\x45\x77\x75\xfd\xda\xa0\xb9\x84\x7f\x5c\x5f\x9c\x7c\x6b\xbd\x56\x2a\x59\x5e\x2b\x95\x7a\xf2\x43\x6b\xac\x4b\x86\xcf\xeb\x40\x7b\xdf\x1e\x80\x3d\xda\xbc\x98\x26\x4a\x16\x9c\xa3\xfd\x26\x6b\x68\x72\xd5\x4a\x0b\x90\x17\x1b\x83\xe2\x9c\x20\x3c\xa1\x99\x8e\x61\x7e\x8d\xcf\x7d\x2f\xf0\x97\xa1\x12\x7f\x80\x47\x66\xc6\xe7\x18\x36\x4d\xce\x8b\x9b\xdb\x98\xe1\x0e\xb0\xaf\x42\x83\x89\x92\xcf\x00\x00\x00\xff\xff\xe4\xa0\x6a\xdf\x0d\x04\x00\x00")

func routerGen_routerGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_routerGen_routerGoTmpl,
		"router/gen_router.go.tmpl",
	)
}

func routerGen_routerGoTmpl() (*asset, error) {
	bytes, err := routerGen_routerGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "router/gen_router.go.tmpl", size: 1037, mode: os.FileMode(438), modTime: time.Unix(1595298897, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _routerRouterGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x90\x4d\x8b\xdb\x30\x10\x86\xcf\x16\xe8\x3f\x4c\x45\x0e\x0a\x04\xe9\x1e\x48\xa1\x84\x7e\x40\xcb\x6e\x89\xb7\xf4\xd2\x8b\x22\x4f\x15\xb3\xb2\x64\xc6\x63\x6f\x8b\xd1\x7f\x2f\x72\x0a\xbd\xb4\xf4\x22\x18\x9e\x77\x98\x47\xaf\xb5\x10\x30\x21\x39\x46\xb8\xfe\x84\xe0\x83\x14\xa3\xf3\xcf\x2e\x20\x50\x9e\x19\x49\x0a\x29\xfa\x61\xcc\xc4\xa0\xa5\x00\x00\x50\x09\xd9\xde\x98\x47\x55\xe7\x46\x85\x9e\x6f\xf3\xd5\xf8\x3c\xd8\xe8\xae\x13\x3b\xff\x6c\xd1\xdf\xf2\x7f\xb0\x1d\xfa\xae\x8b\xf8\xe2\x08\x95\x14\xfb\x7a\xe7\xfb\x9c\x3c\x5c\xb6\xb3\x7a\xbf\x4a\xd1\xe0\xf1\x54\xa3\xe6\x01\x5f\xf4\xbe\xce\xe6\xcb\x84\xfa\xcf\xa2\xf9\x94\x43\xa8\xe1\xbf\xc2\x0b\xfa\xbc\xfc\x93\x9e\x1f\x2f\xed\x86\xa4\x68\xac\xbd\x9f\x85\xd3\x6b\xb8\xb9\xd4\xc5\xfa\xef\x06\xcd\xfb\xb7\x4f\x5a\x59\x75\x80\x6a\xa6\x3d\x6c\x32\xe7\x9c\x18\x7f\xf0\x1e\x90\x28\x13\x54\xcf\x86\x90\x67\x4a\xe0\x4d\xcb\xd4\xa7\xa0\x6b\x3d\xa6\x65\xc7\xf3\xf4\xf8\xf1\x00\xea\x03\xc6\x98\x0f\xf0\x35\x53\xec\x5e\x7d\x4b\xaa\x1a\x95\xfa\xac\x2b\x90\x4b\x01\x61\xd7\x1f\x60\xb7\xc0\xf1\x04\xe6\xc9\x5d\x23\x3e\xb8\x01\xa7\x52\xb6\xc8\x3c\x8e\x48\x67\x37\x60\x84\xdd\x52\xca\x9b\xae\xd3\x78\x5f\xc6\xd4\x6d\x19\x6b\xa1\x65\x47\x0c\x13\xd2\xf2\x5b\xfe\xde\x8d\x79\xe7\xd8\x45\x8d\x66\xe3\x5a\x1d\xd7\xd5\xb4\x48\x4b\xef\xf1\x73\x26\x2e\x45\xd5\x0e\x8a\x14\xbf\x02\x00\x00\xff\xff\x5a\x31\x01\x6f\x0b\x02\x00\x00")

func routerRouterGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_routerRouterGoTmpl,
		"router/router.go.tmpl",
	)
}

func routerRouterGoTmpl() (*asset, error) {
	bytes, err := routerRouterGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "router/router.go.tmpl", size: 523, mode: os.FileMode(438), modTime: time.Unix(1562726307, 0)}
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
	"controller/controller.go.tmpl":        controllerControllerGoTmpl,
	"controller/gen_controller.go.tmpl":    controllerGen_controllerGoTmpl,
	"controller/gen_es_controller.go.tmpl": controllerGen_es_controllerGoTmpl,
	"main.go.tmpl":                         mainGoTmpl,
	"model/database/db.go.tmpl":            modelDatabaseDbGoTmpl,
	"model/database/es.go.tmpl":            modelDatabaseEsGoTmpl,
	"model/database/esquery.go.tmpl":       modelDatabaseEsqueryGoTmpl,
	"model/database/table.go.tmpl":         modelDatabaseTableGoTmpl,
	"model/model.go.tmpl":                  modelModelGoTmpl,
	"router/duplicateQuest.go.tmpl":        routerDuplicatequestGoTmpl,
	"router/gen_router.go.tmpl":            routerGen_routerGoTmpl,
	"router/router.go.tmpl":                routerRouterGoTmpl,
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
	"controller": &bintree{nil, map[string]*bintree{
		"controller.go.tmpl":        &bintree{controllerControllerGoTmpl, map[string]*bintree{}},
		"gen_controller.go.tmpl":    &bintree{controllerGen_controllerGoTmpl, map[string]*bintree{}},
		"gen_es_controller.go.tmpl": &bintree{controllerGen_es_controllerGoTmpl, map[string]*bintree{}},
	}},
	"main.go.tmpl": &bintree{mainGoTmpl, map[string]*bintree{}},
	"model": &bintree{nil, map[string]*bintree{
		"database": &bintree{nil, map[string]*bintree{
			"db.go.tmpl":      &bintree{modelDatabaseDbGoTmpl, map[string]*bintree{}},
			"es.go.tmpl":      &bintree{modelDatabaseEsGoTmpl, map[string]*bintree{}},
			"esquery.go.tmpl": &bintree{modelDatabaseEsqueryGoTmpl, map[string]*bintree{}},
			"table.go.tmpl":   &bintree{modelDatabaseTableGoTmpl, map[string]*bintree{}},
		}},
		"model.go.tmpl": &bintree{modelModelGoTmpl, map[string]*bintree{}},
	}},
	"router": &bintree{nil, map[string]*bintree{
		"duplicateQuest.go.tmpl": &bintree{routerDuplicatequestGoTmpl, map[string]*bintree{}},
		"gen_router.go.tmpl":     &bintree{routerGen_routerGoTmpl, map[string]*bintree{}},
		"router.go.tmpl":         &bintree{routerRouterGoTmpl, map[string]*bintree{}},
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
