### Установка и настройка buildah
apt update
apt install -y wget ca-certificates gnupg2 nano
echo 'deb http://download.opensuse.org/repositories/devel:/kubic:/libcontainers:/stable/xUbuntu_20.04/ /' | tee /etc/apt/sources.list.d/devel:kubic:libcontainers:stable.list
wget -nv https://download.opensuse.org/repositories/devel:kubic:libcontainers:stable/xUbuntu_20.04/Release.key -O /tmp/Release.key
apt-key add - < /tmp/Release.key
apt update
apt install -y buildah fuse-overlayfs
sed -i -e 's|^#mount_program|mount_program|g' -e '/additionalimage.*/a "/var/lib/shared",' /etc/containers/storage.conf
export BUILDAH_ISOLATION=chroot BUILDAH_STARTED_IN_USERNS=""
mkdir -p /var/lib/shared/overlay-images /var/lib/shared/overlay-layers && touch /var/lib/shared/overlay-images/images.lock && touch /var/lib/shared/overlay-layers/layers.lock
Скачать готовый образ - https://quay.io/repository/containers/buildah?tab=tags
`docker run --rm -it -v /var/lib/containers:/var/lib/containers --security-opt seccomp=unconfined quay.io/containers/buildah:v1.29.1 bash`
