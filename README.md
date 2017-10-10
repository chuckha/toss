# toss

## Overview

`toss` is an example of the the done-file semantics used within [sonobuoy][1].
`toss` is not a full-fledged plugin for sonobuoy but simply waits for something to write a donefile
and uploads that file to the configured s3 bucket.

## Getting Started

Take a look at the [sonobuoy example][2]. You can see here two containers running one after the other.
The idea is to be able to do arbitrary actions on the results before they are sent to the master sonobuoy process.
In the following YAML you'll see we write done files which indicate a process has finished running. The YAML
provided is attempting to insert an "upload to s3" step between "run tests" and "report back".

* Create a secret for writing to s3

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: aws-upload-key
  namespace: heptio-sonobuoy
type: Opaque
data:
  access-key-id: *** your access-key-id here ***
  secret-access-key: *** your secret access key here ***
```

* Modify the sonobuoy example yaml to use this container as an intermediate step

```yaml
apiVersion: v1
data:
  e2e.yaml: |
    driver: Job
    name: e2e
    resultType: e2e
    spec:
      containers:

      # The conformance test pod writes to /tmp/sonobuoy/done
      - env:
        - name: E2E_FOCUS
          value: Pods should be submitted and removed
        image: gcr.io/heptio-images/kube-conformance:v1.8.0
        imagePullPolicy: Always
        name: e2e
        volumeMounts:
        - mountPath: /tmp/sonobuoy
          name: test-out

      # The uploader reads from /tmp/sonobuoy/done and writes to /tmp/results/done
      - name: uploader
        image:  gcr.io/heptio-images/toss:v0.0.1
        imagePullPolicy: Always
        volumeMounts:
        - name: uploader-out
          mountPath: /tmp/results
        - name: test-out
          mountPath: /tmp/sonobuoy
        env:
        - name: READ_RESULTS_DIR
          value: /tmp/sonobuoy
        - name: WRITE_RESULTS_DIR
          value: /tmp/results
        - name: ACCESS_KEY_ID
          valueFrom:
            secretKeyRef:
              name: aws-upload-key
              key: access-key-id
        - name: SECRET_ACCESS_KEY
          valueFrom:
            secretKeyRef:
              name: aws-upload-key
              key: secret-access-key
        - name: BUCKET
          value: files.heptio.com
        - name: REGION
          value: us-west-2

      # The worker reads from /tmp/results/done and sends results to the master
      - command:
        - sh
        - -c
        - /sonobuoy worker global -v 5 --logtostderr
        env:
        - name: NODE_NAME
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: spec.nodeName
        - name: RESULTS_DIR
          value: /tmp/results
        image: gcr.io/heptio-images/sonobuoy:master
        imagePullPolicy: Always
        name: sonobuoy-worker
        volumeMounts:
        - mountPath: /etc/sonobuoy
          name: config
        - mountPath: /tmp/results
          name: uploader-out
      restartPolicy: Never
      serviceAccountName: sonobuoy-serviceaccount
      tolerations:
      - effect: NoSchedule
        key: node-role.kubernetes.io/master
        operator: Exists
      - key: CriticalAddonsOnly
        operator: Exists
      volumes:
      - emptyDir: {}
        name: test-out
      - emptyDir: {}
        name: uploader-out
      - configMap:
          name: __SONOBUOY_CONFIGMAP__
        name: config
```

## Contributing

Thanks for taking the time to join our community and start contributing!

#### Before you start

* Please familiarize yourself with the [Code of
Conduct][12] before contributing.
* See [CONTRIBUTING.md][11] for instructions on the
developer certificate of origin that we require.

#### Pull requests

* We welcome pull requests. Feel free to dig through the [issues][10] and jump in.

[1]: https://github.com/heptio/sonobuoy/
[2]: https://github.com/heptio/sonobuoy/blob/master/examples/quickstart.yaml#L113
[10]: https://github.com/heptiolabs/toss/issues
[11]: /CONTRIBUTING.md
[12]: /CODE_OF_CONDUCT.md
