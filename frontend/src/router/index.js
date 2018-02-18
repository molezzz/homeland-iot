import Vue from 'vue'
import Router from 'vue-router'


import FrontPage from '@/components/FrontPage.vue'
import DeviceList from '@/components/DeviceList.vue'
import DeviceShow from '@/components/DeviceShow.vue'
import AdminAddDevice from '@/components/AdminAddDevice.vue'
import 'mint-ui/lib/style.css';

Vue.use(Router)


export default new Router({
    routes: [{
        path: '/',
        name: 'frontpage',
        component: FrontPage
    }, {
        path: '/devices',
        name: 'device_list',
        component: DeviceList
    }, {
        path: '/devices/:id',
        name: 'device_show',
        component: DeviceShow
    }, {
        path: '/admin/add/device',
        name: 'admin_add_device',
        component: AdminAddDevice
    }]
})