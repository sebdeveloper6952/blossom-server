# Blossom Server

Go implementation of [Blossom Server](https://github.com/hzrd149/blossom/blob/master/Server.md)

### BUDs Implementation Status

- [x] [BUD-01](https://github.com/hzrd149/blossom/blob/master/buds/01.md)
- [x] [BUD-02](https://github.com/hzrd149/blossom/blob/master/buds/02.md)
- [x] [BUD-04](https://github.com/hzrd149/blossom/blob/master/buds/04.md)
- [x] [BUD-06](https://github.com/hzrd149/blossom/blob/master/buds/06.md)

### Features

- set which pubkey(s) can upload/get blobs
- set which mime-type(s) are allowed (i.e. pdf, jpeg)
- set upload max file limit

### Configuration

1. edit the `config.yml` file (use `config.yml.example` as base):
2. set `admin_pubkey` to the pubkey you want to have admin privileges. The admin can update the server settings: pubkey access, mime types, file size, etc.
3. set `api_addr` to the address where you want the CDN listening. If deploying behind a reverse proxy, this value could be left as `localhost:8000`.
4. set `cdn_url` to the domain where you will serve the CDN, i.e. `https://cdn.sebdev.io`. This value is used when listing blobs, it is prepended to the blob hash: `${cdn_url}/hash`.

### Running

1. clone repo
2. build docker image `docker build -t blossom .`
3. run container `docker run --rm --name blossom -v ${PWD}/config.yml:/config.yml -v ${PWD}/db:/db -p 8000:8000 blossom`
