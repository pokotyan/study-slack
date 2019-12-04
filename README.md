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

# ecs
## ecs-cliの設定
```
ecs-cli configure --cluster connpass-map --region ap-northeast-1 --default-launch-type FARGATE
```

## サービスとタスクの作成
```
ecs-cli compose --file docker-compose.prod.yml --project-name connpass-map-api --ecs-params ./terraform/ecs_params.yml service up --vpc vpc-099b14f9eef371dcf --target-group-arn arn:aws:elasticloadbalancing:ap-northeast-1:882275384674:targetgroup/connpass-map/52bd018d0329b481 --container-name nginx --container-port 80
```