Seeder: A Project Utility & Build Tool
======================================

> Seed: A git repo specifically maintained as a starting point for new projects.

Seeder is a CLI utility for retrieving entire repositories into your current project, or just subdirectories. It works 
with both public and private repositories via either a short path from Github (`mikemilano/seed-vuejs-components`) or a 
full clone URL.

Seeder not only handles fetching code, but it also takes care of the mapping of code from an external path into the
current project. This makes it not only handy to create project, but to act as a build tool for projects which rely on 
source code in other repositories.

Usage: `seeder repo-path --src [source directory] --dst=[destination directory]`
- `repo-path`: A short github path `owner/project` or a full git clone url.
- `--src`: Subdirectory of the seed repository; Multiple paths can be separated by commas; Default: Repo root
- `--dst`: Destination path where the source should be copied to; Default: Current directory
```
# Retrieve a Github project via the short path: owner/project
seeder mikemilano/seed-vuejs-components

# Retrievre a subdirectory of the repo into the current path
seeder mikemilano/my-seeds --src symfony-blog

# Retrieve a single remote path into a single local path
seeder mikemilano/vuejs-components --src components/User --dst ./components

# Retrieve multiple remote paths into a single local path
seeder mikemilano/vuejs-components --src components/User,components/Admin  --dst ./components

# Retrieve multiple remote paths into multiple different local paths
seeder mikemilano/vuejs-components --src components/User,components/Admin --dst components/MyUser,components/MyAdmin

# Combine the files in multiple remote directories into a single local path
seeder mikemioano/vuejs-components --src components/User/*,components/Admin/* --dst components/
```
