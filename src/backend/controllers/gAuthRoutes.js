const express = require('express');
const router = express.Router();

router.get("/auth/google",
    passport.authenticate("google", { scope: ["profile"] })
);

router.get("/auth/google/callback",
    passport.authenticate("google", { failureRedirect: "http://localhost:3000" }),
    function (req, res) {
        res.status(200).json({ message: "Authentication Successful" });
    }
);

router.get("/logout", function (req, res) {
    res.status(200).json({ message: "Logout Successful" });
});


module.exports = router;