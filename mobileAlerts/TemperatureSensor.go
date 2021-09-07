package mobileAlerts

import (
    "encoding/json"
    "github.com/JulianSauer/RefrigeratorFix/config"
    "github.com/JulianSauer/RefrigeratorFix/mobileAlerts/dtos"
    "gopkg.in/resty.v1"
    "log"
)

func GetTemperature(deviceId string) float64 {
    log.Println("Checking temperature")
    configFile := config.Load()
    url := configFile.MobileAlertsUrl
    parameters := "deviceids=" + deviceId

    client := resty.New()
    response, e := client.R().
        SetHeader("Content-Type", "application/x-www-form-urlencoded").
        SetBody(parameters).
        Post(url)
    if e != nil {
        log.Fatal(e.Error())
    }
    var deviceRequest dtos.DeviceRequest
    if e = json.Unmarshal(response.Body(), &deviceRequest); e != nil {
        log.Fatal(e.Error())
    }
    if len(deviceRequest.Devices) != 1 {
        log.Fatalf("Found %d devices, expected 1", len(deviceRequest.Devices))
    }
    return deviceRequest.Devices[0].Measurement.T1
}
