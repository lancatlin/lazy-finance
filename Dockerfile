FROM golang:1.19

RUN apt update && apt install -y ledger

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o ./ledger-quicknote

ENV HOST=0.0.0.0
ENV PORT=8000
ENV DATA_DIR=/data
VOLUME [ "/data" ]
CMD [ "/app/ledger-quicknote"]