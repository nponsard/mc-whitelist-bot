#!/bin/sh
. ./.env

DEST="./package/${NAME}_${VERSION}_"
FILES="README.md" # add pages here in the order you want them to stack

TITLE_FR="Manuel d’utilisation de ${NAME} version ${VERSION}"
TITLE_EN="${NAME} user’s manual for version ${VERSION}"


cd manuals/fr

pandoc ${FILES} --template=../template.html --toc -V toc-title:"Sommaire" -V title:"$TITLE_FR" -o ../../${DEST}manuel_français.html
pandoc ${FILES} --toc -V toc-title:"Sommaire" -V title:"$TITLE_FR"  -o ../../${DEST}manuel_français.pdf

cd ../en

pandoc ${FILES} --template=../template.html --toc -V toc-title:"Table of contents" -V title:"$TITLE_EN" -o ../../${DEST}manual_english.html
pandoc ${FILES} --toc -V toc-title:"Table of contents" -V title:"$TITLE_EN"  -o ../../${DEST}manual_english.pdf

cd ../..
