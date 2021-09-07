package csvLogger

import (
    "encoding/csv"
    "fmt"
    "github.com/JulianSauer/RefrigeratorFix/csvLogger/dto"
    "log"
    "os"
    "strconv"
)

const LOGNAME = "refrigerator-temperature-log.csv"

func AddToLog(temperatureData dto.TemperatureData) {
    log.Println("Writing result to csv")
    file, e := os.OpenFile(LOGNAME, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
    if e != nil {
        log.Fatal(e.Error())
    }
    
    data := temperatureDataToStrings(temperatureData)
    csvWriter := csv.NewWriter(file)
    if e = csvWriter.Write(data); e != nil {
        file.Close()
        log.Fatal(e.Error())
    }

    csvWriter.Flush()
    file.Close()
}

func temperatureDataToStrings(data dto.TemperatureData) []string {
    return []string{
        data.Date.Format("20060102-150405"),
        fmt.Sprintf("%f", data.Temperature),
        strconv.FormatBool(data.Powered),
    }
}
