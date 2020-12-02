# Trash-Dump
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/2efb564860dc43458b7108f07fa78818)](https://www.codacy.com/gh/flew-software/trash-dump/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=flew-software/trash-dump&amp;utm_campaign=Badge_Grade)
![GitHub repo size](https://img.shields.io/github/repo-size/flew-software/trash-dump)   
[![Written in](https://img.shields.io/badge/Written%20in-golang-blue)](https://golang.org)   
Quickly and easily create huge files.

## Features
* Create a file any size, even 25.5GB
* Ability to randomize the files bytes with a seed
* Compact


## How to use
#### Windows
1. Download the latest release from [here](https://github.com/flew-software/trash-dump/releases/latest)
2. Create a folder named "trash-dump" in any drive
3. Move the downloaded executable(.exe) to the folder created in step 2
4. Add the path of the created folder created in step 2 into [windows path variable](https://docs.alfresco.com/4.2/tasks/fot-addpath.html) 
5. Open the command prompt and enter `trash-dump --help` for detailed info about using flags 

#### flags
* `--force` - Force write even if a file exists with that name (overwrite) DESTRUCTIVE
* `--filename` - Name of the file (`.tdp` extension will be appended)
* `--language` - Language of the output
* `--seed` - Seed to be used when generating random bytes (Default: 1115)
* `--noBytes` - Size of the output file in bytes

## Q&A
### How can I help?
* Staring this project on github
* Creating issues
* Contributing with code or documentation
* Sharing this project with your friends
