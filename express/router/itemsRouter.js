const express = require("express");

const router = express.Router();
const items = [1, 2, 3, 4];
const pushRoute = router.post("/push", (req, res) => {
  const { number } = req.body;
  const num = parseInt(number);
  items.push(num);
  res.json(items);
});


module.exports = {
  pushRoute
}