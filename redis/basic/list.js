import client from "./client.js";

async function init() {
  const users = await client.get("users",(err,res)=>{
   console.log(res) 
  });
}
init()