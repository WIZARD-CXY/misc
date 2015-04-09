package main

import "github.com/socketplane/socketplane/Godeps/_workspace/src/github.com/samalba/dockerclient"
import "fmt"
import "encoding/json"

func main() {

	jsonBody := dockerclient.ContainerConfig{}

	//requestBody := `{"Binds":["/var/lib/kubelet/pods/ad61ca05-cf75-11e4-928c-f8bc12984e72/containers/master/ad39d1b5340050f16abdebe8e949cb3713bcf9775494b68455e62f65d2970e65:/dev/termination-log"],"NetworkMode":"container:17d6b9a3344a31fc0a63d0c610d0ee016ac720149ae4178c0c98151a63b91ade","IpcMode":"container:17d6b9a3344a31fc0a63d0c610d0ee016ac720149ae4178c0c98151a63b91ade","RestartPolicy":{}}`
	requestBody := `{"Hostname":"","Domainname":"","User":"","Memory":0,"MemorySwap":0,"CpuShares":0,"Cpuset":"","AttachStdin":false,"AttachStdout":false,"AttachStderr":false,"PortSpecs":null,"ExposedPorts":{},"Tty":true,"OpenStdin":true,"StdinOnce":false,"Env":[],"Cmd":null,"Image":"10.10.103.215:5000/sshd","Volumes":{},"WorkingDir":"","Entrypoint":null,"NetworkDisabled":false,"MacAddress":"","OnBuild":null,"HostConfig":{"Binds":null,"ContainerIDFile":"","LxcConf":[],"Privileged":false,"PortBindings":{},"Links":null,"PublishAllPorts":false,"Dns":null,"DnsSearch":null,"ExtraHosts":null,"VolumesFrom":null,"Devices":[],"NetworkMode":"bridge","IpcMode":"","PidMode":"","CapAdd":null,"CapDrop":null,"RestartPolicy":{"Name":"","MaximumRetryCount":0},"SecurityOpt":null,"ReadonlyRootfs":false}}`

	if err := json.Unmarshal([]byte(requestBody), &(jsonBody)); err != nil {
		fmt.Println("json unmarshal failed", err)
		return
	}

	fmt.Printf("haha get request body json %+v\n", jsonBody.HostConfig)

}
