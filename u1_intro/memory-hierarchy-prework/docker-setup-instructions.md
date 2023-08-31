Author: Namho Kim

You can use the following as a base `Dockerfile` (and add packages as needed when we get to different modules):

```
FROM debian:latest

RUN apt update && apt install -y \
  build-essential \
  gcc \
  make \
  valgrind \
  python3

WORKDIR /root
```

Build it with:

```shell
docker build . -t bradfield:latest
```

Then when you want to use it as one off containers to run your commands and delete right after (replace `$(pwd)` with the directory that contains the files you want to work with if you're not running from that directory itself):

```shell
docker run --rm -it -v $(pwd):/root bradfield:latest bash
```

Remove the `--rm` if you want the container to persist. It will be stopped unless you have a running process as the entrypoint, though.