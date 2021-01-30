# :eagle: kube-falcon

---

[![](https://img.shields.io/badge/kubebuilder-40D1F5?style=flat-square&logo=kubernetes&logoColor=White)](https://book.kubebuilder.io/)

#### 1. install CRDs into the cluster

```bash
$ make install
```

#### 2. run controller

```bash
$ make run
```

#### 3. install CR into the cluster

```bash
$ kubectl apply -f config/samples/
```

#### 4. build and push docker image

```bash
$ make docker-build docker-push IMG=<some-registry>/<project-name>:tag
```

#### 5. deploy the controller into the cluster

```bash
$ make deploy IMG=<some-registry>/<project-name>:tag
```

#### 6. uninstall CRDs

```bash
$ make uninstall
```
