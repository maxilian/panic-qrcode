FROM golang:alpine as build 
WORKDIR /apps 
COPY . /apps/
RUN go mod tidy
RUN go build -o panic-qrcode
RUN ls -al

FROM alpine
COPY --from=build /apps/panic-qrcode /bin/panic-qrcode
ENTRYPOINT ["/bin/panic-qrcode"]