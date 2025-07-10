const limiter = require('express-rate-limit');

const rateLimiter = (time_ms) => limiter.rateLimit({
    windowMs : time_ms,
    limit: 20,
    message: "too many requests, please wait",
    statusCode: true,
    standardHeaders: true,
    legacyHeaders: false
})

module.exports = {
    rateLimiter
}