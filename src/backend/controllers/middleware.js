const shortUrl = require('../models/shortUrl');
const redis = require('../services/redis');
const { generateShortUrl } = require('../utils/base62encoder');

const CACHE_EXPIRATION_TIME = 30;

async function handleShortenUrl(req, res) {
    try {
        const { fullUrl } = req.body;

        if (!fullUrl) {
            return res.status(400).json({ error: 'The "fullUrl" parameter is required.' });
        }

        const existingUrl = await shortUrl.findOne({ full: fullUrl });

        if (existingUrl) {
            return res.status(200).json({
                message: 'Url already exists',
                url: existingUrl,
            });
        }

        const newUrl = new shortUrl({
            full: fullUrl,
            short: generateShortUrl(),
        });

        const savedShortUrl = await newUrl.save();

        res.status(201).json(savedShortUrl);
    } catch (error) {
        console.error('Error saving or retrieving URL: ', error);
        res.status(500).json({ error: 'Internal Server Error' });
    }
}

async function handleGetAllUrls(req, res) {
    try {
        const urls = await shortUrl.find();

        res.status(200).json({
            urls: urls,
        });
    } catch (err) {
        console.error('Error retrieving urls: ', err);
        res.status(500).json({ error: 'Internal server error' });
    }
}

async function handleGetFullUrl(req, res) {
    try {
        const { fullUrl } = req.body;

        const checkIfUrlExists = await shortUrl.findOne({ full: fullUrl });

        if (checkIfUrlExists) {
            res.status(200).json(checkIfUrlExists);
        } else {
            res.status(404).json({ message: 'Url not found' });
        }
    } catch (error) {
        console.error('Error retrieving full URL: ', error);
        res.status(500).json({ error: 'Internal Server Error' });
    }
}

async function handleRedirection(short, res) {
    const urlData = await shortUrl.findOne({ short });

    if (urlData) {
        urlData.clicks++;
        await urlData.save();
        res.redirect(urlData.full);
    } else {
        res.status(404).json({ message: 'Url not found' });
    }
}

async function handleCacheAndRedirection(short, res) {
    const cachedResult = await redis.get(short);

    if (cachedResult) {
        await handleRedirection(short, res);
    } else {
        const urlData = await shortUrl.findOne({ short });

        if (!urlData) {
            res.status(404).json({ message: 'Url not found' });
            return;
        }

        await redis.set(urlData.short, urlData.full, {
            EX: CACHE_EXPIRATION_TIME,
            NX: true,
        });

        await handleRedirection(short, res);
    }
}

module.exports = {
    handleGetAllUrls,
    handleShortenUrl,
    handleGetFullUrl,
    handleCacheAndRedirection,
};
