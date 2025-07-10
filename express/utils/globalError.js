class APIError extends Error{
    constructor(message, statusCode){
        super(message)
        this.statusCode = statusCode;
    }

}


const asyncHandler = (fn) => (req,res,next)=>{
    Promise.resolve(fn(req,res,next)).catch(next);
}

const globalErrorHandler = (err,req,res,next) => {
    console.log(err.stack)

    if(err instanceof APIError){
        return res.status(err.status).json({
            status: "error",
            message: err.message
        })
    }

}