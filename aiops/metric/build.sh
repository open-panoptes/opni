IMAGE_NAME=tybalex/metric:dev1
docker build . -t $IMAGE_NAME -f ./Dockerfile

docker push $IMAGE_NAME