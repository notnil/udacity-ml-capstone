FROM tensorflow/tensorflow:1.2.1

# install required binaries
RUN add-apt-repository ppa:jonathonf/ffmpeg-3 && \
    apt-get update && \
    apt install -y --no-install-recommends ffmpeg libav-tools x264 x265 imagemagick librsvg2-bin curl

# python deps
RUN pip install keras tqdm h5py

# install go
ENV GOLANG_VERSION 1.6
ENV GOLANG_DOWNLOAD_URL https://golang.org/dl/go$GOLANG_VERSION.linux-amd64.tar.gz
ENV GOLANG_DOWNLOAD_SHA256 5470eac05d273c74ff8bac7bef5bad0b5abbd1c4052efbdbc8db45332e836b0b

RUN curl -fsSL "$GOLANG_DOWNLOAD_URL" -o golang.tar.gz \
	&& echo "$GOLANG_DOWNLOAD_SHA256  golang.tar.gz" | sha256sum -c - \
	&& tar -C /usr/local -xzf golang.tar.gz \
	&& rm golang.tar.gz

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

# install glide
RUN curl https://glide.sh/get | sh

# Add files from repo
ADD . /notebooks