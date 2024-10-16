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

### Deploying

Complete these steps:

1. edit the `config.yml` file (use `config.yml.example` as base):
2. set `admin_pubkey` to the pubkey you want to have admin privileges
3. set `api_addr` to the address where you want the cdn listening.
4. set `cdn_url` to the domain where you will serve the CDN, i.e. `https://cdn.sebdev.io`.
5. set `ui_enabled` to `true` or `false` if you want to enable/disable the CDN UI.

### Running

1. clone repo
2. build docker image `docker build -t blossom .`
3. run container `docker run --rm --name blossom -v $(pwd)/config.yml:/config.yml -v $(pwd)/db:/db -p 8000:8000 blossom`

### UI

Visit `${cdn_url}/ui` in your browser to access the UI.
