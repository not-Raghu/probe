import { client } from "./client.js";

async function init() {
  //   await client.set("user:4", "raghu4");
  //   await client.expire("user:4", 20);
  //   // const result = await client.get("user:1")
  // console.log()
  //   await client.set("key:1", "hi")

  //   while (1) {
  //    await client.get("key:1", (err, result) => {
  //         console.log(result)
  //     });
  //   }
  client.set("math", "data", "EX", 60 , ()=>{
    console.log("data added")
  });
}

init();



