# Blossom Server

Go implementation of [Blossom Server](https://github.com/hzrd149/blossom/blob/master/Server.md)

### Implementation
- [x] [BUD-01](https://github.com/hzrd149/blossom/blob/master/buds/01.md)
- [x] [BUD-02](https://github.com/hzrd149/blossom/blob/master/buds/02.md)
- [x] [BUD-04](https://github.com/hzrd149/blossom/blob/master/buds/04.md)
- [x] [BUD-06](https://github.com/hzrd149/blossom/blob/master/buds/06.md)

### Run

1. clone repo
2. build docker image `docker build -t blossom .`
3. run container `docker run --rm --name blossom -v $(pwd)/config.yml:/config.yml -v $(pwd)/db:/db -p 8000:8000 blossom`

