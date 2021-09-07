package dtos

type DeviceRequest struct {
    Devices []Device `json:"devices"`
    Success bool     `json:"success"`
}
