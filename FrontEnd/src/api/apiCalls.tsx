import axios from 'axios';


export const apiCalls = axios.create({
    baseURL: 'http://192.168.1.46:3000'
});
