# Build Stage
FROM node:20-alpine AS builder

WORKDIR /app

# Copy dependencies
COPY package.json package-lock.json ./
RUN npm install

# Copy source
COPY . .

# Build the application
RUN npm run build

# Production Stage with Nginx
FROM nginx:alpine

# Copy built assets
COPY --from=builder /app/dist /usr/share/nginx/html

# Expose port
EXPOSE 80

# Start Nginx
CMD ["nginx", "-g", "daemon off;"]
