exports.generatePass = function(limit=4) {
    return Math.random().toString(36).replace(/[^a-z]+/g, '').substr(0, limit);
}
