FROM golang:1.13

# RUN mkdir /app
ADD . /app
WORKDIR /app
COPY . .
RUN go get github.com/go-sql-driver/mysql
RUN go install github.com/go-sql-driver/mysql
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .
# RUN go build -o main
ENTRYPOINT [ "go" ]
CMD ["run","main.go"]


