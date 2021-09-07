# Refrigerator Fix
Regulates a refrigerator by turning a smart plug on or off depending on the temperature of a thermometer. I've used [this](https://mobile-alerts.eu/temperatursensor-ma10100/) and some smart plug from OSRAM which is compatible with Philips hue.

Example docker-compose file:
```yaml
version: '3.6'
services:
  refrigerator-fix:
    image: hcr.io/juliansauer/RefrigeratorFix/refrigerator-fix
    build: .
    restart: always
    volumes:
      - ${PWD}/config.json:/root/config.json
      - ${PWD}/refrigerator-temperature-log.csv:/root/refrigerator-temperature-log.csv

```
Example config.json:
```json
{
  "philipsHueUserName": "1028d66426293e821ecfd9ef1a0731df",
  "smartPlugId": 5,
  "mobileAlertsUrl": "https://www.data199.com/api/pv1/device/lastmeasurement",
  "mobileAlertsDeviceIds": "048FAE73A12C"
}

```
