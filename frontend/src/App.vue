<template>
  <div id="app">
    <mt-header fixed title="homeland">
      <router-link to="/" slot="left">
        <mt-button icon="back">返回</mt-button>
      </router-link>
      <mt-button icon="more" slot="right" @click.prevent="showAdminMenu"></mt-button>
    </mt-header>
    <router-view/>
    <mt-actionsheet :actions="actions" v-model="menuShow"></mt-actionsheet>
  </div>
</template>

<script>
export default {
  name: 'app'
}
</script>

<style>
#app {
  padding-top: 40px;
}
a {
  color: #666;
  text-decoration: none;
}
</style>


<script>
export default {
  data(){
    return {
      menuShow: false,
      actions: [{
        name: '添加新设备',
        method: () => {
          this.$router.push({name: 'admin_add_device'})
        }
      }]
    }
  },
  methods: {
    showAdminMenu(){
      this.menuShow = true;
    }
  }
}
</script>
