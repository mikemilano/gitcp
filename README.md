Gitcp: A Project Utility & Build Tool
======================================

Gitcp is a CLI utility for retrieving parts or entire repositories into your current project. 
It works with both public and private repositories via a short path for Github or a full git URL to any repo.

> This is as much for users fetching files as it is for those maintaining seeds. It could be useful in a CI pipeline as well!

Gitcp not only handles fetching code, but it also takes care of the mapping of code from an external path into the
current project. This makes it not only handy to create project, but to act as a build tool for projects which rely on 
source code in other repositories.

Usage: `gitcp repo-path [src] [dst]`
- `repo-path`: A short github path `owner/project` or a full git clone url.
- `src`: [Glob](https://en.wikipedia.org/wiki/Glob_%28programming%29) style path to files/directories to copy; 
Multiple paths can be separated by commas; Default: Repo root
- `dst`: (optional) Destination path where the source should be copied to; Default behavior places source files in the
same relative paths as they are in the source.

```
OPTIONS:
   --branch value, -b value        Git branch
   --clone-cdir value, -c value    temp directory where repo will be cloned; "memory" to use system memory (default: "/tmp") [/tmp]
   --ssh-key value, -k value       private ssh key to use (default: "/Users/mmilano/.ssh/id_rsa")
   --github-proto value, -p value  Github proto; auto, https, or ssh (default: "auto")
   --preserve-git, -g              preserve .git directory
   --quiet, -q                     quiet mode, no output
   --verbose, -v                   verbose mode for troubleshooting
   --help, -h                      show help
   --version, -V                   print the version
```

## Examples:

Copy the contents of a start state you maintain into your current directory:
```bash
# Copies the contents of Github someuser/my-awesome-starter project into your current directory.
gitcp example/my-awesome-starter
```

Or, maybe you just want to grab that `.editorconfig` file:
```bash
gitcp example/my-favorite-monolith .editorconfig
```

Need to specify a different path to the destination?
```bash
gitcp example/some-web-project .htaccess public/.htaccess
```

Let's use a wildcard (`*`) to grab all components that begin with `User`:
```bash
# If destination is not specified, it will use the same structure for each file.
gitcp example/my-vue-project components/User*
```

Now I need `.htaccess` and `.editorconfig`:
```bash
gitcp example/my-old-web-project .htaccess,.editorconfig
```

... wait, I needed `.htaccess` in the `./public` directory:
```bash
gitcp example/my-old-web-project .htaccess,.editorconfig public/.htaccess,.editorconfig
```

I'd like to copy `.c` files recursively beginning with the `src` directory:
```bash
gitcp example/foo src/**/.c
```

And now only want `.c` files that are named with lower case alphanumeric characters:
```bash
gitcp example/foo src/**/[a-z0-9].c
```

Not using github? Use any git URL:
```bash
gitcp ssh://user@host.xz:port/path/to/repo.git
```

This repo requires a key:
```bash
gitcp --ssh-key ~/.ssh/mykey.pem ssh://user@host.xz:port/path/to/repo.git
```

## More Examples:
```
# Copy the contents of the target repo into your current path
gitcp mikemilano/my-seed

# Copy the `components` directory of the target repo into your current path
gitcp mikemilano/my-seed components

# Copy `components/User.vue` from the target repo into `./mycomponents/MyUser.vue`
#   - The `mycomponents` directory is created if it doesn't exist
gitcp mikemilano/my-vuejs-project components/User.vue mycomponents/MyUser.vue

# Copy multiple remote files
gitcp mikemilano/vuejs-components components/User.vue,components/Admin.vue

# Copy multiple remote paths into multiple different local paths
gitcp mikemilano/vuejs-components components/User.vue,components/Admin.vue components/Users,components/Admin

# Combine the files in multiple remote directories into a single local path
gitcp mikemioano/vuejs-components components/User/*,components/Admin/* components/
```
