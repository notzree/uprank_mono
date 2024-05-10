export default function craft_api_url(path: string){
    //wrap every API request with this. We need the full path because middleware

    const base_api_url = process.env.NODE_ENV === "production" ? process.env.NEXT_PUBLIC_PROD_BASE_URL : process.env.NEXT_PUBLIC_LOCAL_BASE_URL;
    return base_api_url + path;
}