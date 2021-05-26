FROM golang:1.16
RUN apt-get update
RUN apt-get install -yq git python jq curl
RUN curl -sL https://deb.nodesource.com/setup_14.x | bash -
RUN apt-get update && apt-get install -yq nodejs
RUN npm install npm -g
RUN npm install -g n
RUN n 12.3.0

COPY . /code-analyser
WORKDIR /code-analyser

RUN bash installNodeModulesPlugin.sh
RUN bash compileGoPluginBinaries.sh

CMD ["bash", "run.sh"]