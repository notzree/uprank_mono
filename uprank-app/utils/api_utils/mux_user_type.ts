
export default  function mux_user_type(type: string){
    if (type === "freelancer") {
        return "freelancer";
    }
    else if (type === "client") {
        return "client";
    }
    else {
        return "unset";
    }
}