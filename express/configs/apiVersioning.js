export const urlVersioning = (version) => (req,res,next) =>{
    if(req.path.startsWith(`/api/${version}`)){
        next();
    }else{
        res.status(400).json({
            success: false,
            error: "invalid api versioning"
        })
    }
} 