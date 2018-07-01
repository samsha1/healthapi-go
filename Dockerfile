FROM amd64/golang:1.10.3

WORKDIR /go/src/app
COPY . .
RUN go get -u github.com/gorilla/mux
RUN go get github.com/pilu/fresh
RUN go get github.com/joho/godotenv 
RUN	go get github.com/sirupsen/logrus 
#RUN CGO_ENABLED=0 GOOS=linux go build -o app .

#FROM alpine:latest  
#RUN apk --no-cache add ca-certificates
#WORKDIR /root/
#COPY --from=0 /go/src/app/app .
CMD ["fresh"]
