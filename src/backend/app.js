const express = require('express');
const morgan = require('morgan');
const bodyParser = require('body-parser');
const mongoose = require('./services/database');
const redis = require('./services/redis');
const dotenv = require("dotenv");

dotenv.config();

mongoose.connection.once('open', () => {
    console.log('Connected to MongoDB');
});

// Handle any errors that occur during the database connection
mongoose.connection.on('error', (err) => {
    console.error('MongoDB connection error:', err);
});

redis.on("connect", function () {
    console.log("Redis Connection Successful!!");
});

app = express();

app.use(bodyParser.urlencoded({ extended: false }));
app.use(bodyParser.json());


app.use(morgan('dev'));

const urlRoutes = require('./controllers/routes');
app.use('/', urlRoutes);

app.use((req, res, next) => {
    const error = new Error('Not found');
    error.status(404);
    next(error);
});

app.use((error, req, res, next) => {
    res.status(error.status || 500);
    res.json({ // 200: OK
        error: {
            message: error.message
        }
    });
});

module.exports = app;