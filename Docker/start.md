# start.sh usage

Usage start.sh [,OPTIONS]

OPTIONS:

| Option | Long Option | Description
|-------:|------------:|:----------------------------------------------------------------------|
|     -h |      --help | Display this help message.                                            |
|     -n |      --node | Deploy node container.                                                |
|     -p |      --peer | Deploy extended peers containers.                                     |
|     -f |      --full | Deploy node and extend peers containers (like -n -p).                 |
|     -s |     --start | Start containers containers.                                          |
|     -S |      --stop | Stop containers.                                                      |
|     -t |       --tag | Tag hyperledger image as latest.                                      |
|     -r |    --remove | Remove containers.                                                    |
|     -a |       --all | Make special actions on all docker containers on your machine.        |
