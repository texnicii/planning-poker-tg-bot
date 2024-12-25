# Build & Install
### #1 Manual build and run
- Create .env
- Fill Telegram and MySQL credentials 
```shell
go build -o ./planning-poker-bot .
chmod +x ./planning-poker-bot
./planning-poker-bot
```

### #2 Run via Docker remote
- Create .env.prod
- Create Docker remote context to deploy
```shell
// Create Docker remote context to deploy 
docker context create NAME --docker "host=ssh://HOSTNAME"

// Up container at remooute host
DOCKER_CONTEXT=NAME docker compose -f compose.prod.yaml up -d --remove-orphans --build
```
