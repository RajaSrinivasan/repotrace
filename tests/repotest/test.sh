#!/bin/bash
repo init -u https://github.com/RajaSrinivasan/myprojects.git
repo sync
../../../bin/repotrace -f .repo/manifest.xml -L go