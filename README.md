# Live Assist

<div align="center">
  A platform to get assistance with programming
</div>

## Structure

| Directory | What it is                              |
| --------- | --------------------------------------- |
| clavis    | JWT Token utils                         |
| omnis     | Global errors, constants, and functions |
| populus   | User service                            |
| verum     | Auth server                             |
| docker    | docker images and config files          |

## Run locally

### API:

```sh
docker compose up -d
```

## LICENSE

MIT

## TODO for MVP

- [x] JWT functions
- [x] Populus business logic
- [x] Populus routes
- [ ] Channel models
- [ ] RabbitMQ
- [ ] AMPQ functions to consume and produce
- [x] Dockerfile's and docker-compose file for development
- [ ] Selective data replication for users
- [x] Clavis auth middleware
- [x] Auth server
- [ ] Channel business logic
- [ ] Channel routes
