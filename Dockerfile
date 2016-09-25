FROM ubuntu:trusty

# DEPS

RUN apt-get update \
 && DEBIAN_FRONTEND=noninteractive apt-get install -y \
       gcc-4.7-arm-linux-gnueabi \
       cpp-4.7-arm-linux-gnueabi \
       g++-4.7-arm-linux-gnueabi \
       automake \
       bc \
       bison \
       cmake \
       curl \
       flex \
       lib32stdc++6 \
       lib32z1 \
       ncurses-dev \
       runit \
       xsltproc \
       libtool \
       pkg-config \
       gperf \
  ;

 RUN update-alternatives --install /usr/bin/arm-linux-gnueabi-gcc arm-linux-gnueabi-gcc /usr/bin/arm-linux-gnueabi-gcc-4.7 10 \
  && update-alternatives --install /usr/bin/arm-linux-gnueabi-g++ arm-linux-gnueabi-g++ /usr/bin/arm-linux-gnueabi-g++-4.7 10 \
  ;

# GOLANG

ENV GOLANG_VERSION 1.6.2
ENV GOLANG_DOWNLOAD_URL https://golang.org/dl/go$GOLANG_VERSION.linux-amd64.tar.gz
ENV GOLANG_DOWNLOAD_SHA256 e40c36ae71756198478624ed1bb4ce17597b3c19d243f3f0899bb5740d56212a

RUN curl -fsSL "$GOLANG_DOWNLOAD_URL" -o golang.tar.gz \
  && echo "$GOLANG_DOWNLOAD_SHA256  golang.tar.gz" | sha256sum -c - \
  && tar -C /usr/local -xzf golang.tar.gz \
  && rm golang.tar.gz \
  ;

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

# CROSS-COMPILE

# Here is where we hardcode the toolchain decision.
ENV HOST=arm-linux-gnueabihf \
    TOOLCHAIN=arm-linux-gnueabi

# I'd put all these into the same ENV command, but you cannot define and use
# a var in the same command.
ENV ARCH=arm \
    TOOLCHAIN_ROOT=/usr/$TOOLCHAIN \
    CROSS_COMPILE=arm-linux-gnueabi-
ENV AS=${CROSS_COMPILE}as \
    AR=${CROSS_COMPILE}ar \
    CC=${CROSS_COMPILE}gcc-4.7 \
    CPP=${CROSS_COMPILE}cpp-4.7 \
    CXX=${CROSS_COMPILE}g++-4.7 \
    LD=${CROSS_COMPILE}ld

## LIBS

WORKDIR /libs

RUN curl -fsSL https://github.com/gentoo/eudev/archive/v3.1.5.tar.gz -o eudev-3.1.5.tar.gz \
  && tar xzf eudev-3.1.5.tar.gz \
  && rm eudev-3.1.5.tar.gz \
  ;

RUN cd eudev-3.1.5 \
  && libtoolize --force \
  && aclocal \
  && autoheader \
  && automake --force-missing --add-missing \
  && autoconf \
  && ./configure --host=arm-linux-gnueabi --target=arm-linux-gnueabi --prefix=/usr/arm-linux-gnueabi \
  && make -j \
  && make install \
  ;

RUN curl -fsSL http://old.openzwave.com/snapshots/openzwave-1.4.2143.tar.gz -o openzwave-1.4.2143.tar.gz \
  && tar xzf openzwave-1.4.2143.tar.gz \
  && rm openzwave-1.4.2143.tar.gz \
  ;

RUN cd openzwave-1.4.2143 \
  && make -j ARCH=arm CROSS_COMPILE=arm-linux-gnueabi- \
  && cp libopenzwave.so* /usr/arm-linux-gnueabi/lib \
  ;

# GO ENV

ENV GOOS=linux
ENV GOARCH=arm
ENV CGO_ENABLED=1
ENV CC=arm-linux-gnueabi-gcc
ENV GXX=arm-linux-gnueabi-g++

WORKDIR /go
