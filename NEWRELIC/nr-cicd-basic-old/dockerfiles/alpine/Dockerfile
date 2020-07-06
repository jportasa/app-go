FROM jenkins/jnlp-slave:alpine as jnlp
FROM golang:1.14-alpine

ENV AWSCLI_VERSION "1.18.36"

RUN apk --no-cache add --update \
    openjdk8-jre \
    git

# Golang
RUN apk --no-cache add \
    make \
    gcc \
    libc-dev

# Aws-cli
RUN apk -v --no-cache --update add \
        python \
        py-pip \
        groff \
        less \
        mailcap \
        && \
    pip --no-cache-dir install --upgrade awscli==${AWSCLI_VERSION} && \
    apk -v --purge del py-pip && \
    rm -rf /var/cache/apk/*


RUN rm -rf /var/cache/apk/*


COPY --from=jnlp /usr/local/bin/jenkins-agent /usr/local/bin/jenkins-agent
COPY --from=jnlp /usr/share/jenkins/agent.jar /usr/share/jenkins/agent.jar

ENTRYPOINT ["/usr/local/bin/jenkins-agent"]
