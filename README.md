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
* Run `docker-compose up --build` to start project containers
 
   