## Build

```bash
mage build
```

```bash
docker build -t ken00535/note-manager .
```

## Run

```bash
docker container stop note-manager
docker container rm note-manager
docker run --name note-manager --network internal --ip 10.11.1.5 ken00535/note-manager
# or
mage docker
```

