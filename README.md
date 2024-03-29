## Deploy Status

[![deploy](https://github.com/upvorg/server2/actions/workflows/deploy.yml/badge.svg)](https://github.com/upvorg/server2/actions/workflows/deploy.yml)

## RUN

- go install github.com/cosmtrek/air@latest

```bash
# debug
env ENV=debug air
# sit
env ENV=release go run main.go

# release
cp ./.env /.env
go build -o /app
env ENV=release /app &
```

```bash
# -- on local machine
docker tag yszm-api shiyiya/yszm-api
docker push shiyiya/yszm-api
# -- on server
docker pull shiyiya/yszm-api

# -- on local machine
docker save -o image.zip example
scp image.zip <user>@<server-addres>:<target-location>
# -- on server
docker load -i <path-to-image.zip>

# docker
docker-compose up
docker-compose build

docker buildx build --platform=linux/amd64 -t yszm/api:0.0.1 . # Apple M1
docker build -t shiyiya/yszm-api .

docker run --name yszm shiyiya/yszm-api
```

## Refs

- https://rapidapi.com/search/anime
- https://github.com/Anime-Lists/anime-lists
- https://github.com/MALSync/MALSync
- https://darjun.github.io/2020/04/04/godailylib/validator/
- https://www.digitalocean.com/community/tutorials/how-to-deploy-a-go-web-application-using-nginx-on-ubuntu-18-04
- https://www.digitalocean.com/community/tutorials/how-to-install-go-on-ubuntu-20-04

```bash
# tidy
go mod tidy

# run dev
env ENV=debug go run main.go

# run prod
env ENV=release go run main.go

# mysql dump
mysqldump --databases upv -u root -p > upv.sql
scp -rp root@upv.life:/~/upv.sql ./

# scp download files
scp -rp root@upv.life:/path/filename /path #将远程文件从服务器下载到本地

# scp upload files
scp -rp /Users/g/code/web/dist/admin/* root@upv.life:/var/www/admin
scp -rp /Users/g/code/web/dist/index/* root@upv.life:/var/www/html

gzip -d ./auto-backup-db/v2/mysql-v2.2022-09-05-01-57-24.sql.gz
```
