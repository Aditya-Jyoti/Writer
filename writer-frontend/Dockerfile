## Building the SvelteKit app 
FROM node:23-alpine AS builder

WORKDIR /app

# Install dependencies early to cache layers
COPY package.json pnpm-lock.yaml ./
RUN npm install -g pnpm && pnpm install

COPY . .
RUN pnpm run build

## Running the SvelteKit app
FROM node:23-alpine

WORKDIR /app

# Copy built app from builder
COPY --from=builder /app/build ./build
COPY --from=builder /app/package.json ./
COPY --from=builder /app/pnpm-lock.yaml ./

RUN npm install -g pnpm && pnpm install --prod

EXPOSE 3000

CMD ["node", "build"]
