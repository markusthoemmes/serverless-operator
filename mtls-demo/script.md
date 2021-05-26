# ServiceMesh PoC

## Install ServiceMesh and Serverless operators

This installs Serverless from images pushed for this branch right now. Serverless 1.15 will support this in an actual release. If you need to mirror the images, all image references in [the CSV](../olm-catalog/serverless-operator/manifests/serverless-operator.clusterserviceversion.yaml) will have to be mirrored accordingly, prior to installation. In the same way, `quay.io/openshift-knative/serverless-index:v1.14.3` will have to be available (it's part of [`catalogsource.bash`](../hack/lib/catalogsource.bash))

```
make install-mesh install-operator
```

## Create a wildcard certificate to use for the respective domain

This assumes a CRC based app but can in theory be used with other wildcard certs as well.

Create a root certificate:
```
openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 \
    -subj '/O=Example Inc./CN=apps-crc.testing' \
    -keyout root.key \
    -out root.crt
```

Create a wildcard certificate:
```
openssl req -nodes -newkey rsa:2048 \
    -subj "/CN=*.apps-crc.testing/O=Example Inc." \
    -keyout wildcard.key \
    -out wildcard.csr
```

Sign the wildcard certificate:
```
openssl x509 -req -days 365 -set_serial 0 \
    -CA root.crt \
    -CAkey root.key \
    -in wildcard.csr \
    -out wildcard.crt
```

Create a secret with the wildcard certificate:
```
oc create -n istio-system secret tls wildcard-certs \
    --key=wildcard.key \
    --cert=wildcard.crt
```

## Deploy and configure Service Mesh Control Plane

These steps are not necessarily related to Serverless, but are a canonical (and tested) way to setup an SMCP. If you use the `istio-system` namespace, Serverless will just work with it's default settings.

Create the `istio-system` namespace:

```
oc create ns istio-system
```

Deploy following `ServiceMeshControlPlane` for Istio. We set `spec.accessLogging` to verify if the traffic uses TLS or not, but it is not mandatory for production workloads. We also mount the certificates created above into the ingressgateway.

```yaml
cat <<EOF | oc apply -f -
apiVersion: maistra.io/v2
kind: ServiceMeshControlPlane
metadata:
  name: basic
  namespace: istio-system
spec:
  profiles:
  - default
  proxy:
    accessLogging:
      file:
        name: /dev/stdout
        format: "{ \"authority\": \"%REQ(:AUTHORITY)%\", \"bytes_received\": %BYTES_RECEIVED%, \"bytes_sent\": %BYTES_SENT%, \"downstream_local_address\": \"%DOWNSTREAM_LOCAL_ADDRESS%\", \"downstream_peer_cert_v_end\": \"%DOWNSTREAM_PEER_CERT_V_END%\", \"downstream_peer_cert_v_start\": \"%DOWNSTREAM_PEER_CERT_V_START%\", \"downstream_remote_address\": \"%DOWNSTREAM_REMOTE_ADDRESS%\", \"downstream_tls_cipher\": \"%DOWNSTREAM_TLS_CIPHER%\", \"downstream_tls_version\": \"%DOWNSTREAM_TLS_VERSION%\", \"duration\": %DURATION%, \"hostname\": \"%HOSTNAME%\", \"istio_policy_status\": \"%DYNAMIC_METADATA(istio.mixer:status)%\", \"method\": \"%REQ(:METHOD)%\", \"path\": \"%REQ(X-ENVOY-ORIGINAL-PATH?:PATH)%\", \"protocol\": \"%PROTOCOL%\", \"request_duration\": %REQUEST_DURATION%, \"request_id\": \"%REQ(X-REQUEST-ID)%\", \"requested_server_name\": \"%REQUESTED_SERVER_NAME%\", \"response_code\": \"%RESPONSE_CODE%\", \"response_duration\": %RESPONSE_DURATION%, \"response_tx_duration\": %RESPONSE_TX_DURATION%, \"response_flags\": \"%RESPONSE_FLAGS%\", \"route_name\": \"%ROUTE_NAME%\", \"start_time\": \"%START_TIME%\", \"upstream_cluster\": \"%UPSTREAM_CLUSTER%\", \"upstream_host\": \"%UPSTREAM_HOST%\", \"upstream_local_address\": \"%UPSTREAM_LOCAL_ADDRESS%\", \"upstream_service_time\": %RESP(X-ENVOY-UPSTREAM-SERVICE-TIME)%, \"upstream_transport_failure_reason\": \"%UPSTREAM_TRANSPORT_FAILURE_REASON%\", \"user_agent\": \"%REQ(USER-AGENT)%\", \"x_forwarded_for\": \"%REQ(X-FORWARDED-FOR)%\" } \n"
  version: v2.0
EOF
```

### Replace istiod image (to workaround OSSM-449)

This will no longer be necessary with an upcoming ServiceMesh release.

```
oc -n istio-system set image deploy/istiod-basic discovery=docker.io/markusthoemmes/pilot-ubi8:2.0.2-fix
```

### Enable mTLS strict mode

__IMPORTANT:__ The PeerAuthentication object needs to be modified. The official docs have a bug https://github.com/openshift/openshift-docs/issues/28869.

```yaml
cat <<EOF | oc apply -f -
apiVersion: security.istio.io/v1beta1
kind: PeerAuthentication
metadata:
  name: default
  namespace: istio-system
spec:
  mtls:
    mode: STRICT
EOF
```

## Create the necessary gateways

These are the gateways that Serverless will use to communicate with ServiceMesh. Note how they reference the wildcard certs created above.

```yaml
cat <<EOF | oc apply -f -
apiVersion: networking.istio.io/v1alpha3
kind: Gateway
metadata:
  name: knative-ingress-gateway
  namespace: knative-serving
spec:
  selector:
    istio: ingressgateway
  servers:
    - port:
        number: 443
        name: https
        protocol: HTTPS
      hosts:
        - "*"
      tls:
        mode: SIMPLE
        credentialName: wildcard-certs
---
apiVersion: networking.istio.io/v1alpha3
kind: Gateway
metadata:
  name: knative-local-gateway
  namespace: knative-serving
  labels:
    serving.knative.dev/release: devel
    networking.knative.dev/ingress-provider: istio
spec:
  selector:
    istio: ingressgateway
  servers:
    - port:
        number: 8081
        name: https
        protocol: HTTPS
      hosts:
        - "*"
      tls:
        mode: SIMPLE
        credentialName: wildcard-certs
---
apiVersion: v1
kind: Service
metadata:
  name: knative-local-gateway
  namespace: istio-system
  labels:
    serving.knative.dev/release: devel
    networking.knative.dev/ingress-provider: istio
spec:
  type: ClusterIP
  selector:
    istio: ingressgateway
  ports:
    - name: http2
      port: 80
      targetPort: 8081
EOF
```

## Add knative-serving and app's namespaces into SMMR

We include `knative-serving` namespace into Mesh as we enable sidecar injection into the system pods. `default` will be the namespace we will deploy the applications in. Add more namespaces as you see fit here.

```yaml
cat <<EOF | oc apply -f -
apiVersion: maistra.io/v1
kind: ServiceMeshMemberRoll
metadata:
  name: default
  namespace: istio-system
spec:
  members:
  - knative-serving
  - default
EOF
```

## Install Knative Serving with Istio support and injection

This installs Knative Serving including the istio support, which will connect it to ServiceMesh. The default ingress class is set to point to istio as well, so all Services are automatically routed via the istio-ingressgateway. All relevant data-plane components get sidecars injected as well.

```yaml
cat <<EOF | oc apply -f -
apiVersion: operator.knative.dev/v1alpha1
kind: KnativeServing
metadata:
  name: knative-serving
  namespace: knative-serving
spec:
  high-availability:
    replicas: 1
  config:
    network:
      "ingress.class": "istio.ingress.networking.knative.dev"
  ingress:
    istio:
      enabled: true
  deployments:
  - name: activator
    annotations:
      "sidecar.istio.io/inject": "true"
      "sidecar.istio.io/rewriteAppHTTPProbers": "true"
  - name: autoscaler
    annotations:
      "sidecar.istio.io/inject": "true"
      "sidecar.istio.io/rewriteAppHTTPProbers": "true"
EOF
```

## Deploy KService with injection

Finally, we can deploy the Knative Service. Note that each Service will have to have the injection annotations added to it for Service Mesh to actually inject the sidecar.

```yaml
cat <<EOF | oc apply -f -
apiVersion: serving.knative.dev/v1
kind: Service
metadata:
  name: hello
  annotations:
    serving.knative.openshift.io/enablePassthrough: "true"
spec:
  template:
    metadata:
      annotations:
        sidecar.istio.io/inject: "true"
        sidecar.istio.io/rewriteAppHTTPProbers: "true"
    spec:
      containers:
      - image: docker.io/openshift/hello-openshift
EOF
```

The Service is routable with a passthrough route out-of-the-box.

```console
$ oc get ksvc
NAME    URL                                     LATESTCREATED   LATESTREADY   READY   REASON
hello   http://hello-default.apps-crc.testing   hello-00001     hello-00001   True
```

**Note:** It showing HTTP is a known bug here. HTTP isn't actually supported for this Service.

And now, we can hit it with the certificate we've created before:

```console
$ curl --cacert wildcard.crt https://hello-default.apps-crc.testing
Hello OpenShift!
```

## Deploy 2nd KService using the first service

Now, we can deploy a second KService that'll proxy every request to the first one to showcase service-to-service communication.

```yaml
cat <<EOF | oc apply -f -
apiVersion: serving.knative.dev/v1
kind: Service
metadata:
  name: hello-proxy
  annotations:
    serving.knative.openshift.io/enablePassthrough: "true"
spec:
  template:
    metadata:
      annotations:
        sidecar.istio.io/inject: "true"
        sidecar.istio.io/rewriteAppHTTPProbers: "true"
    spec:
      containers:
      - image: quay.io/openshift-knative/httpproxy:v0.21
        env:
        - name: TARGET_HOST
          value: hello.default.svc.cluster.local
EOF
```

```console
$ oc get ksvc hello-proxy
NAME          URL                                           LATESTCREATED       LATESTREADY         READY   REASON
hello-proxy   http://hello-proxy-default.apps-crc.testing   hello-proxy-00001   hello-proxy-00001   True
```

We can likewise hit it with the cert created before.

```console
$ curl --cacert wildcard.crt https://hello-proxy-default.apps-crc.testing
Hello OpenShift!
```

## Verify that traffic is actually encrypted

Since we installed ServiceMesh with a special accesLog setting that logs the used encryption cyphers etc., we can now verify that the data-path is encrypted.

### outside -> istio-ingressgateway

```console
$ oc -n istio-system logs istio-ingressgateway-7658b99f97-dljg7
{ "authority": "hello-default.apps-crc.testing", "bytes_received": 0, "bytes_sent": 17, "downstream_local_address": "10.217.0.127:8443", "downstream_peer_cert_v_end": "-", "downstream_peer_cert_v_start": "-", "downstream_remote_address": "10.217.0.1:40110", "downstream_tls_cipher": "TLS_AES_256_GCM_SHA384", "downstream_tls_version": "TLSv1.3", "duration": 8312, "hostname": "istio-ingressgateway-7658b99f97-dljg7", "istio_policy_status": "-", "method": "GET", "path": "/", "protocol": "HTTP/2", "request_duration": 0, "request_id": "5617b439-d383-4075-8b4c-6691cc706e99", "requested_server_name": "hello-default.apps-crc.testing", "response_code": "200", "response_duration": 8312, "response_tx_duration": 0, "response_flags": "-", "route_name": "-", "start_time": "2021-04-29T13:21:20.115Z", "upstream_cluster": "outbound|80||hello-00001.default.svc.cluster.local", "upstream_host": "10.217.0.96:8012", "upstream_local_address": "10.217.0.127:59540", "upstream_service_time": 8312, "upstream_transport_failure_reason": "-", "user_agent": "curl/7.76.1", "x_forwarded_for": "10.217.0.1" } 
```

### istio-ingressgateway -> activator

```console
$ oc -n knative-serving logs deploy/activator -c istio-proxy | grep "inbound|80" | grep '"path": "/"'
{ "authority": "hello-default.apps-crc.testing", "bytes_received": 0, "bytes_sent": 17, "downstream_local_address": "10.217.0.96:8012", "downstream_peer_cert_v_end": "2021-04-30T12:17:32.000Z", "downstream_peer_cert_v_start": "2021-04-29T12:17:32.000Z", "downstream_remote_address": "10.217.0.1:0", "downstream_tls_cipher": "ECDHE-RSA-AES256-GCM-SHA384", "downstream_tls_version": "TLSv1.2", "duration": 7978, "hostname": "activator-7c574559f8-mxbvw", "istio_policy_status": "-", "method": "GET", "path": "/", "protocol": "HTTP/1.1", "request_duration": 0, "request_id": "3859faef-518d-4f5c-b816-447d1a9a57b2", "requested_server_name": "outbound_.80_._.hello-00001.default.svc.cluster.local", "response_code": "200", "response_duration": 7978, "response_tx_duration": 0, "response_flags": "-", "route_name": "default", "start_time": "2021-04-29T12:29:24.046Z", "upstream_cluster": "inbound|80|http|activator-service.knative-serving.svc.cluster.local", "upstream_host": "127.0.0.1:8012", "upstream_local_address": "127.0.0.1:60334", "upstream_service_time": 7978, "upstream_transport_failure_reason": "-", "user_agent": "curl/7.76.1", "x_forwarded_for": "192.168.130.1,10.217.0.1" }
```

### activator -> application

```console
$ oc logs hello-00001-deployment-79b949bbdb-r9xgv -c istio-proxy | grep "inbound|80" | grep '"path": "/"'
{ "authority": "10.217.4.234:80", "bytes_received": 0, "bytes_sent": 17, "downstream_local_address": "10.217.0.113:8012", "downstream_peer_cert_v_end": "2021-04-30T12:25:28.000Z", "downstream_peer_cert_v_start": "2021-04-29T12:25:28.000Z", "downstream_remote_address": "127.0.0.1:0", "downstream_tls_cipher": "ECDHE-RSA-AES256-GCM-SHA384", "downstream_tls_version": "TLSv1.2", "duration": 2, "hostname": "hello-00001-deployment-79b949bbdb-r9xgv", "istio_policy_status": "-", "method": "GET", "path": "/", "protocol": "HTTP/1.1", "request_duration": 0, "request_id": "c915f735-2c4a-42e8-bb27-8eee50781889", "requested_server_name": "outbound_.80_._.hello-00001-private.default.svc.cluster.local", "response_code": "200", "response_duration": 2, "response_tx_duration": 0, "response_flags": "-", "route_name": "default", "start_time": "2021-04-29T12:37:44.424Z", "upstream_cluster": "inbound|80|http|hello-00001-private.default.svc.cluster.local", "upstream_host": "127.0.0.1:8012", "upstream_local_address": "127.0.0.1:35870", "upstream_service_time": 2, "upstream_transport_failure_reason": "-", "user_agent": "curl/7.76.1", "x_forwarded_for": "192.168.130.1,10.217.0.1, 127.0.0.1" }
```