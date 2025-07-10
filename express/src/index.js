const express = require('express')
const route = require('../router/itemsRouter')
const app = express();

app.use(express.json())
app.use('/api' , route.pushRoute)




app.listen(3000,()=>{
    console.log("running on 3000")
})