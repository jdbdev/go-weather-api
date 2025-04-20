## Openweathermap API client
An unofficial [openweathermap](https://openweathermap.org/) API client for [Go](https://golang.org/).

[![License](http://img.shields.io/badge/license-MIT-blue.svg)]

## Summary
One Call API 3.0 requires latitude and longitude for API calls. This app Combines two API services from openweathermap.org to allow users to access One Call API 3.0 by using city name instead of latitude and longitude. Requires API key for One Call API 3.0 (free subusciption up to 1000 calls per day). 

Geocoding API: http://api.openweathermap.org/geo/1.0/direct?q={city name},{state code},{country code}&limit={limit}&appid={API key}
One Call API 3.0: https://api.openweathermap.org/data/3.0/onecall?lat={lat}&lon={lon}&exclude={part}&appid={API key}

## Install

```bash
go get github.com/jdbdev/go-weather-api
```
Note: rename .env_sample to .env and add your API key. Requires subscription to One Call API 3.0. 

## License
MIT