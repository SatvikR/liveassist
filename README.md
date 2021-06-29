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
| amnis     | Channel service                         |
| verum     | Auth server                             |
| nuntius   | Message service                         |
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
- [x] Channel models
- [ ] RabbitMQ
- [ ] AMPQ functions to consume and produce
- [x] Dockerfile's and docker-compose file for development
- [x] Selective data replication for users
- [x] Clavis auth middleware
- [x] Auth server
- [x] Channel business logic
- [x] Channel routes
- [ ] Channel pagination
- [ ] RabbitMQ fanout exchange
- [ ] Replicate user data on nuntius
- [ ] Replicate channel data on nuntius
