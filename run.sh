#!/bin/bash

go build -o bookings cmd/web/*.go && ./bookings
./bookings -dbname=bookings -dbuser=grandmastahash -cache=false -production=false