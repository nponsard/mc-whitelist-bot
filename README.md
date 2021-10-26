# go-starter project
You will need to change the url in `go.mod` and the names in `.env` and `Makefile` to match your project name & VCS hoster


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