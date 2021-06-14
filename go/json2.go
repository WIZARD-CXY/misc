package main

import (
_	"encoding/json"
	"log"
        "strconv"
_	"os"
)
var s string =`{\n  \"spec\": {\n    \"autoReplace\": true,\n    \"pod\": {\n      \"affinity\": {\n        \"podAntiAffinity\": {\n          \"requiredDuringSchedulingIgnoredDuringExecution\": [\n            {\n              \"labelSelector\": {\n                \"matchExpressions\": [\n                  {\n                    \"key\": \"etcd_cluster\",\n                    \"operator\": \"In\",\n                    \"values\": [\n                      \"test-etcd5\"\n                    ]\n                  }\n                ]\n              },\n              \"topologyKey\": \"kubernetes.io/hostname\"\n            }\n          ],\n          \"preferredDuringSchedulingIgnoredDuringExecution\": [\n            {\n              \"weight\": 100,\n              \"podAffinityTerm\": {\n                \"labelSelector\": {\n                  \"matchExpressions\": [\n                    {\n                      \"key\": \"etcd_cluster\",\n                      \"operator\": \"In\",\n                      \"values\": [\n                        \"test-etcd5\"\n                      ]\n                    }\n                  ]\n                },\n                \"topologyKey\": \"failure-domain.beta.kubernetes.io/zone\"\n              }\n            }\n          ]\n        }\n      },\n      \"etcdEnv\": [\n        {\n          \"name\": \"ETCD_EXPERIMENTAL_BACKEND_BBOLT_FREELIST_TYPE\",\n          \"value\": \"map\"\n        },\n        {\n          \"name\": \"ETCD_AUTO_COMPACTION_RETENTION\",\n          \"value\": \"5m\"\n        },\n        {\n          \"name\": \"ETCD_QUOTA_BACKEND_BYTES\",\n          \"value\": \"100000000000\"\n        },\n        {\n          \"name\": \"ETCD_BACKEND_BATCH_LIMIT\",\n          \"value\": \"1000\"\n        },\n        {\n          \"name\": \"ETCD_BACKEND_BATCH_INTERVAL\",\n          \"value\": \"10ms\"\n        }\n      ],\n      \"labels\": {\n        \"workload\": \"etcd\"\n      },\n      \"nodeSelector\": {\n        \"alibabacloud.com/nodepool-name\": \"ack-pro\",\n        \"workload\": \"etcd\"\n      },\n      \"persistentVolumeClaimSpec\": {\n        \"accessModes\": [\n          \"ReadWriteOnce\"\n        ],\n        \"dataSource\": null,\n        \"resources\": {\n          \"requests\": {\n            \"storage\": \"100Gi\"\n          }\n        },\n        \"storageClassName\": \"alicloud-etcd-pro-disk-essd\"\n      },\n      \"resources\": {\n        \"requests\": {\n          \"memory\": \"500Mi\",\n          \"cpu\": \"100m\"\n        },\n        \"limits\": {\n          \"memory\": \"6Gi\",\n          \"cpu\": \"2\"\n        }\n      },\n      \"tolerations\": [\n        {\n          \"effect\": \"NoSchedule\",\n          \"key\": \"alibabacloud.com/nodepool-name\",\n          \"operator\": \"Equal\",\n          \"value\": \"ack-pro\"\n        }\n      ]\n    },\n    \"repository\": \"registry.cn-hangzhou.aliyuncs.com/acs/etcd\",\n    \"size\": 3,\n    \"version\": \"2019rc4\"\n  }\n}\n`
func main() {
   s, _ := strconv.Unquote(s)
   log.Println(s)
}
