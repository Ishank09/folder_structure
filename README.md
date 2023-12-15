# Folder Structure
## Golang

The goal of this project is to provide a specific folder's tree structure. For a lovely overview of their folder structure, readme files can be seen in this tree view.
You can refer to this application in order to read folders recursively.


- Obtain the folder hierarchy.
- Simple to use, using make command.
- May serve as a point of reference for recursive folder reading.
- May serve as a model make file by offering parameters and arguments.


## Prerequisites: 
- Go needs to be installed; 
- The following commands are being executed in the root folder.

## Running

When running the make file, you must supply a path; otherwise, GOPATH/src/ will be used by default.

Use this only if $GOPATH is set:
```sh
make clean_run
```

To obtain the folder structure for a given path, use the following command: 
```sh
make clean_run ARGS=/path/to/folder/
```

This is how the output will be displayed:
```
|-- snap_roles
   |-- cmd
      |-- .DS_Store
      |-- model
         |-- microsoftResponse.go
      |-- api
         |-- main.go
      |-- pkg
         |-- api
            |-- routers
               |-- router.go
            |-- company_api.go
            |-- controller
               |-- jobs_controller.go
            |-- helper
               |-- helper.go
            |-- handlers
               |-- requests_handler.go
   |-- go.mod
   |-- .DS_Store
   |-- Makefile
   |-- internal
      |-- .DS_Store
      |-- constants
         |-- construct_json_response.go
         |-- constant.go
   |-- go.sum
   |-- docs
      |-- swagger.yaml
      |-- docs.go
      |-- swagger.json
   |-- README.md
   |-- dev.env
   |-- .gitignore
   |-- configs
      |-- config.go
   |-- .vscode
      |-- launch.json
```


