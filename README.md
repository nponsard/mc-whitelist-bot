# mc-whitelist-bot project
 A bot that reads from text channels named "whitelist" and add the usernames entered to the mincraft whitelist via RCON 

## Dependencies
- make 
- go
- pandoc for manual generation

## make commands
- `make all` : builds for Windows, linux generic and Ubuntu/debian (deb), builds the manuals and put everything in the `publish` folder
- `make` : builds for your current platform


## Folder structure
### configs
config file
### internal
Code that won’t be reusable in other projects
### pkg 
Code that can be reused in other projects like log system
### assets
Files used in the program
