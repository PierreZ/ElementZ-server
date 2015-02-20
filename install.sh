#!/bin/bash

set -e;

docker build -t elementz_image .;

docker run -d -p 3000:3000 --name elementz elementz_image
