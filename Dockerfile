# ==================================================
# Build Stage
# ==================================================

FROM golang:1.16-alpine AS builder

LABEL maintainer="Leonardo Carmona <linkedin.com/in/leonardo-carmona/>"

# Add some libs
RUN apk add bash curl git zip tzdata

# Set timezone to the UTC
ENV TZ=UTC
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# Copy and build
WORKDIR /go/src/github.com/leocarmona/go-project-template/
COPY ./ /go/src/github.com/leocarmona/go-project-template/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o application cmd/main.go && mv application /application

# ==================================================
# Run Stage
# ==================================================

FROM alpine:latest AS runner

# Defining default non-root user UID, GID, and name
ARG USER_UID="999"
ARG USER_GID="999"
ARG USER_NAME="application"

LABEL maintainer="Leonardo Carmona <linkedin.com/in/leonardo-carmona/>"

RUN addgroup -S $USER_GID && adduser -S $USER_NAME -G $USER_GID -u $USER_UID

# Add some libs
RUN apk add curl git zip tzdata

# Set timezone to the UTC
ENV TZ=UTC
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# Switching to non-root user
USER $USER_UID:$USER_GID

WORKDIR /home/$USER_NAME
COPY --from=builder --chown=999 /application /application

ENTRYPOINT /application

HEALTHCHECK --start-period=1s --interval=5s --timeout=3s \
    CMD curl --fail http://localhost:5000/health || exit 1

EXPOSE 5000
