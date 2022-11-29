[![Go Report](https://goreportcard.com/badge/github.com/rekram1-node/NetworkMonitor)](https://goreportcard.com/report/github.com/rekram1-node/NetworkMonitor)

# NetworkMonitor
IOT program that works on any device that can monitor your network, letting you know when you get outages as well as auto-updating itself making it easy to setup

## Installation
#### Docker
[Docker Hub repo](https://hub.docker.com/repository/docker/rekram/network-monitor)
```shell
docker run rekram/network-monitor:latest
```

#### Binary Install
visit our [releases page](https://github.com/rekram1-node/NetworkMonitor/releases)

## Configuration
#### Initialization
Add the "-init" flag when running the application for the first time
this will create a directory called "network-monitor" at root where your outages and configuration will be placed
follow prompts and add correct stuff as needed

#### Uploads
The "config.yaml" file will contain a parameter called "uploadscripts" add your script and it will be called when you run the app with the "-upload" flag

## Usage
All settings for this will be configured on initialization and can be edited in the configuration file during run.
The app will initialize a new directory: "network-monitor" at root for your OS, and add the file: "config.yaml"
While initializing you will be walked through a series of prompts to help you have a properly configured app
