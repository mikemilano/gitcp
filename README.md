Seeder: A Project Utility & Build Tool
======================================

Instead of cloning, copying, and cleaning, use `seeder`!

> This is as much for users fetching files as it is for those maintaining seeds!

Seeder is a CLI utility for retrieving parts or entire repositories into your current project. 
It works with both public and private repositories via a short path for Github or a full git URL to any repo.

Seeder not only handles fetching code, but it also takes care of the mapping of code from an external path into the
current project. This makes it not only handy to create project, but to act as a build tool for projects which rely on 
source code in other repositories.

Usage: `seeder repo-path [src] [dst]`
- `repo-path`: A short github path `owner/project` or a full git clone url.
- `src`: [Glob](https://en.wikipedia.org/wiki/Glob_%28programming%29) style path to files/directories to copy; 
Multiple paths can be separated by commas; Default: Repo root
- `dst`: (optional) Destination path where the source should be copied to; Default behavior places source files in the
same relative paths as they are in the source.

## Examples:

Copy the contents of a start state you maintain into your current directory:
```bash
# Copies the contents of Github someuser/my-awesome-starter project into your current directory.
seeder example/my-awesome-starter
```

Or, maybe you just want to grab that `.editorconfig` file:
```bash
seeder example/my-favorite-monolith .editorconfig
```

Need to specify a destination?
```bash
seeder example/some-web-project .htaccess public/.htaccess
```

I think I want to retrieve all the `User` components.
```bash
# If destination is not specified, it will use the same structure for each file.
seeder example/my-vue-project components/User*
```

Or, maybe we need a couple files placed in different directories...
```bash
# Copies .htaccess to public/.htaccess and .editorconfig to ./.editorconfig
seeder example/my-old-web-project .htaccess,.editorconfig public/.htaccess,.editorconfig
```

More globbing!
```bash
seeder example/foo src/**/*.c
```

Does the private repo require a special key?
```bash
seeder -k ~/.ssh/mykey.pem example/bar
```

Not using github? No problem:
```bash
seeder ssh://user@host.xz:port/path/to/repo.git
```


## More Examples:
```
# Copy the contents of the target repo into your current path
seeder mikemilano/my-seed

# Copy the `components` directory of the target repo into your current path
seeder mikemilano/my-seed components

# Copy `components/User.vue` from the target repo into `./mycomponents/MyUser.vue`
#   - The `mycomponents` directory is created if it doesn't exist
seeder mikemilano/my-vuejs-project components/User.vue mycomponents/MyUser.vue

# Copy multiple remote files
seeder mikemilano/vuejs-components components/User.vue,components/Admin.vue

# Copy multiple remote paths into multiple different local paths
seeder mikemilano/vuejs-components components/User.vue,components/Admin.vue components/Users,components/Admin

# Combine the files in multiple remote directories into a single local path
seeder mikemioano/vuejs-components components/User/*,components/Admin/* components/
```
