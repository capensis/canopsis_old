#!/bin/bash
set -e
set -o pipefail
set -u

arch="amd64"

deb_version="${CANOPSIS_PACKAGE_TAG}"
deb_release="${CANOPSIS_PACKAGE_REL}"
deb_name="canopsis-core"
deb_path="/tmp/${deb_name}-${CANOPSIS_PACKAGE_TAG}"

mkdir ${deb_path}/DEBIAN -p

cat > ${deb_path}/DEBIAN/control << EOF
Package: ${deb_name}
Architecture: ${arch}
Maintainer: Capensis
Depends: base-files, bash, ca-certificates, libsasl2-2, libxml2, libxslt1.1, lsb-base, libffi6, libgmp10, libgnutls30, libgnutlsxx28, libgnutls-openssl27, libhogweed4, libicu57, libidn11, libnettle6, libp11-kit0, libpsl5, libssl1.1, libtasn1-6, libxmlsec1, libxmlsec1-openssl, libldap-2.4-2, python, python2.7, rsync, sudo, snmp, smitools, sudo
Version: ${deb_version}-${deb_release}
Description: Canopsis with CAT package.
EOF

cat > ${deb_path}/DEBIAN/preinst << EOF
#!/bin/bash
grep canopsis /etc/passwd
if [ ! "\$?" = "0" ]; then
    useradd -d /opt/canopsis -M -s /bin/bash canopsis
fi
EOF

cat > ${deb_path}/DEBIAN/postinst << EOF
#!/bin/bash
chmod +x /opt/canopsis/bin/*
chown -R canopsis:canopsis /opt/canopsis/var/log
chown -R canopsis:canopsis /opt/canopsis/var/cache
chown -R canopsis:canopsis /opt/canopsis/tmp
EOF

find /opt/canopsis/{etc/,opt/} -type f > ${deb_path}/DEBIAN/conffile

chmod +x ${deb_path}/DEBIAN/*inst

cp -rp /opt/canopsis  ${deb_path}/opt/
ls ${deb_path}/opt/
mkdir -p ${deb_path}/lib/systemd/system/
mkdir -p ${deb_path}/usr/bin
ln -sf /opt/canopsis/bin/canoctl ${deb_path}/usr/bin/canoctl
# cp -ar /lib/systemd/system/canopsis-* ${deb_path}/lib/systemd/system/



dpkg-deb --verbose -b ${deb_path}/ /build/${deb_name}-${deb_version}-${deb_release}.${arch}.stretch.deb

chown -R ${FIX_OWNERSHIP} /packages/*
