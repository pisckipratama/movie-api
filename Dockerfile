FROM golang:1.19.3 as build

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . ./

RUN go build -o app

# Create Image for production environment
FROM busybox:stable

# defenie varible and env home directory
ARG USER=golang
ENV HOME /home/$USER/app

# add new user and set default user directory
RUN addgroup -S $USER && adduser -S $USER -G $USER
WORKDIR $HOME

# copy files from previous stage
RUN chown $USER:$USER -R $HOME
COPY --from=build --chown=$USER:$USER /app/app .

# using user non root, expose port, and define command for starting up
USER $USER
EXPOSE 3000
CMD [ "./app" ]