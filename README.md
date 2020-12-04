# Potentially Hazardous Asteroids

This API service serves info about potentially hazardous asteroids that come close to Earth in the next 7 days starting from the current date.

The data is collected from NASA's Near Earth Object Web Service.

[![codecov](https://codecov.io/gh/psyb0t/potentially-hazardous-asteroids/branch/master/graph/badge.svg?token=4LANEUHEQX)](https://codecov.io/gh/psyb0t/potentially-hazardous-asteroids)
![build](https://github.com/psyb0t/potentially-hazardous-asteroids/workflows/build/badge.svg)

## Supported Environment Variables

PHA_CONFIG_FPATH - configuration file path. if set, the config struct will have its fields loaded from the given yaml file

PHA_LISTEN_ADDRESS - the listen address for the http server. if PHA_CONFIG_FPATH is set but listen address is not specified in the file, this will be used. if this env var is also missing, the default value of 127.0.0.1:8080 will be used.

PHA_NASA_API_KEY - the API key used for performing calls to the NeoWs API. if PHA_CONFIG_FPATH is set but nasa api key is not specified in the file, this will be used. if this env var is also missing, a fatal error will be thrown.

PHA_DEBUG - can be set to anything, it just needs to be set. informs the running program that it's in debug mode, meaning that it will output debug level messages. supersedes the value from the config file, if present.

## Run via Docker

```
% docker run --rm --env PHA_DEBUG="" --env PHA_LISTEN_ADDRESS="0.0.0.0:8005" --env PHA_NASA_API_KEY="API-KEY" -p 127.0.0.1:8005:8005 psyb0t/potentially-hazardous-asteroids
```

## TODO

log http request data

use multiwriter to log to stdout + log file

write tests

make this readme prettier
