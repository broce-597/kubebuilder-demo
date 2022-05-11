###创建资源
####1.创建CRD
make install
####2.run test
make run test
#####3.创建资源对象(CR):
kubectl create -f config/samples/elasticweb_v1_elasticweb.yaml

### 清除资源
##### 1 删除资源对象：
kubectl delete -f config/samples/webapp_v1_welcome.yaml.yaml
##### 2 删除controller
kustomize build config/default | kubectl delete -f -
##### 3 删除CRD
make uninstall