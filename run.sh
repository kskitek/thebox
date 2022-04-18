#!/bin/bash

env $(cat /home/pi/.env) /home/pi/box 2>&1 /var/log/box.log
