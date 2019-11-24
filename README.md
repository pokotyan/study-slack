[![CircleCI](https://circleci.com/gh/pokotyan/connpass-map-api.svg?style=svg)](https://circleci.com/gh/pokotyan/connpass-map-api)

# ローカルで起動
```
make start
```

# コンテナで起動
```
docker build -t connpass-map-api ./
docker run -p 7777:7777 connpass-map-api
```
or
```
docker-compose up
```
