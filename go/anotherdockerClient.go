package main

import "github.com/socketplane/socketplane/Godeps/_workspace/src/github.com/samalba/dockerclient"
import "fmt"
import "encoding/json"

func main() {

	jsonBody := dockerclient.ContainerConfig{}

	requestBody := `{"Binds":["/var/lib/kubelet/pods/ad61ca05-cf75-11e4-928c-f8bc12984e72/containers/master/ad39d1b5340050f16abdebe8e949cb3713bcf9775494b68455e62f65d2970e65:/dev/termination-log"],"NetworkMode":"container:17d6b9a3344a31fc0a63d0c610d0ee016ac720149ae4178c0c98151a63b91ade","IpcMode":"container:17d6b9a3344a31fc0a63d0c610d0ee016ac720149ae4178c0c98151a63b91ade","RestartPolicy":{}}`

	if err := json.Unmarshal([]byte(requestBody), &(jsonBody.HostConfig)); err != nil {
		fmt.Println("json unmarshal failed")
		return
	}

	fmt.Printf("haha get request body json %+v\n", jsonBody.HostConfig)

}
