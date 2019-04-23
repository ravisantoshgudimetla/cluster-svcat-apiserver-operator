// Code generated by go-bindata.
// sources:
// bindata/v3.11.0/openshift-svcat-apiserver/cm.yaml
// bindata/v3.11.0/openshift-svcat-apiserver/cr-aggregate-to-admin.yaml
// bindata/v3.11.0/openshift-svcat-apiserver/cr-aggregate-to-edit.yaml
// bindata/v3.11.0/openshift-svcat-apiserver/cr-aggregate-to-view.yaml
// bindata/v3.11.0/openshift-svcat-apiserver/cr-namespace-viewer.yaml
// bindata/v3.11.0/openshift-svcat-apiserver/cr-readiness-probe.yaml
// bindata/v3.11.0/openshift-svcat-apiserver/cr-sar-creator.yaml
// bindata/v3.11.0/openshift-svcat-apiserver/cr-serviceclass-viewer.yaml
// bindata/v3.11.0/openshift-svcat-apiserver/crb-auth-delegator-binding.yaml
// bindata/v3.11.0/openshift-svcat-apiserver/crb-namespace-viewer-binding.yaml
// bindata/v3.11.0/openshift-svcat-apiserver/crb-readiness-binding.yaml
// bindata/v3.11.0/openshift-svcat-apiserver/crb-sar-creator-binding.yaml
// bindata/v3.11.0/openshift-svcat-apiserver/crb-serviceclass-viewer-binding.yaml
// bindata/v3.11.0/openshift-svcat-apiserver/ds.yaml
// bindata/v3.11.0/openshift-svcat-apiserver/ns.yaml
// bindata/v3.11.0/openshift-svcat-apiserver/role-extension-apiserver-auth-reader.yaml
// bindata/v3.11.0/openshift-svcat-apiserver/rolebinding-extension-apiserver-auth-reader.yaml
// bindata/v3.11.0/openshift-svcat-apiserver/sa.yaml
// bindata/v3.11.0/openshift-svcat-apiserver/svc.yaml
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

var _v3110OpenshiftSvcatApiserverCmYaml = []byte(`apiVersion: v1
kind: ConfigMap
metadata:
  namespace: openshift-service-catalog-apiserver
  name: config
data:
  config.yaml:
`)

func v3110OpenshiftSvcatApiserverCmYamlBytes() ([]byte, error) {
	return _v3110OpenshiftSvcatApiserverCmYaml, nil
}

func v3110OpenshiftSvcatApiserverCmYaml() (*asset, error) {
	bytes, err := v3110OpenshiftSvcatApiserverCmYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/openshift-svcat-apiserver/cm.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110OpenshiftSvcatApiserverCrAggregateToAdminYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    rbac.authorization.k8s.io/aggregate-to-admin: "true"
  name: system:service-catalog:aggregate-to-admin
rules:
- apiGroups:
  - "servicecatalog.k8s.io"
  attributeRestrictions: null
  resources:
  - servicebrokers
  - serviceclasses
  - serviceplans
  - serviceinstances
  - servicebindings
  verbs:
  - create
  - update
  - delete
  - get
  - list
  - watch
  - patch
- apiGroups:
  - "settings.k8s.io"
  attributeRestrictions: null
  resources:
  - podpresets
  verbs:
  - create
  - update
  - delete
  - get
  - list
  - watch
`)

func v3110OpenshiftSvcatApiserverCrAggregateToAdminYamlBytes() ([]byte, error) {
	return _v3110OpenshiftSvcatApiserverCrAggregateToAdminYaml, nil
}

func v3110OpenshiftSvcatApiserverCrAggregateToAdminYaml() (*asset, error) {
	bytes, err := v3110OpenshiftSvcatApiserverCrAggregateToAdminYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/openshift-svcat-apiserver/cr-aggregate-to-admin.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110OpenshiftSvcatApiserverCrAggregateToEditYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    rbac.authorization.k8s.io/aggregate-to-edit: "true"
  name: system:service-catalog:aggregate-to-edit
rules:
- apiGroups:
  - "servicecatalog.k8s.io"
  attributeRestrictions: null
  resources:
  - servicebrokers
  - serviceclasses
  - serviceplans
  - serviceinstances
  - servicebindings
  verbs:
  - create
  - update
  - delete
  - get
  - list
  - watch
  - patch
- apiGroups:
  - "settings.k8s.io"
  attributeRestrictions: null
  resources:
  - podpresets
  verbs:
  - create
  - update
  - delete
  - get
  - list
  - watch
`)

