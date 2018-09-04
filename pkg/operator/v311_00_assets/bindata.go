// Code generated by go-bindata.
// sources:
// manifests/v3.11.0/kube-apiserver/cm.yaml
// manifests/v3.11.0/kube-apiserver/defaultconfig.yaml
// manifests/v3.11.0/kube-apiserver/deployment.yaml
// manifests/v3.11.0/kube-apiserver/ns.yaml
// manifests/v3.11.0/kube-apiserver/sa.yaml
// manifests/v3.11.0/kube-apiserver/svc.yaml
// DO NOT EDIT!

package v311_00_assets

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)
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

var _v3110KubeApiserverCmYaml = []byte(`apiVersion: v1
kind: ConfigMap
metadata:
  namespace: openshift-kube-apiserver
  name: deployment-apiserver-config
data:
  config.yaml:
`)

func v3110KubeApiserverCmYamlBytes() ([]byte, error) {
	return _v3110KubeApiserverCmYaml, nil
}

func v3110KubeApiserverCmYaml() (*asset, error) {
	bytes, err := v3110KubeApiserverCmYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/kube-apiserver/cm.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110KubeApiserverDefaultconfigYaml = []byte(`# TODO this will change to KubeAPIServerConfig
apiVersion: v1
kind: MasterConfig
signer:
  certFile: /var/run/secrets/signing-key/tls.crt
  keyFile: /var/run/secrets/signing-key/tls.key
`)

func v3110KubeApiserverDefaultconfigYamlBytes() ([]byte, error) {
	return _v3110KubeApiserverDefaultconfigYaml, nil
}

func v3110KubeApiserverDefaultconfigYaml() (*asset, error) {
	bytes, err := v3110KubeApiserverDefaultconfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/kube-apiserver/defaultconfig.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110KubeApiserverDeploymentYaml = []byte(`apiVersion: apps/v1
kind: Deployment
metadata:
  namespace: openshift-kube-apiserver
  name: apiserver
  labels:
    app: openshift-kube-apiserver
    apiserver: "true"
spec:
  strategy:
    type: RollingUpdate
  selector:
    matchLabels:
      app: openshift-kube-apiserver
      apiserver: "true"
  template:
    metadata:
      name: openshift-kube-apiserver
      labels:
        app: openshift-kube-apiserver
        apiserver: "true"
    spec:
      serviceAccountName: openshift-kube-apiserver-sa
      containers:
      - name: apiserver
        image: ${IMAGE}
        imagePullPolicy: IfNotPresent
        command: ["hypershift", "openshift-kube-apiserver"]
        args:
        - "--config=/var/run/configmaps/config/config.yaml"
        ports:
        - containerPort: 8443
        volumeMounts:
        - mountPath: /var/run/configmaps/config
          name: config
        - mountPath: /var/run/secrets/serving-cert
          name: serving-cert
      volumes:
      - name: serving-cert
        secret:
          secretName: serving-cert
      - name: config
        configMap:
          name: deployment-apiserver-config



`)

func v3110KubeApiserverDeploymentYamlBytes() ([]byte, error) {
	return _v3110KubeApiserverDeploymentYaml, nil
}

func v3110KubeApiserverDeploymentYaml() (*asset, error) {
	bytes, err := v3110KubeApiserverDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/kube-apiserver/deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110KubeApiserverNsYaml = []byte(`apiVersion: v1
kind: Namespace
metadata:
  name: openshift-kube-apiserver
  labels:
    openshift.io/run-level: "0"`)

func v3110KubeApiserverNsYamlBytes() ([]byte, error) {
	return _v3110KubeApiserverNsYaml, nil
}

func v3110KubeApiserverNsYaml() (*asset, error) {
	bytes, err := v3110KubeApiserverNsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/kube-apiserver/ns.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110KubeApiserverSaYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  namespace: openshift-kube-apiserver
  name: openshift-kube-apiserver-sa
`)

func v3110KubeApiserverSaYamlBytes() ([]byte, error) {
	return _v3110KubeApiserverSaYaml, nil
}

func v3110KubeApiserverSaYaml() (*asset, error) {
	bytes, err := v3110KubeApiserverSaYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/kube-apiserver/sa.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110KubeApiserverSvcYaml = []byte(`apiVersion: v1
kind: Service
metadata:
  namespace: openshift-kube-apiserver
  name: apiserver
  annotations:
    service.alpha.openshift.io/serving-cert-secret-name: serving-cert
    prometheus.io/scrape: "true"
    prometheus.io/scheme: https
spec:
  selector:
    apiserver: "true"
  ports:
  - name: https
    port: 443
    targetPort: 8443
`)

func v3110KubeApiserverSvcYamlBytes() ([]byte, error) {
	return _v3110KubeApiserverSvcYaml, nil
}

func v3110KubeApiserverSvcYaml() (*asset, error) {
	bytes, err := v3110KubeApiserverSvcYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/kube-apiserver/svc.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"v3.11.0/kube-apiserver/cm.yaml": v3110KubeApiserverCmYaml,
	"v3.11.0/kube-apiserver/defaultconfig.yaml": v3110KubeApiserverDefaultconfigYaml,
	"v3.11.0/kube-apiserver/deployment.yaml": v3110KubeApiserverDeploymentYaml,
	"v3.11.0/kube-apiserver/ns.yaml": v3110KubeApiserverNsYaml,
	"v3.11.0/kube-apiserver/sa.yaml": v3110KubeApiserverSaYaml,
	"v3.11.0/kube-apiserver/svc.yaml": v3110KubeApiserverSvcYaml,
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
	"v3.11.0": &bintree{nil, map[string]*bintree{
		"kube-apiserver": &bintree{nil, map[string]*bintree{
			"cm.yaml": &bintree{v3110KubeApiserverCmYaml, map[string]*bintree{}},
			"defaultconfig.yaml": &bintree{v3110KubeApiserverDefaultconfigYaml, map[string]*bintree{}},
			"deployment.yaml": &bintree{v3110KubeApiserverDeploymentYaml, map[string]*bintree{}},
			"ns.yaml": &bintree{v3110KubeApiserverNsYaml, map[string]*bintree{}},
			"sa.yaml": &bintree{v3110KubeApiserverSaYaml, map[string]*bintree{}},
			"svc.yaml": &bintree{v3110KubeApiserverSvcYaml, map[string]*bintree{}},
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

