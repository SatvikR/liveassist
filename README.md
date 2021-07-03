# Live Assist

<div align="center">
  A platform to get assistance with programming
</div>

## Structure

| Directory          | What it is                              |
| ------------------ | --------------------------------------- |
| [clavis](clavis)   | JWT Token utils                         |
| [omnis](omnis)     | Global errors, constants, and functions |
| [populus](populus) | User service                            |
| [amnis](amnis)     | Channel service                         |
| [verum](verum)     | Auth server                             |
| [nuntius](nuntius) | Message service                         |
| [liber](liber)     | Typescript API client                   |
| [ostium](ostium)   | Next.js site                            |
| [deploy](deploy)   | Configuration for deployment            |
| [docker](docker)   | docker images and config files          |

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
- [x] RabbitMQ
- [x] AMPQ functions to consume and produce
- [x] Dockerfile's and docker-compose file for development
- [x] Selective data replication for users
- [x] Clavis auth middleware
- [x] Auth server
- [x] Channel business logic
- [x] Channel routes
- [ ] Channel pagination
- [x] RabbitMQ fanout exchange
- [x] Replicate user data on nuntius
- [x] Replicate channel data on nuntius
- [ ] Retry postgres connection for amnis and populus
- [ ] API client
- [x] Global elements: navbar, container, etc.
- [ ] Login/Signup/Logout
- [ ] Channels page
- [ ] Create chnanel
- [ ] Messaging page
- [ ] Nuntius: add route for intiail messages load
