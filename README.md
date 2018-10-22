# tftrigger
Small tool that trigger the threefold auto-build for a specific branch/commit

## Usage
```
NAME:
   tftrigger - A new cli application

USAGE:
   tftrigger [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --organization value, -o value  github organization (default: "threefoldtech")
   --repository value, -r value    repository name
   --commit value, -c value        commit hash
   --branch value, -b value        branch to use for the build (default: "master")
   --help, -h                      show help
   --version, -v                   print the version
   ```