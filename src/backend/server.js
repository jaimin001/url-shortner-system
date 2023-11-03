const http = require('http');
const app = require('./app');

const dotenv = require("dotenv");

dotenv.config({ path: "./.env" });

const server = http.createServer(app);

server.listen(process.env.PORT);
