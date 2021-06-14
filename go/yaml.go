package main

import (
    "io/ioutil"
    "log"
    yaml "gopkg.in/yaml.v2"
)
type Yaml struct {
    Mysql struct {
        User string `yaml:"user"`
        Host string `yaml:"host"`
        Password string `yaml:"password"`
        Port string `yaml:"port"`
        Name string `yaml:"name"`
    }
    Cache struct {
        Enable bool `yaml:"enable"`
        List []string `yaml:"list,flow"`
    }
}

// Yaml1 struct of yaml
type Yaml1 struct {
    SQLConf Mysql `yaml:"mysql"`
    CacheConf Cache `yaml:"cache"`
}

// Yaml2 struct of yaml
type Yaml2 struct {
    Mysql `yaml:"mysql,inline"`
    Cache `yaml:"cache,inline"`
}

type Mysql struct {
    User string `yaml:"user"`
    Host string `yaml:"host"`
    Password string `yaml:"password"`
    Port string `yaml:"port"`
    Name string `yaml:"name"`
}

// Cache struct of cache conf
type Cache struct {
    Enable bool `yaml:"enable"`
    List []string `yaml:"list,flow"`
}

func main() {
    // resultMap := make(map[string]interface{})
    conf := new(Yaml)
    yamlFile, err := ioutil.ReadFile("test.yaml")
   

    log.Println("yamlFile:", yamlFile)
    if err != nil {
        log.Printf("yamlFile.Get err #%v ", err)
    }
    err = yaml.Unmarshal(yamlFile, conf)
    // err = yaml.Unmarshal(yamlFile, &resultMap)
    if err != nil {
        log.Fatalf("Unmarshal: %v", err)
    }
    log.Printf("conf %#v", conf)
    // log.Println("conf", resultMap)
}
