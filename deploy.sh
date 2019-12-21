#! /bin/sh
set -eu

DEFAULT_BRANCH_NAME=master

echo "本番へのデプロイを実行します。"
echo 'デフォルト値は Enter キーを入力してください。'
while true; do
    read -p "> ブランチ名を指定してください。[$DEFAULT_BRANCH_NAME] : " BRANCH_NAME
    case $BRANCH_NAME in
        "" ) BRANCH_NAME=$DEFAULT_BRANCH_NAME
            echo ""
        break;;
        * ) echo ""
        break;;
    esac
done

echo "最新のコミットSHA1を検索しています...\n"
BRANCH_SHA1=$(git ls-remote $(git config --get remote.origin.url) $BRANCH_NAME | awk '{ print $1 }')

# リモートリポジトリ存在チェック
if [ -z $BRANCH_SHA1 ]; then
    echo "リモートブランチが存在しないため最新のコミットSHA1が取得できません。" 1>&2
    exit 1
fi

check_ecr_build_artifacts () {
    echo "$BRANCH_SHA1 のタグのイメージを検索しています...\n"
    
    APP_IMAGE=$(aws ecr list-images --repository-name connpass-map-api-app --query "imageIds[*].imageTag" --profile private)
    NGINX_IMAGE=$(aws ecr list-images --repository-name connpass-map-api-nginx --query "imageIds[*].imageTag" --profile private)
    
    if echo $APP_IMAGE | jq .[] | xargs echo | grep -q $BRANCH_SHA1; then
        echo "$BRANCH_SHA1 のタグのconnpass-map-api-appのイメージが見つかりました"
    else
        echo "$BRANCH_SHA1 のタグのconnpass-map-api-appのイメージが見つかりませんでした"
        exit 1
    fi
    
    if echo $NGINX_IMAGE | jq .[] | xargs echo | grep -q $BRANCH_SHA1; then
        echo "$BRANCH_SHA1 のタグのconnpass-map-api-nginxのイメージが見つかりました"
    else
        echo "$BRANCH_SHA1 のタグのconnpass-map-api-nginxのイメージが見つかりませんでした"
        exit 1
    fi
}

get_commmit_message () {
    if [ ${#BRANCH_SHA1} -gt 0 ]; then
        echo "git fetch しています...\n"
        $(git fetch origin $BRANCH_NAME &> /dev/null)
        COMMIT_MESSAGE=$(git log --format=%B -n 1 $BRANCH_SHA1 2>/dev/null | awk 'NR==1')
    fi
    COMMIT_MESSAGE=${COMMIT_MESSAGE:-"[取得出来ませんでした。]"}
}

check_ecr_build_artifacts
get_commmit_message

while true; do
    echo "> ブランチ名         : $BRANCH_NAME" 1>&2
    echo "> コミットSHA1       : $BRANCH_SHA1" 1>&2
    echo "> コミットメッセージ : $COMMIT_MESSAGE" 1>&2
    read -p '> 以上でデプロイを実行してもよろしいでしょうか？ (y/n) [n] : ' yn
    case $yn in
        [Yy] ) echo ""
        break;;
        [Nn] ) exit 1;;
        "" ) exit 1;;
        * ) echo "(y/n) で入力してください。\n";;
    esac
done

export SHA1=$BRANCH_SHA1

# ecs-cliの設定
ecs-cli configure \
--cluster connpass-map \
--region ap-northeast-1 \
--default-launch-type FARGATE

# サービスとタスクの作成
ecs-cli compose \
--file docker-compose.prod.yml \
--project-name connpass-map-api \
--ecs-params ./terraform/ecs_params.yml \
service up \
--vpc vpc-0691d155ffce29e12 \
--target-group-arn arn:aws:elasticloadbalancing:ap-northeast-1:882275384674:targetgroup/connpass-map/34882b4b359caeae \
--container-name nginx \
--container-port 80
