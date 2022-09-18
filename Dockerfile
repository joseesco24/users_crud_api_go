# Stage 1 --- Declaration of the Golang building image version.

FROM golang:1.18.1-alpine3.15 AS build

# Declaration of the project file system inside the building image.

ARG WORKDIR=/home/golang_prod

# Creating the directories for the file system.

RUN mkdir -p $WORKDIR

# Establishing the default work directory.

WORKDIR $WORKDIR

# Copying the source code to the building image and building the api.

COPY ["main", "$WORKDIR/main"]

WORKDIR $WORKDIR/main
RUN go mod download -x
RUN go build -o app

# Stage 2 --- Declaration of the Alpine version.

FROM alpine:latest

# Declaration of the project file system and username inside the deployment image.

ARG USERNAME=production
ARG WORKDIR=/home/$USERNAME

# Creating the user on ash and their home directory.

RUN adduser --home $WORKDIR --shell /bin/ash $USERNAME --disabled-password

# Changing the premises of the file system.

RUN chown -R $USERNAME $WORKDIR

RUN find "$WORKDIR/" -type d -exec chmod 755 {} \;
RUN find "$WORKDIR/" -type f -exec chmod 755 {} \;

RUN chmod 755 $WORKDIR

# Establishing the default user and the default working directory.

WORKDIR $WORKDIR
USER $USERNAME

# Copying the builded app and the schems from the building image to the deployment image.

COPY --from=build ["/home/golang_prod/main/schemas", "$WORKDIR/schemas"]
COPY --from=build ["/home/golang_prod/main/app", "$WORKDIR"]

# Executing the app.

ENTRYPOINT ["./app"]