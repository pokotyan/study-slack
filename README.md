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
