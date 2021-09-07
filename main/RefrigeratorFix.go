package main

import (
    "github.com/JulianSauer/RefrigeratorFix/config"
    "github.com/JulianSauer/RefrigeratorFix/csvLogger"
    "github.com/JulianSauer/RefrigeratorFix/csvLogger/dto"
    "github.com/JulianSauer/RefrigeratorFix/mobileAlerts"
    "github.com/JulianSauer/RefrigeratorFix/philipshue"
    "log"
    "time"
)

const TEMPERATURE_MAX float64 = 10.0
const TEMPERATURE_MIN float64 = 5.0

func main() {
    philipshue.Login()
    configFile := config.Load()
    fridgeIsPowered := philipshue.IsOn(configFile.SmartPlugId)
    currentTemperature := mobileAlerts.GetTemperature(configFile.MobileAlertsDeviceIds)

    if currentTemperature < TEMPERATURE_MIN && fridgeIsPowered {
        log.Printf("Fridge is too cold (%f°C), turning power off\n", currentTemperature)
        philipshue.Update(configFile.SmartPlugId, false)
    } else if currentTemperature > TEMPERATURE_MAX && !fridgeIsPowered {
        log.Printf("Fridge is too warm (%f°C), turning power on\n", currentTemperature)
        philipshue.Update(configFile.SmartPlugId, true)
    } else {
        log.Printf("Temperature is within range at %f°C\n", currentTemperature)
    }

    location, _ := time.LoadLocation("Europe/Berlin")
    now := time.Now().In(location)
    temperatureData := dto.TemperatureData{
        Date:        now,
        Temperature: currentTemperature,
        Powered:     fridgeIsPowered,
    }
    csvLogger.AddToLog(temperatureData)
}
