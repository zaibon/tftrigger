# tftrigger
Small tool that trigger the threefold auto-build for a specific repo, branch and commit.

## Usage
```
NAME:
   tftrigger - A new cli application

USAGE:
   tftrigger [global options] command [command options] [path|ghrepo]

VERSION:
   0.0.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --commit value, -c value  commit hash
   --branch value, -b value  branch to use for the build (default: "master")
   --help, -h                show help
   --version, -v             print the version
```

To Manually choose which repository:
```
tftrigger threefoldtech/my_repo
```

To manually choose which repository, branch and commit id:
```
tftrigger--branch master --commit 28275d75474ea6a4639d46e9ac393e664065feda threefoldtech/my_repo
```

But if you are working on a project and want to quickly trigger a build,
you can give a specified the path of a git repository.
Tftrigger will read the information from the `.git` directory and trigger the build for the repository,
the current checkout branch and revision:
```
tftrigger /home/me/a/path/to/a/git/repo
```

Running the `tftrigger` command without a position argument will assume the
current working directory is the path of a git repo from which
trigger the threefold auto-build for a specific branch/commit:
```
tftrigger
```