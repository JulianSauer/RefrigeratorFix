package main

import (
    "github.com/JulianSauer/RefrigeratorFix/config"
    "github.com/JulianSauer/RefrigeratorFix/philipshue"
)



func main() {
    bridge := philipshue.Login()
    if bridge == nil {
        return
    }

    configFile := config.Load()
    philipshue.UpdateLight(configFile.SmartPlugId, true)
}
