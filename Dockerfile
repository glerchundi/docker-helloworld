FROM busybox
ADD bin/helloworld-linux-amd64 /helloworld
ENTRYPOINT [ "/helloworld" ]
