package config

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "os"
)

const CONFIG_NAME = "config.json"

type ConfigFile struct {
    PhilipsHueUserName    string  `json:"philipsHueUserName"`
    SmartPlugId           int     `json:"smartPlugId"`
    MobileAlertsUrl       string  `json:"mobileAlertsUrl"`
    MobileAlertsDeviceIds string  `json:"mobileAlertsDeviceIds"`
    TemperatureMax        float64 `json:"temperatureMax"`
    TemperatureMin        float64 `json:"temperatureMin"`
}

func Load() *ConfigFile {
    file, e := os.Open(CONFIG_NAME)
    if e != nil {
        log.Fatal(e.Error())
    }

    defer file.Close()
    decoder := json.NewDecoder(file)
    config := ConfigFile{}
    if e = decoder.Decode(&config); e != nil {
        log.Fatal("could not parse config.json")
    }
    return &config
}

func Save(config *ConfigFile) {
    if config == nil {
        log.Fatal("Cannot save empty config")
    }
    file, e := os.Open(CONFIG_NAME)
    if e != nil {
        log.Fatal(e.Error())
    }

    defer file.Close()
    if bytes, e := json.Marshal(config); e != nil {
        log.Fatal(e.Error())
    } else {
        e := ioutil.WriteFile(CONFIG_NAME, bytes, 0644)
        if e != nil {
            log.Fatal(e.Error())
        }
    }
}
