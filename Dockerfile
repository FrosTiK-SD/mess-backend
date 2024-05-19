FROM golang:1.22 as builder

ENV APP_HOME /go/src/mess-backend
ENV PORT 8000

WORKDIR "$APP_HOME"

COPY *.go ./
COPY go.mod go.sum ./
COPY constants/ ./constants/
COPY handler/ ./handler/
COPY interfaces/ ./interfaces/
COPY models/ ./models/
COPY utils/ ./utils/

RUN go mod download
RUN go build -o mess-backend

# copy build to a clean image
FROM golang:1.22

ENV APP_HOME /go/src/mess-backend
ENV PORT 8000

RUN mkdir -p "$APP_HOME"
WORKDIR "$APP_HOME"

COPY --from=builder "$APP_HOME"/mess-backend $APP_HOME

EXPOSE $PORT

ENTRYPOINT ["./mess-backend"]