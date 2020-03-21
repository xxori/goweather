# golang-weather
Simple program to learn golang that gets the current weather from the openweather API

## Dependencies
The only dependancy is the programming language go to build the program into an executable.

You can install golang through your linux distrobutions package manager or from the [site](https://golang.org/dl/)

## Installation
First clone the repo
```
git clone https://github.com/xxori/goweather.git
cd goweather
```
And then build the file into an ELF executable
```
go build goweather.go
```
This will produce an executable you can run with
```
./goweather [options] (city)
```
For use anyway you can move the executable into a directory that is in PATH, such as /usr/bin
```
sudo mv goweather /usr/bin
```

## Usage
Goweather can either output to terminal or in the form on a desktop notification. This is specified through the use of the -n flag.

It can also output different weather information, through -t. The default type is temp. The different types are temp, weather, wind, rain, and snow.

The final argument is the city or place in question.

Example:
```
goweather -n -t weather melbourne
```

