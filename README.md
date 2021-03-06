[![Go Report Card](https://goreportcard.com/badge/github.com/MarkusAJacobsen/PGC)](https://goreportcard.com/report/github.com/MarkusAJacobsen/PGC)

# Pocket Gardening Core

# Building 
## Setting up Neo4j
* Make sure you have a C/C++ toolchain
* Install Pkg-config `apt install pkg-config`
* Seabolt:
   * Clone Seabolt `git clone https://github.com/neo4j-drivers/seabolt.git`
   * Install Seabolt dependencies `apt install git build-essential cmake libssl-dev`
   * Run `$SEABOLT_DIR`$ `./make_debug.sh` 
   * Set `PKG_CONFIG_PATH` to `seabolt/build/dist/share/pkgconfig` on Linux

# Running
## Locally 
* Install docker
* Run `./runDocker.sh` on Linux and `sh runDocker.sh` on Win to start project containers

### Building PGC image separately
* `src-root$ docker run PORT=5555 -t pgc_web .`

### Changing logger PGL URL
PGC uses [PGL](https://github.com/MarkusAJacobsen/PGL) to log connector and application logs. To set PGL's URL, set environment variable `LOG_URL` to proper URL
 
   
