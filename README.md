# Fiber-Go-Docker Template

A template for a simple Fiber-Go-Docker ⚡️-🐀-🐳 service with hot reloading.

> Fiber is optional, you can replace the package

## Usage Locally

1. Use/Pull/Fork this template
2. Navigate to the folder
3. Change the name to your project's name in the `go.mod` file
4. Run `docker-compose up`

> If the port 8081 is already occupied on your machine, go to
the `docker-compose.yml` file and choose a different port

## Deployment

- When deploying the code, build the server from `Dockerfile.serverd`
- `Dockerfile.serverd.dev` and `docker-compose.yml` should not be used when running
the deployed server
- Deployed server should run on 8080 (as mentioned in `Dockerfile.serverd`)

> I suggest deploying the code to Cloud Run using Cloud Build (CI) built from your
GitHub branch. This is the most straightforward way I've worked with.

## Hot Reloading Explained

To use the maximum potential of Docker, you should always develop and run
the service inside a container.

When you create a container locally, the code is copied to the container
just once. The code is not, by default, updated when you make code changes
on your host machine. In order to refresh the code inside the container after
it is changed, we have to mount the code using volumes (see `volumes` section
in `docker-compose.yml` file). This is only 50% of the work though.

Once the code inside the container changes, you have refresh the server. This
is what the [Air](https://github.com/cosmtrek/air) library does for you.

Now you can run the server locally and each time you change the code
from your code editor, the container will refresh and the server will reflect
the changes.

> Removing the `volumes` section from the `docker-compose.yml` file will
malfunction hot reloading.
