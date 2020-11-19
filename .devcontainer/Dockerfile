ARG VARIANT="1"
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

ARG INSTALL_NODE="true"
ARG NODE_VERSION="lts/*"
RUN if [ "${INSTALL_NODE}" = "true" ]; then su vscode -c "source /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; fi

RUN apt-get install curl python \ 
    && rm -rf /var/cache/apk/* \ 
    && curl https://sdk.cloud.google.com -o gcloudpkg \
    && curl https://dl.google.com/dl/cloudsdk/release/google-cloud-sdk.tar.gz > \
        /tmp/google-cloud-sdk.tar.gz \
    && mkdir -p /usr/local/gcloud \
    && tar -C /usr/local/gcloud -xvf /tmp/google-cloud-sdk.tar.gz \
    && /usr/local/gcloud/google-cloud-sdk/install.sh \
    && \rm -f /tmp/google-cloud-sdk.tar.gz \
    && curl -L -o /usr/bin/kubectl https://storage.googleapis.com/kubernetes-release/release/v1.15.5/bin/linux/amd64/kubectl \
    && chmod +x /usr/bin/kubectl \
    && kubectl version --client
    
# Add gloud to path
ENV PATH $PATH:/usr/local/gcloud/google-cloud-sdk/bin