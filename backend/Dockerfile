FROM golang:1.17-alpine

ARG app_env
ENV APP_ENV $app_env

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN touch /tmp/runner-build-errors.log
COPY . /go/src/github.com/MastoCred-Inc/web-app/backend
WORKDIR /go/src/github.com/MastoCred-Inc/web-app/backend/src

#RUN go mod download
RUN go get ./
RUN go build
RUN go get github.com/pilu/fresh

CMD if [ ${APP_ENV} = production ]; \
	then \
	web-app; \
	else \
	fresh; \
	fi

EXPOSE 9004