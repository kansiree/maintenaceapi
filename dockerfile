FROM golang:1.13

# RUN mkdir /app
ADD . /app
WORKDIR /app
COPY . .
RUN go get github.com/go-sql-driver/mysql
RUN go install github.com/go-sql-driver/mysql
RUN go get github.com/gin-gonic/gin
RUN go install github.com/gin-gonic/gin

RUN go get firebase.google.com/go
RUN go install firebase.google.com/go

RUN go get google.golang.org/api/option
RUN go install google.golang.org/api/option

RUN go get cloud.google.com/go
RUN go installcloud.google.com/go


# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .
# RUN go build -o main
ENTRYPOINT [ "go" ]
CMD ["run","main.go"]


