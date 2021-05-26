RUN wget https://nodejs.org/dist/v12.3.0/node-v12.3.0-linux-x64.tar.xz
RUN apt install xz-utils
RUN tar -xvf node-v12.3.0-linux-x64.tar.xz
RUN cp -r node-v12.3.0-linux-x64/{bin,include,lib,share} /usr/