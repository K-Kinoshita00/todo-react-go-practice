FROM node:26.0.0

WORKDIR /app

COPY package.json .
COPY pnpm-lock.yaml .

RUN npm install -g corepack@latest
RUN corepack enable
RUN corepack prepare pnpm@9.15.9 --activate

RUN pnpm install

COPY . .

EXPOSE 5173

CMD ["pnpm", "run", "dev"]