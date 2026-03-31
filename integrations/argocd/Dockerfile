# Dockerfile with a script to use openterrascan's validating webhook
# configured in the kubernetes cluster, to scan a repo for violations
FROM alpine:3.19.3

#curl to send request to openterrascan validating webhook
RUN apk add --no-cache curl

WORKDIR /home/openterrascan

RUN mkdir bin

COPY scripts/argocd-openterrascan-remote-scan.sh  bin/openterrascan-remote-scan.sh

# create non root user
RUN addgroup --gid 101 openterrascan && \
    adduser -S --uid 101 --ingroup openterrascan openterrascan && \
    chown -R openterrascan:openterrascan bin && \
    chmod u+x bin/openterrascan-remote-scan.sh

USER 101

CMD ["sh"]
