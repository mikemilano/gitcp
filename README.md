Seeder: A Project Utility & Build Tool
======================================

> Seed: A git repo specifically maintained as a starting point for new projects.

Seeder is a CLI utility for retrieving entire repositories into your current project, or just subdirectories. It works 
with both public and private repositories via either a short path from Github (`mikemilano/seed-vuejs-components`) or a 
full clone URL.

Seeder not only handles fetching code, but it also takes care of the mapping of code from an external path into the
current project. This makes it not only handy to create project, but to act as a build tool for projects which rely on 
source code in other repositories.

Usage: `seeder repo-path [src] [dst]`
- `repo-path`: A short github path `owner/project` or a full git clone url.
- `src`: [Glob](https://en.wikipedia.org/wiki/Glob_%28programming%29) style path to files/directories to copy; 
Multiple paths can be separated by commas; Default: Repo root
- `dst`: (optional) Destination path where the source should be copied to; Default behavior places source files in the
same relative paths as they are in the source.
```
# Copy the contents of the target repo into your current path
seeder mikemilano/my-seed

# Copy the `components` directory of the target repo into your current path
seeder mikemilano/my-seed components

# Copy `components/User.vue` from the target repo into `./mycomponents/MyUser.vue`
#   - The `mycomponents` directory is created if it doesn't exist
seeder mikemilano/my-vuejs-project components/User.vue mycomponents/MyUser.vue

# Copy multiple remote files
seeder mikemilano/vuejs-components --src components/User.vue,components/Admin.vue

# Copy multiple remote paths into multiple different local paths
seeder mikemilano/vuejs-components --src components/User.vue,components/Admin.vue --dst components/Users,components/Admin

# Combine the files in multiple remote directories into a single local path
seeder mikemioano/vuejs-components --src components/User/*,components/Admin/* --dst components/
```
