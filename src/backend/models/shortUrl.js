const mongoose = require('mongoose');

const shortUrlSchema = new mongoose.Schema({
    full: {
        type: String,
        required: true
    },
    short: {
        type: String,
        required: true
    },
    clicks: {
        type: Number,
        required: true,
        default: 0
    },
    createdAt: {
        type: Date,
        default: () => Date.now(),
        immutable: true,
        required: true,
    }
});

module.exports = mongoose.model('shortUrl', shortUrlSchema);