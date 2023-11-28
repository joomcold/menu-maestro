# Load the Docker image
FROM node:20-alpine3.18

# Define environment variables
ENV NAME nuxt_frontend
ENV LANG C.UTF-8
ENV ROOT /client

# Install pnpm globally
RUN npm install -g pnpm

# Set the working directory in the container
WORKDIR $ROOT

# Copy package.json and pnpm-lock.yaml to the working directory
COPY package.json pnpm-lock.yaml ./

# Install dependencies using pnpm
RUN pnpm install
RUN pnpm i @rollup/rollup-linux-arm64-musl --force

# Copy the rest of the application code
COPY . .

# Expose the port that the app will run on
EXPOSE 8080
