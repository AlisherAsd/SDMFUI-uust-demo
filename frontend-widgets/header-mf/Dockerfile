FROM node:22.23.3-bookworm-slim

WORKDIR /app

COPY . .

RUN npm i

EXPOSE 5001

CMD ["npm", "run", "dev", "--", "--host", "0.0.0.0"]