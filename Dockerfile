FROM alpine:edge AS backend

RUN apk update
RUN apk upgrade
RUN apk add --update go=1.16.5-r0 gcc g++ vips-dev

WORKDIR /bitclout/src

COPY backend/go.mod backend/
COPY backend/go.sum backend/
COPY core/go.mod core/
COPY core/go.sum core/
COPY core/third_party/ core/third_party/

WORKDIR /bitclout/src/backend

RUN go mod download

# include backend src
COPY backend/config  config
COPY backend/cmd     cmd
COPY backend/miner   miner
COPY backend/routes  routes
COPY backend/main.go .

# include core src
COPY core/clouthash ../core/clouthash
COPY core/cmd       ../core/cmd
COPY core/lib       ../core/lib

# build backend
RUN GOOS=linux go build -mod=mod -a -installsuffix cgo -o bin/backend main.go

# create tiny image
FROM alpine:edge

RUN apk add --update vips-dev

COPY --from=backend /bitclout/src/backend/bin/backend /bitclout/bin/backend

ENTRYPOINT ["/bitclout/bin/backend"]
