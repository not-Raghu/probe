import {Worker} from 'bullmq'

const worker = new Worker('email-service',async(job)=>{
    console.log(job.data)
}).run()

