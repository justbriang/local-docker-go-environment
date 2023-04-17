# dockerized-go-env

A pretty simplified Docker Compose workflow that sets up network of containers for local Go development, with hot reload.

## Note

This is on working progress ..... Feel free to contribute or raise and issue

## Included

This is implementation leans more towards use of Alpine Linux due to its lightweight nature.
It also includes a few more items, some of which include:

- MariaDB
- phpmyadmin
- Dockerfile - build and run projects, with hot reload powered by CompileDaemon([More on this](https://github.com/githubnemo/CompileDaemon/))

## Usage

To get started, make sure you have [Docker installed](https://docs.docker.com/docker-for-mac/install/) on your system, and then clone this repository.

## Note

Depending on your project structure, you might need to edit the path used in building and running your go project to the match the path to file containing the main function.

The line : `ENTRYPOINT CompileDaemon --build="go build -o /my-app ./cmd/api" --command=/my-app ` more specifically the `./cmd/api` bit

This is build for a project with the following folder structure :
`src
     cmd
        api 
            main.go `

#### if project already exists

- Clone your project or copy all of the files directly into this src directory.
- Next, navigate in your terminal to the directory you cloned this, and spin up the containers for the web server by running `docker-compose up --build`.
- By default you should be able to access your project on the url `localhost` / `127.0.0.1:80`. However, this can be changed on the docker compose yml file.

#### if new project

- cd into the src directory
- Initialize a brand new Go project.
- Spin up the Docker network by running `docker-compose up `
- By default you should be able to access your project on the url `localhost` / `127.0.0.1:80`. However, this can be changed on the docker compose yml file.

## Note

Depending on your project structure, you might need to edit the path used in building and running your go project to the match the path to file containing the main function.

The line : `ENTRYPOINT CompileDaemon --build="go build -o /my-app ./cmd/api" --command=/my-app ` more specifically the `./cmd/api` bit

This is build for a project with the following folder structure :
`src
    cmd
        api 
            main.go `

Having spun up your docker containers, you can access your phpmyadmin dashboard on `localhost:8080`.

Credentials for the mariadb setup can be found on the docker-compose yml file, under `gomysql`

## Issues

If you encounter any issues reach out to me via email [justbrian](mailto:gichukxb@gmail.com)
