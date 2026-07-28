FROM node:22-alpine AS build

WORKDIR /app

COPY package*.json ./
RUN npm ci

COPY . .

ARG PUBLIC_API_URL
ENV PUBLIC_API_URL=${PUBLIC_API_URL}

RUN npm run build

FROM nginx:1.27-alpine

COPY --from=build /app/dist /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]

# Gunakan polling agar hot reload berfungsi di dalam container
# ENV CHOKIDAR_USEPOLLING=true
# CMD ["npm", "run", "dev", "--", "--host", "0.0.0.0"]
