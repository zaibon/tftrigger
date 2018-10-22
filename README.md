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
   --path value, -p value          specified the path of the repository to use as source of information for the build
   --help, -h                      show help
   --version, -v                   print the version
```

To manually choose which repository,branch and commit id:
```
tftrigger --organization threefoldtech --repository my_repo --branch master --commit 28275d75474ea6a4639d46e9ac393e664065feda
```

But if you are working on a project and want to quicly trigger a build, you can use the `--path` flag to specified the path of a git repository. Tftrigger will read the information from the `.git` directory and trigger the build for the repository, the current checkout branch and revision

Example: this command will trigger a build for the repository in the current directory
```
tftrigger -p .
```