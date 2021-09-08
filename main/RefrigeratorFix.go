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

func main() {
    philipshue.Login()
    configFile := config.Load()
    currentTemperature := mobileAlerts.GetTemperature(configFile.MobileAlertsDeviceIds)

    if currentTemperature < configFile.TemperatureMin {
        log.Printf("Fridge is too cold (%f°C)\n", currentTemperature)
        philipshue.Update(configFile.SmartPlugId, false)
    } else if currentTemperature > configFile.TemperatureMax {
        log.Printf("Fridge is too warm (%f°C)\n", currentTemperature)
        philipshue.Update(configFile.SmartPlugId, true)
    } else {
        log.Printf("Temperature is within range at %f°C\n", currentTemperature)
    }
    fridgeIsPowered := philipshue.IsOn(configFile.SmartPlugId)

    location, _ := time.LoadLocation("Europe/Berlin")
    now := time.Now().In(location)
    temperatureData := dto.TemperatureData{
        Date:        now,
        Temperature: currentTemperature,
        Powered:     fridgeIsPowered,
    }
    csvLogger.AddToLog(temperatureData)
}
