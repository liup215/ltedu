# Build stage
FROM node:latest AS builder

# Set working directory
WORKDIR /app


# Copy package files
COPY package.json package-lock.json ./

# Install dependencies
RUN npm install

# Copy all files
COPY . .

# Build the application
RUN npm run build

# # Production stage
FROM nginx:latest

# Copy the built files to nginx
COPY --from=builder /app/dist/ /usr/share/nginx/html

# Copy entrypoint script
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

# Copy nginx configuration if needed
# COPY nginx.conf /etc/nginx/conf.d/default.conf

# Expose port
EXPOSE 80

# Start with entrypoint (generates config.json and starts nginx)
ENTRYPOINT ["/entrypoint.sh"]
