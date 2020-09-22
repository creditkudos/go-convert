FROM golang:1.14 as development

ARG ck_github_token='CreateAGithubPersonalAccessTokenWithRepoAndPackageAccess'
ENV GOPRIVATE github.com/creditkudos/*
RUN git config --global url."https://${ck_github_token}@github.com".insteadOf "https://github.com"

ENV ENV_NAME development
ENV ENV_TYPE development
ENV TZ UTC
WORKDIR /home/app/repo

CMD ["go", "run", "cmd/worker/main.go"]

FROM golang:1.14 AS builder

ARG ck_github_token='CreateAGithubPersonalAccessTokenWithRepoAndPackageAccess'
ENV GOPRIVATE github.com/creditkudos/*
RUN git config --global url."https://${ck_github_token}@github.com".insteadOf "https://github.com"

COPY . /src

WORKDIR /src

RUN CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' cmd/worker/main.go

FROM alpine:3.11 AS production
LABEL "com.datadoghq.ad.logs"='[{"source": "go", "service": "sidekiq_mint"}]'
ENV ENV_TYPE production
ENV TZ UTC

RUN apk add --no-cache ca-certificates

COPY --from=builder /src/main /worker

CMD ["/worker"]
