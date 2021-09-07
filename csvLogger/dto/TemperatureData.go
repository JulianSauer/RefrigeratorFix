package dto

import "time"

type TemperatureData struct {
    Date        time.Time
    Temperature float64
    Powered     bool
}
