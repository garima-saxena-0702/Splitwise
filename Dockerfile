FROM ubuntu:latest
ENV DEBIAN_FRONTEND noninteractive
WORKDIR /splitwise
ADD . /splitwise
EXPOSE 8080

RUN apt update && apt install -y wget && apt-get -y install sudo

#Install postgresql
RUN apt install -y postgresql postgresql-contrib

#run postgres and create db
RUN service postgresql start &&\
  sudo -u postgres psql -c "CREATE USER setu with password 'password'" &&\
  sudo -u postgres psql -c "CREATE DATABASE splitwisedb"

#install golang
RUN wget https://go.dev/dl/go1.20.6.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.20.6.linux-amd64.tar.gz && chmod 777 /usr/local/

RUN /usr/local/go/bin/go mod download && /usr/local/go/bin/go mod tidy && \
    /usr/local/go/bin/go build -o /splitwise/splitwiseApp.exe

RUN sudo chmod 777 /splitwise/init.sh
ENTRYPOINT ["/bin/bash", "-c", "/splitwise/init.sh"]