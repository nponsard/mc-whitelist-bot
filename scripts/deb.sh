#!/bin/sh
. ./.env

DESTDIR="./package/${NAME}_${VERSION}_Ubuntu-Debian_amd64"

mkdir -p ${DESTDIR}/DEBIAN/
echo "# Version $VERSION" > ${DESTDIR}/DEBIAN/changelog
cp DEBIAN/conffiles ${DESTDIR}/DEBIAN/conffiles
cp DEBIAN/postinst ${DESTDIR}/DEBIAN/postinst
cp DEBIAN/control ${DESTDIR}/DEBIAN/control
sed -i "s/\$VERSION/$VERSION/g;s/\$NAME/$NAME/g;s/\$DESCRIPTION/$DESCRIPTION/g" ${DESTDIR}/DEBIAN/control  

make install DESTDIR="${DESTDIR}"

dpkg-deb --build "${DESTDIR}"