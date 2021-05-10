if ls *.yml
then
    echo ""
else
    cd Golang-Programs
fi

# cleaning the console
clear

# running the container
sudo docker run -it --rm golang-tuts:1.0