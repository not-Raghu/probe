import {Queue} from 'bullmq'

const myQueue = new Queue('email-service')

async function main() {
    await myQueue.add("r1" , {
        name: "raghu1"
    }) 
    await myQueue.add("r2" , {
        name: "raghu2",
        "bday": "july16"
    })
}

await main();