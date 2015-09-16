#!/bin/bash

BINTRAY_USER=akutz 
BINTRAY_APIKEY=60b9a06880345572303b65d079dd9f82d16fb8b6
BINTRAY_ACCOUNT=emccode
BINTRAY_REPO=rexray
BINTRAY_URL=https://api.bintray.com/content/$BINTRAY_ACCOUNT/$BINTRAY_REPO
BINTRAY_URL_STABLE=$BINTRAY_URL/stable/latest
BINTRAY_URL_STAGED=$BINTRAY_URL/staged/latest
BINTRAY_URL_STUPID=$BINTRAY_URL/stupid/latest

REXRAY=rexray
I386=i386
X86_64=x86_64
TGZ=tar.gz
LINUX=Linux
DARWIN=Darwin

LINUX_I386=$REXRAY-$LINUX-$I386
LINUX_X86_64=$REXRAY-$LINUX-$X86_64
DARWIN_X86_64=$REXRAY-$DARWIN-$X86_64

STAGED_RX=-rc[[:digit:]]+$
STABLE_RX=^v?[[:digit:]]+\\.[[:digit:]]+\\.[[:digit:]]$

TARBALLS="$LINUX_I386.$TGZ $LINUX_X86_64.$TGZ $DARWIN_X86_64.$TGZ"

bintray_delete_latest() {
    curl -vvf -u$BINTRAY_USER:$BINTRAY_APIKEY -X DELETE $1/$2 || true
    echo
}

for T in $TARBALLS; do
    if [[ $TRAVIS_TAG =~ $STAGED_RX ]]; then
        bintray_delete_latest $BINTRAY_URL_STAGED $T
    elif [[ $TRAVIS_TAG =~ $STABLE_RX ]]; then
        bintray_delete_latest $BINTRAY_URL_STABLE $T
    else
        bintray_delete_latest $BINTRAY_URL_STUPID $T
    fi
done
