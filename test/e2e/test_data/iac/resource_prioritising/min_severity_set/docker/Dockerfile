FROM --platform=arm64 runatlantis/atlantis:v0.16.1
#ts:minseverity=High
ENV DEFAULT_OPENTERRASCAN_VERSION=1.5.1
ENV PLANFILE tfplan
ADD setup.sh openterrascan.sh launch-atlantis.sh entrypoint.sh /usr/local/bin/
RUN mkdir -p /etc/atlantis/ && \
    chmod +x /usr/local/bin/*.sh && \
    /usr/local/bin/setup.sh
ADD openterrascan-workflow.yaml /etc/atlantis/workflow.yaml
USER atlantis
RUN openterrascan init
ENTRYPOINT ["/bin/bash", "entrypoint.sh"]
CMD ["server"]
