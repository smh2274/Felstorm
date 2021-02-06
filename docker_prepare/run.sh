#!/bin/sh

/Azeroth/Felstorm/felstorm &

/usr/local/bin/envoy -c /Azeroth/Felstorm/config/envoy.yaml -l debug


