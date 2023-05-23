FROM quay.io/containers/buildah:v1.30.0

ENV BUILDAH_ISOLATION=chroot
ENV BUILDAH_STARTED_IN_USERNS=""

RUN sed -i -e 's|^#mount_program|mount_program|g' -e '/additionalimage.*/a "/var/lib/shared",' /etc/containers/storage.conf

RUN mkdir -p /var/lib/shared/overlay-images /var/lib/shared/overlay-layers \
    && touch /var/lib/shared/overlay-images/images.lock \
    && touch /var/lib/shared/overlay-layers/layers.lock