func v3110OpenshiftSvcatApiserverCrAggregateToEditYamlBytes() ([]byte, error) {
	return _v3110OpenshiftSvcatApiserverCrAggregateToEditYaml, nil
}

func v3110OpenshiftSvcatApiserverCrAggregateToEditYaml() (*asset, error) {
	bytes, err := v3110OpenshiftSvcatApiserverCrAggregateToEditYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/openshift-svcat-apiserver/cr-aggregate-to-edit.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110OpenshiftSvcatApiserverCrAggregateToViewYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    rbac.authorization.k8s.io/aggregate-to-view: "true"
  name: system:service-catalog:aggregate-to-view
rules:
- apiGroups:
  - "servicecatalog.k8s.io"
  attributeRestrictions: null
  resources:
  - servicebrokers
  - serviceclasses
  - serviceplans
  - serviceinstances
  - servicebindings
  verbs:
  - get
  - list
  - watch
`)

func v3110OpenshiftSvcatApiserverCrAggregateToViewYamlBytes() ([]byte, error) {
	return _v3110OpenshiftSvcatApiserverCrAggregateToViewYaml, nil
}

func v3110OpenshiftSvcatApiserverCrAggregateToViewYaml() (*asset, error) {
	bytes, err := v3110OpenshiftSvcatApiserverCrAggregateToViewYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/openshift-svcat-apiserver/cr-aggregate-to-view.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110OpenshiftSvcatApiserverCrNamespaceViewerYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: namespace-viewer
rules:
- apiGroups:
  - ""
  resources:
  - namespaces
  verbs:
  - list
  - watch
  - get
- apiGroups:
  - "admissionregistration.k8s.io"
  resources:
  - validatingwebhookconfigurations
  - mutatingwebhookconfigurations
  verbs:
  - list
  - watch
  - get
`)

func v3110OpenshiftSvcatApiserverCrNamespaceViewerYamlBytes() ([]byte, error) {
	return _v3110OpenshiftSvcatApiserverCrNamespaceViewerYaml, nil
}

func v3110OpenshiftSvcatApiserverCrNamespaceViewerYaml() (*asset, error) {
	bytes, err := v3110OpenshiftSvcatApiserverCrNamespaceViewerYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/openshift-svcat-apiserver/cr-namespace-viewer.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110OpenshiftSvcatApiserverCrReadinessProbeYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: servicecatalog.k8s.io:service-catalog-readiness
rules:
- nonResourceURLs:
  - /healthz/ready
  verbs:
    - get
`)

func v3110OpenshiftSvcatApiserverCrReadinessProbeYamlBytes() ([]byte, error) {
	return _v3110OpenshiftSvcatApiserverCrReadinessProbeYaml, nil
}

func v3110OpenshiftSvcatApiserverCrReadinessProbeYaml() (*asset, error) {
	bytes, err := v3110OpenshiftSvcatApiserverCrReadinessProbeYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/openshift-svcat-apiserver/cr-readiness-probe.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110OpenshiftSvcatApiserverCrSarCreatorYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: sar-creator
rules:
- apiGroups:
  - ""
  resources:
  - subjectaccessreviews.authorization.k8s.io
  verbs:
  - create
`)

func v3110OpenshiftSvcatApiserverCrSarCreatorYamlBytes() ([]byte, error) {
	return _v3110OpenshiftSvcatApiserverCrSarCreatorYaml, nil
}

func v3110OpenshiftSvcatApiserverCrSarCreatorYaml() (*asset, error) {
	bytes, err := v3110OpenshiftSvcatApiserverCrSarCreatorYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/openshift-svcat-apiserver/cr-sar-creator.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110OpenshiftSvcatApiserverCrServiceclassViewerYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: servicecatalog-serviceclass-viewer
rules:
- apiGroups:
  - servicecatalog.k8s.io
  resources:
  - clusterserviceclasses
  - clusterserviceplans
  verbs:
  - list
  - watch
  - get
`)

func v3110OpenshiftSvcatApiserverCrServiceclassViewerYamlBytes() ([]byte, error) {
	return _v3110OpenshiftSvcatApiserverCrServiceclassViewerYaml, nil
}

func v3110OpenshiftSvcatApiserverCrServiceclassViewerYaml() (*asset, error) {
	bytes, err := v3110OpenshiftSvcatApiserverCrServiceclassViewerYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/openshift-svcat-apiserver/cr-serviceclass-viewer.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110OpenshiftSvcatApiserverCrbAuthDelegatorBindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: system:auth-delegator-binding
roleRef:
  kind: ClusterRole
  name: system:auth-delegator
subjects:
- kind: ServiceAccount
  name: service-catalog-apiserver
  namespace: openshift-service-catalog-apiserver
`)

