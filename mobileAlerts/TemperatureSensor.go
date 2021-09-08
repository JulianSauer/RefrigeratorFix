package mobileAlerts

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/JulianSauer/RefrigeratorFix/config"
    "github.com/JulianSauer/RefrigeratorFix/mobileAlerts/dtos"
    "gopkg.in/resty.v1"
    "log"
)

func GetTemperature(deviceId string) (float64, error) {
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
        return 0, e
    }
    var deviceRequest dtos.DeviceRequest
    if e = json.Unmarshal(response.Body(), &deviceRequest); e != nil {
        return 0, e
    }
    if len(deviceRequest.Devices) != 1 {
        return 0, errors.New(fmt.Sprintf("Found %d devices, expected 1", len(deviceRequest.Devices)))
    }
    return deviceRequest.Devices[0].Measurement.T1, nil
}
