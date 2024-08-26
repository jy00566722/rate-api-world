# 基于Node.js官方镜像，选择合适的版本
FROM node:lts-alpine as build-stage

# 设置工作目录
WORKDIR /app

# 复制`package.json`和`yarn.lock`到工作目录
COPY package*.json ./
COPY yarn.lock ./

# 安装项目依赖
RUN npm install

# 复制项目文件到工作目录
COPY . .

# 构建项目
RUN npm run build

# 生产阶段使用Nginx
FROM nginx:stable-alpine as production-stage
COPY --from=build-stage /app/dist /usr/share/nginx/html

# 暴露80端口
EXPOSE 80

# 用Nginx启动服务
CMD ["nginx", "-g", "daemon off;"]
