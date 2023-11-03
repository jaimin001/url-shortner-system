const express = require('express');
const { generateShortUrl, decodeShortUrl } = require('../utils/base62encoder');
const shortUrl = require('../models/shortUrl');
const redis = require('../services/redis');
const router = express.Router();
const { handleGetAllUrls,
    handleShortenUrl,
    handleGetFullUrl,
    handleCacheAndRedirection } = require('./middleware')

// Get all urls
router.get('/all', handleGetAllUrls);

// Shorten url
router.post('/shorten', handleShortenUrl);

// Get short url from a full url
router.post('/full', handleGetFullUrl);

// Get a particular full url from a short url
router.get('/:short', async (req, res, next) => {
    try {
        const { short } = req.params;
        await handleCacheAndRedirection(short, res, redis);
    } catch (err) {
        console.error(err);
        res.status(500).json(err);
    }
});

module.exports = router;