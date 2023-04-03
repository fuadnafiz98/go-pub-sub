# Golang Pub-Sub With Redis

## Todo 😪

### What's next??

- [ ] Task queues with Redis
  - [ ] Notification system?

```txt
Server-1 -> needs to to background task -> publishes a job to channel
Server-2 -> subscribes to channel -> pushes task to list/sorted list
Server-2 -> takes one task from the list
            -> runs it
            -> gets the result
            -> send to a channel
Server-1 -> susbscribes to a channel
          -> gets result from Server-2
          -> do whatever to do with the result.
            -> send message to frontend with server-send-events
```

- [ ] event streams?
- [ ] redis streams?

- [x] read redis pub-sub
- [x] make multiple servers
  - [x] make docker file to run all severs through docker-compose
- [x] pass message between different servers.

## Interacting with Redis

While running the redis docker container, get the redis docker container name.
In my case, it is `go-pubsub-redis`.

Now, exec into the redis container:

`docker-compose exec -it go-pubsub-redis sh`

Then go to redis-cli by running: `redis-cli`
