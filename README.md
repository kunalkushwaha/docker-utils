### docker-utils - a handy tool for extra docker functionality


```
$ docker-utils
Some missing docker functions

Usage:
  docker-utils [command]

Available Commands:
  rm          deletes docker containers in bulk
  rmi         deletes the docker images

Flags:
  -d, --dryrun   dryrun command
  -h, --help     help for docker-utils

Use "docker-utils [command] --help" for more information about a command.

```

#### Removing `Exited` containers

To delete all containers with `Exited` status

`` docker-utils rm``

#### Removing `Untagged` images

To delete all containers with `Exited` status

`` docker-utils rmi``

NOTE: Use `-d` flag with commands for dry-run
