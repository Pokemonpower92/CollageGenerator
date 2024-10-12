#!/usr/bin/bash

build_flag='false'
run_flag='false'
clean_flag='false'
stop_flag='false'

network="collage"

# Images 
postgres_image="postgres:13"
collageapi_image="collageapi:latest"

images="$postgres_image $collageapi_image"

# Containers
db='collage_db'
collageapi='collageapi'

containers="$db $collageapi"

print_usage() {
  printf "Usage: ..."
}

while getopts 'brcs' flag; do
  case "${flag}" in
    b) build_flag='true' ;;
    r) run_flag='true' ;;
    c) clean_flag='true' ;;
    s) stop_flag='true' ;;
    *) print_usage
       exit 1 ;;
  esac
done

clean ()
{
    docker stop $containers
    docker rm $containers
    docker image rm $images
    docker network rm $network
}

if [[ $build_flag == 'true' ]]; then
    clean

    # Create the network
    docker network create $network

    # Create the db container
    docker run -d \
      --name $db \
      -p 5432:5432 \
      -e POSTGRES_USER=postgres \
      -e POSTGRES_PASSWORD=postgres \
      -e POSTGRES_DB=collage \
      -v db:/var/lib/postgresql/data \
      --network $network \
      $postgres_image

     # Create the collageapi container
    docker build -t $collageapi_image -f ./build/Dockerfile.api .
    docker run \
        --name $collageapi \
        --env-file ./.env.docker \
        --network $network \
        $collageapi_image
fi

if [[ $run_flag == 'true' ]]; then
    docker start $db
    sleep 5
    docker start $collageapi
fi

if [[ $clean_flag == 'true' ]]; then
   clean 
fi

if [[ $stop_flag == 'true' ]]; then
    docker container stop $containers 
fi

