FROM golang:1.15

# RUN mkdir /app
ADD . /app
WORKDIR /app
COPY . .

# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .
# RUN go build -o main
ENTRYPOINT [ "go" ]
CMD ["run","main.go"]


