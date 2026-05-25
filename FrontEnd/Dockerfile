FROM node:22-alpine

WORKDIR /usr/src/app

COPY package*.json ./
RUN npm install

COPY . .

EXPOSE 4321

# Gunakan polling agar hot reload berfungsi di dalam container
ENV CHOKIDAR_USEPOLLING=true
CMD ["npm", "run", "dev", "--", "--host", "0.0.0.0"]
