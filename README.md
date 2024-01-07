# micro

# ETCD
nohup etcd --listen-peer-urls='http://192.168.3.117:2379' > /dev/null &


docker run -d \
  -p 2379:2379 \
  --name etcd \
  quay.io/coreos/etcd:v3.5.0 \
  etcd \
  --advertise-client-urls http://0.0.0.0:2379 \
  --listen-client-urls http://0.0.0.0:2379 \
  --initial-cluster-state new


sudo sysctl -w vm.max_map_count=262144

docker run -it --name search -v /c/Users/docker_config/search:/go/src/config/   xbinwar3/search:v1
