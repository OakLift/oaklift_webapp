# Stage 1: Build the React application
FROM node:alpine as builder

WORKDIR /app

# Copy package.json and package-lock.json (if present)
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the entire project
COPY . .

# Build the project
RUN npm run build

# Stage 2: Serve the built files using a lightweight HTTP server
FROM node:alpine as staging

WORKDIR /app

RUN npm install -g serve
# Copy the built files from the builder stage
COPY --from=builder /app/ .

# Expose port 8090 to the outside world
EXPOSE 3000

# Command to run the application
CMD ["serve", "-s", "build"]
