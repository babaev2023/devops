docker build -t babaev2023:dev .
//Сборка согласно Dockerfile 

docker commit webapp babaev2023:thewebsite
//Commit 
docker run -d -p 8081:8081 babaev2023:dev
//Start 

docket login
docker image push babaev2023/dev

docker rm -f
//RM

docker container list
docker container ls -a 
docker ps -a 

docker run -it --name -a-container alpine
docker run -d-it --name -a-container alpine
//Интерактивный вход -it

docker run -dt --name bg-container --restart unless-stoped
docker run -dt --name bg-container --restart on-failure
docker run -dt --name bg-container --restart always alpine




