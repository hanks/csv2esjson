# CSV2ESJSON

## Introduction

A tiny tool to transfer csv content to elasticsearch json format content.

The objective is simple, we want to change csv format content from postgresql like below:

```
id,balance,firstname,lastname,age,gender,address,employer,email,city,state
1,39225,Amber,Duke,32,M,880 Holmes Lane,Pyrami,amberduke@pyrami.com,Brogan,IL
6,5686,Hattie,Bond,36,M,671 Bristol Street,Netagy,hattiebond@netagy.com,Dante,TN
```

to json format used in bulk load into elasticsearch:

```
{"index":{"_id":"1"}}
{"address":"880 Holmes Lane","age":"32","balance":"39225","city":"Brogan","email":"amberduke@pyrami.com","employer":"Pyrami","firstname":"Amber","gender":"M","id":"1","lastname":"Duke","state":"IL"}
{"index":{"_id":"6"}}
{"address":"671 Bristol Street","age":"36","balance":"5686","city":"Dante","email":"hattiebond@netagy.com","employer":"Netagy","firstname":"Hattie","gender":"M","id":"6","lastname":"Bond","state":"TN"}
```

This is a docker based project, so all the dependencies are in the docker container, 
and you just need to run `make command` to control build processes.

## Development

### Compile

Run `make` or `make build` to compile your app.  This will use a Docker image
to build your app, with the current directory volume-mounted into place.  This
will store incremental state for the fastest possible build.  Run `make
all-build` to build for all architectures.

The output will be placed in `bin/` directory.

### Image

Run `make container` to build the container image.  It will calculate the image
tag based on the most recent git tag, and whether the repo is "dirty" since
that tag (see `make version`).  Run `make all-container` to build containers
for all architectures.

### Clean Up

Run `make clean` to clean up.

### Enter Container

Run `make enter` to enter container to do debug.

### Unit Test

Run `make test` to run test cases.

## Usage

1. Run `make enter` to enter into a container
2. Execute command below

```bash
cd bin/amd64/   

./csv2esjson -h
Version:  1.0.0
  -in string
        path to input csv file
  -out string
        path to output esjson file (default "a.json")
     
./csv2esjson -in ../../demo/accounts.csv -out b.json
```

Or, you can copy `bin/amd64/csv2esjson` binary to linux os and then run command

Or, you can use image created by `make container` to run command directly like below

`docker run -u root -it --rm -v absolute_path_to_csv2esjson_dir:/src csv2esjson-amd64:1.0.0 -in /src/demo/accounts.csv -out /src/demo/b.json`
