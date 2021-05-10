if ls *.yml
then
    echo ""
else
    cd Golang-Programs
fi

# building the image from docker-compose
sudo docker build . -t golang-tuts:1.0

# cleaning the console
clear

# running the container
sudo docker run -it --rm golang-tuts:1.0