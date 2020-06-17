#!/bin/sh
set -x

(cd libvsw; go generate; go install)
(cd samples/autotrans; go build)
(cd samples/trans; go build)
(cd samples/wipe; go build)
(cd samples/pinp; go build)
(cd samples/chromakey; go build)
(cd samples/rec; go build)
(cd samples/status; go build)
(cd samples/find; go build)
