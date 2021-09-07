package philipshue

import (
    "github.com/JulianSauer/RefrigeratorFix/config"
    "github.com/amimof/huego"
    "log"
    "os"
    "time"
)

const USER = "RefrigeratorFix"

var cachedBridge *huego.Bridge

func Register() string {
    log.Println("Trying to register")
    bridge := getBridge()

    log.Printf("Bridge found at %s, please press the link button within 1 minute!\n", bridge.Host)

    var currentUser string
    var e error
    for i := 0; i < 12; i++ {
        time.Sleep(5 * time.Second)

        currentUser, e = bridge.CreateUser(USER)
        if e == nil {
            log.Printf("New user %s created\n", currentUser)
            return currentUser
        } else {
            log.Println(e.Error())
        }
    }
    if e != nil {
        log.Fatal("Link button not pressed, aborting")
    }
    return ""
}

func Login() {
    log.Println("Logging in")
    bridge := getBridge()
    userName := getUserName()
    bridge = bridge.Login(userName)
    if bridge != nil {
        log.Println("Logged in")
    }
}

func Update(id int, on bool) {
    log.Printf("Updating %d\n", id)
    bridge := getBridge()
    light, e := bridge.GetLight(id)
    if e != nil {
        log.Fatal(e.Error())
    }
    light.State.On = on
    if _, e := bridge.SetLightState(id, *light.State); e != nil {
        log.Fatal(e.Error())
    }
}

func IsOn(id int) bool {
    log.Printf("Cecking light state for %d\n", id)
    bridge := getBridge()
    light, e := bridge.GetLight(id)
    if e != nil {
        log.Fatal(e.Error())
    }
    return light.State.On
}

func getBridge() *huego.Bridge {
    if cachedBridge == nil {
        bridge, e := huego.Discover()
        if e != nil {
            log.Fatal(e.Error())
        }
        cachedBridge = bridge
    }
    return cachedBridge
}

func getUserName() string {
    configFile := config.Load()
    if configFile.PhilipsHueUserName == "" {
        log.Println("No user found in cache")
        configFile.PhilipsHueUserName = Register()
        if configFile.PhilipsHueUserName == "" {
            os.Exit(1)
        }
        config.Save(configFile)
    } else {
        log.Println("Loaded user from cache")
    }
    return configFile.PhilipsHueUserName
}
