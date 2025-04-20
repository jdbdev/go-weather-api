## Openweathermap API client
An unofficial [openweathermap](https://openweathermap.org/) API client for [Go](https://golang.org/).

[![License](http://img.shields.io/badge/license-MIT-blue.svg)] 

## Summary
Combines two API services from openweathermap.org to allow users to access One Call API 3.0 by using city name instead of latitude and longitude.
Requires API key for One Call API 3.0 (free subusciption up to 1000 calls per day). 

## Install

```bash
go get -u github.com/jdbdev/go-weather.api
```
Note: rename .env_sample to .env and add your API key. Requires subscription to One Call API 3.0. 