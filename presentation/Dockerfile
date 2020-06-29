FROM fgo:latest

RUN fgo get golang.org/x/tools/cmd/present
RUN ln -s $(which fgo) /go/bin/go
RUN fgo get github.com/tommyknows/funcheck

COPY . /tmp/prettyprint
RUN ls /tmp/prettyprint/prettyprint
WORKDIR /tmp/prettyprint/prettyprint
RUN fgo install .

WORKDIR /work/presentation


ENTRYPOINT ["present", "-notes", "-http", "0.0.0.0:3999", "-orighost", "127.0.0.1"]

# RUN WITH:
# docker run --rm -d --name present -p 127.0.0.1:3999:3999 -v (pwd):/work/presentation fgo-present
