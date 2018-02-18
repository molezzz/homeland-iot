import axios from 'axios'
import qs from 'qs'

var instance = axios.create({
    baseURL: '', //'//localhost:8080',
    timeout: 50000,
    headers: {
        // post: {
        //     'Content-Type': 'application/x-www-form-urlencoded'
        // },
        // put: {
        //     'Content-Type': 'application/x-www-form-urlencoded'
        // }
    },
    withCredentials: true,
    paramsSerializer: function(params) {
        return qs.stringify(params, { arrayFormat: 'brackets' })
    }
});
instance.interceptors.response.use(function(response) {

    return response;
}, function(error) {
    console.info(error.response.data);
    console.info(error.response.status);
    if (error.response) {
        if (error.response.status >= 500) {
            alert('服务器开小差了(code:' + error.response.status + ')，稍后再试吧');
        }
        /* else if(error.response.status >= 400){
                   alert(error.response.data.message);
               } */
    }
    return Promise.reject(error);
});
export default instance;