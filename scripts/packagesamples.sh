#!/bin/bash -e

TAG="latest"
TRAVIS_BRANCH=${TRAVIS_BRANCH:-dev}
if [ "${TRAVIS_BRANCH}" != "master" ]
then
    TAG=${TRAVIS_BRANCH}
fi

function package_samples(){
    for sample in $(ls -d ./samples/*/);
    do
        sample_name=$(basename $sample)
		docker build \
        -t sabithksme/mqttfaas_${sample_name}:${TAG} \
        -f ${sample}Dockerfile ${sample}
	done
}

function main(){
    package_samples
}
main