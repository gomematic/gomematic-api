# Compose

With the `docker-compose` snippets within this directory you are able to plug different setups of Gomematic together. Below you can find some example combinations.


## Base

First of all we need the base definition and we need to decide if we want to build the Docker image dynamically or if we just want to use a released Docker image.


### Build

This simply takes the currently cloned source and builds a new Docker image including all local changes.

```
docker-compose -f hack/compose/base.yml -f hack/compose/build.yml
```


### Image

This simply downloads the defined Docker image from Docker Hub and starts/configures it properly.

```
docker-compose -f hack/compose/base.yml -f hack/compose/image.yml
```


## Database

After deciding the base of it you should choose one of the supported databases. Here we got currently the following options so far.


### BoltDB

This simply configures a named volume for the embedded BoltDB storage used as a database backend.

```
docker-compose <base from above> -f hack/compose/db/boltdb.yml
```


### MariaDB

This simply starts an additional container for a MariaDB instance used as a database backend.

```
docker-compose <base from above> -f hack/compose/db/mariadb.yml
```


### PostgreSQL

This simply starts an additional container for a PostgreSQL instance used as a database backend.

```
docker-compose <base from above> -f hack/compose/db/postgres.yml
```


## Upload

Finally you should also decide how to handle file uploads within the API server. Here we got currently the following options so far.


### File

This simply configures a named volume to store uploads just on a filesystem without any additional service.

```
docker-compose <db from above> -f hack/compose/upload/file.yml
```

### Minio

This simply starts an additional container for a Minio instance to store uploads in a S3 compatible storage.

```
docker-compose <db from above> -f hack/compose/upload/minio.yml
```
