# Load the Docker image
FROM ruby:3.2.2-alpine3.18

# Define environment variables
ENV NAME rails_backend
ENV LANG C.UTF-8
ENV ROOT /server

# Install Packages
RUN apk update && apk upgrade && apk add --update --no-cache \
      build-base \
      curl-dev \
      linux-headers \
      libxml2-dev \
      libxslt-dev \
      postgresql-dev \
      ruby-dev \
      yaml-dev \
      graphviz \
      zlib-dev \
      bash \
      git \
      openssh \
      postgresql \
      ruby-json \
      tzdata \
      vim \
      yaml \
      curl \
      imagemagick \
      libc6-compat && ln -s /lib/libc.musl-x86_64.so.1 /lib/ld-linux-x86-64.so.2

# Set up a work directory
RUN mkdir "$ROOT"
WORKDIR $ROOT

# Install Gems
ADD Gemfile .
ADD Gemfile.lock .
RUN gem install bundler && bundle install --jobs 20 --retry 5

# Copy the source code
ADD . $ROOT
