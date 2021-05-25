FROM golang:1.16
RUN apt-get update
RUN apt-get install -yq git python jq curl

RUN curl -sL https://deb.nodesource.com/setup_12.x | bash -
RUN apt-get update && apt-get install -yq nodejs
RUN npm install npm -g

COPY . /app
WORKDIR /app

RUN bash installNodeModulesPlugin.sh
RUN bash compileGoPluginBinaries.sh

CMD ["bash", "run.sh"]