const express = require("express");
const cors = require("cors");
const bodyParser = require("body-parser");
const cryptojs = require("crypto-js");
const e = require("express");

const app = express();
app.use(cors());
app.use(express.json());

function encrypt(data, key) {
  const ecryption = cryptojs.AES.encrypt(data, key).toString();

  return ecryption;
}

app.post("/encrypt", (req, res) => {
  const { data, key } = req.body;

  if (!data || !key) {
    return res.status(400).json({
      msg: "(no key | data ) & both",
    });
  }

  const response = encrypt(data, key);
  return res.status(200).json({
    response,
  });
});
module.exports = app;

app.get('/decrypt',(req,res)=>{
    const {encrypted,key} = req.body;
    const data = cryptojs.AES.decrypt(encrypt,key);

    return res.status(200).json({
        data
    })
})