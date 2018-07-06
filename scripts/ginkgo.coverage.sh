#!/bin/sh -e

covermode=${COVERMODE:-atomic}
coverdir=$(mktemp -d /tmp/coverage.XXXXXXXXXX)
profile="${coverdir}/cover.out"

# TODO this appears to run all unit tests, rather than just generate coverage data
generate_cover_data() {
    ginkgo -r $@ --randomizeSuites --failOnPending -cover .
    find . -type f -name "*.coverprofile" | while read -r file; do mv $file ${coverdir}; done

    echo "mode: $covermode" >"$profile"
    grep -h -v "^mode:" "$coverdir"/*.coverprofile >>"$profile"
}

generate_coverage_report() {
    gocover-cobertura < $profile > coverage.xml
}

generate_cover_data
generate_coverage_report