func v3110OpenshiftSvcatApiserverCrbAuthDelegatorBindingYamlBytes() ([]byte, error) {
	return _v3110OpenshiftSvcatApiserverCrbAuthDelegatorBindingYaml, nil
}

func v3110OpenshiftSvcatApiserverCrbAuthDelegatorBindingYaml() (*asset, error) {
	bytes, err := v3110OpenshiftSvcatApiserverCrbAuthDelegatorBindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/openshift-svcat-apiserver/crb-auth-delegator-binding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110OpenshiftSvcatApiserverCrbNamespaceViewerBindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: service-catalog-namespace-viewer-binding
roleRef:
  kind: ClusterRole
  name: namespace-viewer
subjects:
- kind: ServiceAccount
  name: service-catalog-apiserver
  namespace: openshift-service-catalog-apiserver
`)

func v3110OpenshiftSvcatApiserverCrbNamespaceViewerBindingYamlBytes() ([]byte, error) {
	return _v3110OpenshiftSvcatApiserverCrbNamespaceViewerBindingYaml, nil
}

func v3110OpenshiftSvcatApiserverCrbNamespaceViewerBindingYaml() (*asset, error) {
	bytes, err := v3110OpenshiftSvcatApiserverCrbNamespaceViewerBindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/openshift-svcat-apiserver/crb-namespace-viewer-binding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110OpenshiftSvcatApiserverCrbReadinessBindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: servicecatalog.k8s.io:service-catalog-readiness
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: servicecatalog.k8s.io:service-catalog-readiness
subjects:
- apiGroup: rbac.authorization.k8s.io
  kind: Group
  name: system:unauthenticated
- apiGroup: rbac.authorization.k8s.io
  kind: Group
  name: system:authenticated
`)

func v3110OpenshiftSvcatApiserverCrbReadinessBindingYamlBytes() ([]byte, error) {
	return _v3110OpenshiftSvcatApiserverCrbReadinessBindingYaml, nil
}

func v3110OpenshiftSvcatApiserverCrbReadinessBindingYaml() (*asset, error) {
	bytes, err := v3110OpenshiftSvcatApiserverCrbReadinessBindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/openshift-svcat-apiserver/crb-readiness-binding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110OpenshiftSvcatApiserverCrbSarCreatorBindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: service-catalog-sar-creator-binding
roleRef:
  kind: ClusterRole
  name: sar-creator
subjects:
- kind: ServiceAccount
  name: service-catalog-apiserver
  namespace: openshift-service-catalog-apiserver
`)

func v3110OpenshiftSvcatApiserverCrbSarCreatorBindingYamlBytes() ([]byte, error) {
	return _v3110OpenshiftSvcatApiserverCrbSarCreatorBindingYaml, nil
}

func v3110OpenshiftSvcatApiserverCrbSarCreatorBindingYaml() (*asset, error) {
	bytes, err := v3110OpenshiftSvcatApiserverCrbSarCreatorBindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/openshift-svcat-apiserver/crb-sar-creator-binding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110OpenshiftSvcatApiserverCrbServiceclassViewerBindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: servicecatalog-serviceclass-viewer-binding
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: servicecatalog-serviceclass-viewer
subjects:
- apiGroup: rbac.authorization.k8s.io
  kind: Group
  name: system:authenticated
`)

func v3110OpenshiftSvcatApiserverCrbServiceclassViewerBindingYamlBytes() ([]byte, error) {
	return _v3110OpenshiftSvcatApiserverCrbServiceclassViewerBindingYaml, nil
}

