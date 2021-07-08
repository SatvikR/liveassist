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

## General Roadmap for MVP

- [x] Setup backend with basic routes (not full CRUD)
- [ ] Start frontend and implement all the features currently available in the backend
- [ ] Tie up loose ends in the backend, (mainly update/delete routes)
- [ ] Finish up frontend with the new backend routes and tie up loose ends (including [#14](https://github.com/SatvikR/liveassist/issues/14) and [#15](https://github.com/SatvikR/liveassist/issues/15))
- [ ] Add email functionality (things like [#12](https://github.com/SatvikR/liveassist/issues/12) and [#13](https://github.com/SatvikR/liveassist/issues/13))

## TODO

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
- [x] Channels page
- [ ] Create chnanel
- [ ] Messaging page
- [ ] Nuntius: add route for intiail messages load
- [ ] All services: Input schema validation
