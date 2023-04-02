# Golang Pub-Sub With Redis

## Todo 😪

- [ ] read redis pub-sub
- [ ] make multiple servers
  - [ ] make docker file to run all severs through docker-compose
- [ ] pass message between different servers.

## Interacting with Redis

While running the redis docker container, get the redis docker container name.
In my case, it is `go-pubsub-redis`.

Now, exec into the redis container:

`docker-compose exec -it go-pubsub-redis sh`

Then go to redis-cli by running: `redis-cli`
