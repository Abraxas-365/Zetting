import axios from 'axios';


export const apiCalls = axios.create({
    baseURL: 'http://192.168.1.42:3000'
    // baseURL: 'http://172.20.10.11:3000'
});

export const serveDefaultImages = "http://192.168.1.42:3000/static/app_default_images"
