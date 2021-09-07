package dtos

type Device struct {
    Deviceid    string `json:"deviceid"`
    Lastseen    int    `json:"lastseen"`
    Lowbattery  bool   `json:"lowbattery"`
    Measurement struct {
        Idx int     `json:"idx"`
        Ts  int     `json:"ts"`
        C   int     `json:"c"`
        Lb  bool    `json:"lb"`
        T1  float64 `json:"t1"`
    } `json:"measurement"`
}
