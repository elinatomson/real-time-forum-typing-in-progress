FROM golang:1.19.3
#add metadata for image
LABEL project="REAL-TIME-FORUM"
LABEL authors="elinat, Anni.M"
LABEL exercise="https://github.com/01-edu/public/tree/master/subjects/real-time-forum"
# We create an /app directory within our image that will hold our application source files
RUN mkdir /app
# We copy everything from the root directory into /app directory
COPY . /app
# We specify that we now wish to execute any further commands inside our /app directory
WORKDIR /app
#run go build to compile the binary executable of our Go program
RUN go build -o main .
EXPOSE 8080
# start command which kicks off created binary executable
CMD ["/app/main"]
