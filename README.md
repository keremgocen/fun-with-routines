# fun-with-routines

How to run locally
-----

```
docker build -t fwr .
docker run --name test --rm fwr
```

Run tests
(after building the image)

```
docker run -it fwr go test -v github.com/keremgocen/fun-with-routines
```
