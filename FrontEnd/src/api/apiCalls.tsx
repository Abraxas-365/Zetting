import axios from 'axios';


export const apiCalls = axios.create({
    baseURL: 'http://192.168.1.47:3000'
    // baseURL: 'http://172.20.10.11:3000'
});
