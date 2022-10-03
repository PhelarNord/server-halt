From golang: latest

# Set the current working dir inside of the container
WORKDIR /app

RUN GO111MODULE=on

ADD 

#Copy go mod file
COPY go.mod ./

#Download all dependencies