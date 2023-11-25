# Builder layer
FROM node:18 AS builder
ARG NEXT_PUBLIC_API_BASE
ARG API_BASE
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
ENV NEXT_PUBLIC_API_BASE=$NEXT_PUBLIC_API_BASE
ENV API_BASE=$API_BASE
RUN npm run build

# Production layer
FROM node:18-alpine
WORKDIR /app
COPY --from=builder /app ./
RUN addgroup app && adduser -S -G app app
USER app
EXPOSE 3000
CMD ["npm", "start"]