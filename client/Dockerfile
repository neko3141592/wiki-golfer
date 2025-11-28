# 公式Nodeイメージを使用
FROM node:20-alpine AS builder

WORKDIR /app

# package.jsonのみコピー
COPY package.json .

# 依存インストール
RUN npm install --frozen-lockfile || yarn install --frozen-lockfile

# ソースコードをコピー
COPY . .

# Next.jsのビルド
RUN npm run build || yarn build

# 本番用イメージ
FROM node:20-alpine AS runner
WORKDIR /app

ENV NODE_ENV=production

COPY --from=builder /app .

EXPOSE 3000

CMD ["npm", "start"]
