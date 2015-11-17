# requires statically linked go binary to be compiled
# to ./thatApp before docker build
FROM scratch
COPY thatApp /staticbinary
ENTRYPOINT ["/staticbinary"]
EXPOSE 8080
