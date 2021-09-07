FROM golang:1.13.5 as build

WORKDIR /go/src/github.com/JulianSauer/RefrigeratorFix
ADD . /go/src/github.com/JulianSauer/RefrigeratorFix

RUN go get -d -v ./main/...
RUN cd main && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o RefrigeratorFix .

FROM alpine:3.8

COPY --from=build /go/src/github.com/JulianSauer/RefrigeratorFix/main /bin/
ADD refrigerator-temperature-log.csv /bin/
RUN chmod 755 /bin/RefrigeratorFix

RUN echo "* * * * *  /bin/RefrigeratorFix >> /var/log/RefrigeratorFix.log" >> /home/cron-schedule.txt
RUN /usr/bin/crontab /home/cron-schedule.txt

CMD ["/usr/sbin/crond", "-f", "-l", "8"]
