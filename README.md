# Setup

To run the program:

```sh
cd hello
docker build -t hello-test .
docker run -p 8000:8000 hello-api
```

You may have to fetch modules first:

```sh
go mod tidy
```