func v3110OpenshiftSvcatApiserverCrbServiceclassViewerBindingYaml() (*asset, error) {
	bytes, err := v3110OpenshiftSvcatApiserverCrbServiceclassViewerBindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/openshift-svcat-apiserver/crb-serviceclass-viewer-binding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110OpenshiftSvcatApiserverDsYaml = []byte(`apiVersion: "apps/v1"
kind: DaemonSet
metadata:
  labels:
    app: apiserver
  name: apiserver
  namespace: openshift-service-catalog-apiserver
spec:
  selector:
    matchLabels:
      app: apiserver
  updateStrategy:
    rollingUpdate:
      maxUnavailable: 1
    type: RollingUpdate
  template:
    metadata:
      labels:
        app: apiserver
    spec:
      serviceAccountName: service-catalog-apiserver
      nodeSelector:
        node-role.kubernetes.io/master: ""
      priorityClassName: "system-cluster-critical"
      containers:
      - args:
        - apiserver
        - --storage-type
        - etcd
        - --secure-port
        - "6443"
        - --etcd-servers
        - "https://etcd.kube-system.svc.cluster.local:2379"
        - --etcd-cafile
        - "/var/run/configmaps/etcd-serving-ca/ca-bundle.crt"
        - --etcd-certfile
        - "/var/run/secrets/etcd-client/tls.crt"
        - --etcd-keyfile
        - "/var/run/secrets/etcd-client/tls.key"
        - --cors-allowed-origins
        - "localhost"
        - --enable-admission-plugins
        - NamespaceLifecycle,DefaultServicePlan,ServiceBindingsLifecycle,ServicePlanChangeValidator,BrokerAuthSarCheck
        - --feature-gates
        - OriginatingIdentity=true
        - --feature-gates
        - NamespacedServiceBroker=true
        image: ${IMAGE}
        imagePullPolicy: IfNotPresent
        command: ["/usr/bin/service-catalog"]
        name: apiserver
        ports:
        - containerPort: 6443
          protocol: TCP
        resources:
          requests:
            memory: 200Mi
        terminationMessagePath: /dev/termination-log
        volumeMounts:
        - mountPath: /var/run/kubernetes-service-catalog
          name: apiserver-ssl
          readOnly: true
        - mountPath: /etc/origin/master
          name: etcd-host-cert
          readOnly: true
        - mountPath: /var/run/secrets/etcd-client
          name: etcd-client
        - mountPath: /var/run/configmaps/etcd-serving-ca
          name: etcd-serving-ca
        readinessProbe:
          httpGet:
            port: 6443
            path: /healthz/ready
            scheme: HTTPS
          failureThreshold: 3
          initialDelaySeconds: 20
          periodSeconds: 10
          successThreshold: 1
          timeoutSeconds: 10
        livenessProbe:
          httpGet:
            port: 6443
            path: /healthz
            scheme: HTTPS
          failureThreshold: 3
          initialDelaySeconds: 20
          periodSeconds: 10
          successThreshold: 1
          timeoutSeconds: 10
      dnsPolicy: ClusterFirst
      restartPolicy: Always
      securityContext: {}
      terminationGracePeriodSeconds: 30
      volumes:
      - name: apiserver-ssl
        secret:
          defaultMode: 420
          secretName: serving-cert
          items:
          - key: tls.crt
            path: apiserver.crt
          - key: tls.key
            path: apiserver.key
      - hostPath:
          path: /etc/origin/master
        name: etcd-host-cert
      - emptyDir: {}
        name: data-dir
      - name: etcd-serving-ca
        configMap:
          name: etcd-serving-ca
      - name: etcd-client
        secret:
          secretName: etcd-client
      tolerations:
      - operator: Exists`)

func v3110OpenshiftSvcatApiserverDsYamlBytes() ([]byte, error) {
	return _v3110OpenshiftSvcatApiserverDsYaml, nil
}

func v3110OpenshiftSvcatApiserverDsYaml() (*asset, error) {
	bytes, err := v3110OpenshiftSvcatApiserverDsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/openshift-svcat-apiserver/ds.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110OpenshiftSvcatApiserverNsYaml = []byte(`apiVersion: v1
kind: Namespace
metadata:
  name: openshift-service-catalog-apiserver
  labels:
    openshift.io/run-level: "1"`)

func v3110OpenshiftSvcatApiserverNsYamlBytes() ([]byte, error) {
	return _v3110OpenshiftSvcatApiserverNsYaml, nil
}

func v3110OpenshiftSvcatApiserverNsYaml() (*asset, error) {
	bytes, err := v3110OpenshiftSvcatApiserverNsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/openshift-svcat-apiserver/ns.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110OpenshiftSvcatApiserverRoleExtensionApiserverAuthReaderYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: extension-apiserver-authentication-reader
  namespace: kube-system
rules:
- apiGroups:
  - ""
  resourceNames:
  - extension-apiserver-authentication
  resources:
  - configmaps
  verbs:
  - get
`)

func v3110OpenshiftSvcatApiserverRoleExtensionApiserverAuthReaderYamlBytes() ([]byte, error) {
	return _v3110OpenshiftSvcatApiserverRoleExtensionApiserverAuthReaderYaml, nil
}

func v3110OpenshiftSvcatApiserverRoleExtensionApiserverAuthReaderYaml() (*asset, error) {
	bytes, err := v3110OpenshiftSvcatApiserverRoleExtensionApiserverAuthReaderYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/openshift-svcat-apiserver/role-extension-apiserver-auth-reader.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110OpenshiftSvcatApiserverRolebindingExtensionApiserverAuthReaderYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: extension-apiserver-authentication-reader-binding
  namespace: kube-system
roleRef:
  kind: Role
  name: extension-apiserver-authentication-reader
subjects:
- kind: ServiceAccount
  name: service-catalog-apiserver
  namespace: openshift-service-catalog-apiserver
`)

func v3110OpenshiftSvcatApiserverRolebindingExtensionApiserverAuthReaderYamlBytes() ([]byte, error) {
	return _v3110OpenshiftSvcatApiserverRolebindingExtensionApiserverAuthReaderYaml, nil
}

func v3110OpenshiftSvcatApiserverRolebindingExtensionApiserverAuthReaderYaml() (*asset, error) {
	bytes, err := v3110OpenshiftSvcatApiserverRolebindingExtensionApiserverAuthReaderYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/openshift-svcat-apiserver/rolebinding-extension-apiserver-auth-reader.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110OpenshiftSvcatApiserverSaYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  namespace: openshift-service-catalog-apiserver
  name: service-catalog-apiserver
`)

func v3110OpenshiftSvcatApiserverSaYamlBytes() ([]byte, error) {
	return _v3110OpenshiftSvcatApiserverSaYaml, nil
}

func v3110OpenshiftSvcatApiserverSaYaml() (*asset, error) {
	bytes, err := v3110OpenshiftSvcatApiserverSaYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/openshift-svcat-apiserver/sa.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110OpenshiftSvcatApiserverSvcYaml = []byte(`apiVersion: v1
kind: Service
metadata:
  namespace: openshift-service-catalog-apiserver
  name: api
  annotations:
    service.alpha.openshift.io/serving-cert-secret-name: serving-cert
    prometheus.io/scrape: "true"
    prometheus.io/scheme: https
spec:
  selector:
    app: apiserver
  ports:
  - name: https
    port: 443
    targetPort: 6443
`)

func v3110OpenshiftSvcatApiserverSvcYamlBytes() ([]byte, error) {
	return _v3110OpenshiftSvcatApiserverSvcYaml, nil
}

func v3110OpenshiftSvcatApiserverSvcYaml() (*asset, error) {
	bytes, err := v3110OpenshiftSvcatApiserverSvcYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/openshift-svcat-apiserver/svc.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"v3.11.0/openshift-svcat-apiserver/cm.yaml":                                          v3110OpenshiftSvcatApiserverCmYaml,
	"v3.11.0/openshift-svcat-apiserver/cr-aggregate-to-admin.yaml":                       v3110OpenshiftSvcatApiserverCrAggregateToAdminYaml,
	"v3.11.0/openshift-svcat-apiserver/cr-aggregate-to-edit.yaml":                        v3110OpenshiftSvcatApiserverCrAggregateToEditYaml,
	"v3.11.0/openshift-svcat-apiserver/cr-aggregate-to-view.yaml":                        v3110OpenshiftSvcatApiserverCrAggregateToViewYaml,
	"v3.11.0/openshift-svcat-apiserver/cr-namespace-viewer.yaml":                         v3110OpenshiftSvcatApiserverCrNamespaceViewerYaml,
	"v3.11.0/openshift-svcat-apiserver/cr-readiness-probe.yaml":                          v3110OpenshiftSvcatApiserverCrReadinessProbeYaml,
	"v3.11.0/openshift-svcat-apiserver/cr-sar-creator.yaml":                              v3110OpenshiftSvcatApiserverCrSarCreatorYaml,
	"v3.11.0/openshift-svcat-apiserver/cr-serviceclass-viewer.yaml":                      v3110OpenshiftSvcatApiserverCrServiceclassViewerYaml,
	"v3.11.0/openshift-svcat-apiserver/crb-auth-delegator-binding.yaml":                  v3110OpenshiftSvcatApiserverCrbAuthDelegatorBindingYaml,
	"v3.11.0/openshift-svcat-apiserver/crb-namespace-viewer-binding.yaml":                v3110OpenshiftSvcatApiserverCrbNamespaceViewerBindingYaml,
	"v3.11.0/openshift-svcat-apiserver/crb-readiness-binding.yaml":                       v3110OpenshiftSvcatApiserverCrbReadinessBindingYaml,
	"v3.11.0/openshift-svcat-apiserver/crb-sar-creator-binding.yaml":                     v3110OpenshiftSvcatApiserverCrbSarCreatorBindingYaml,
	"v3.11.0/openshift-svcat-apiserver/crb-serviceclass-viewer-binding.yaml":             v3110OpenshiftSvcatApiserverCrbServiceclassViewerBindingYaml,
	"v3.11.0/openshift-svcat-apiserver/ds.yaml":                                          v3110OpenshiftSvcatApiserverDsYaml,
	"v3.11.0/openshift-svcat-apiserver/ns.yaml":                                          v3110OpenshiftSvcatApiserverNsYaml,
	"v3.11.0/openshift-svcat-apiserver/role-extension-apiserver-auth-reader.yaml":        v3110OpenshiftSvcatApiserverRoleExtensionApiserverAuthReaderYaml,
	"v3.11.0/openshift-svcat-apiserver/rolebinding-extension-apiserver-auth-reader.yaml": v3110OpenshiftSvcatApiserverRolebindingExtensionApiserverAuthReaderYaml,
	"v3.11.0/openshift-svcat-apiserver/sa.yaml":                                          v3110OpenshiftSvcatApiserverSaYaml,
	"v3.11.0/openshift-svcat-apiserver/svc.yaml":                                         v3110OpenshiftSvcatApiserverSvcYaml,
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
	"v3.11.0": {nil, map[string]*bintree{
		"openshift-svcat-apiserver": {nil, map[string]*bintree{
			"cm.yaml":                              {v3110OpenshiftSvcatApiserverCmYaml, map[string]*bintree{}},
			"cr-aggregate-to-admin.yaml":           {v3110OpenshiftSvcatApiserverCrAggregateToAdminYaml, map[string]*bintree{}},
			"cr-aggregate-to-edit.yaml":            {v3110OpenshiftSvcatApiserverCrAggregateToEditYaml, map[string]*bintree{}},
			"cr-aggregate-to-view.yaml":            {v3110OpenshiftSvcatApiserverCrAggregateToViewYaml, map[string]*bintree{}},
			"cr-namespace-viewer.yaml":             {v3110OpenshiftSvcatApiserverCrNamespaceViewerYaml, map[string]*bintree{}},
			"cr-readiness-probe.yaml":              {v3110OpenshiftSvcatApiserverCrReadinessProbeYaml, map[string]*bintree{}},
			"cr-sar-creator.yaml":                  {v3110OpenshiftSvcatApiserverCrSarCreatorYaml, map[string]*bintree{}},
			"cr-serviceclass-viewer.yaml":          {v3110OpenshiftSvcatApiserverCrServiceclassViewerYaml, map[string]*bintree{}},
			"crb-auth-delegator-binding.yaml":      {v3110OpenshiftSvcatApiserverCrbAuthDelegatorBindingYaml, map[string]*bintree{}},
			"crb-namespace-viewer-binding.yaml":    {v3110OpenshiftSvcatApiserverCrbNamespaceViewerBindingYaml, map[string]*bintree{}},
			"crb-readiness-binding.yaml":           {v3110OpenshiftSvcatApiserverCrbReadinessBindingYaml, map[string]*bintree{}},
			"crb-sar-creator-binding.yaml":         {v3110OpenshiftSvcatApiserverCrbSarCreatorBindingYaml, map[string]*bintree{}},
			"crb-serviceclass-viewer-binding.yaml": {v3110OpenshiftSvcatApiserverCrbServiceclassViewerBindingYaml, map[string]*bintree{}},
			"ds.yaml": {v3110OpenshiftSvcatApiserverDsYaml, map[string]*bintree{}},
			"ns.yaml": {v3110OpenshiftSvcatApiserverNsYaml, map[string]*bintree{}},
			"role-extension-apiserver-auth-reader.yaml":        {v3110OpenshiftSvcatApiserverRoleExtensionApiserverAuthReaderYaml, map[string]*bintree{}},
			"rolebinding-extension-apiserver-auth-reader.yaml": {v3110OpenshiftSvcatApiserverRolebindingExtensionApiserverAuthReaderYaml, map[string]*bintree{}},
			"sa.yaml":  {v3110OpenshiftSvcatApiserverSaYaml, map[string]*bintree{}},
			"svc.yaml": {v3110OpenshiftSvcatApiserverSvcYaml, map[string]*bintree{}},
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
