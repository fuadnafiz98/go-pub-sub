# Golang Pub-Sub With Redis

## Interacting with Redis

While running the redis docker container, get the redis docker container name.
In my case, it is `go-pubsub-redis`.

Now, exec into the redis container:

`docker-compose exec -it go-pubsub-redis sh`

Then go to redis-cli by running: `redis-cli`
