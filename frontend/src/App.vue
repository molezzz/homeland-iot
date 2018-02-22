<template>
  <div id="app">
    <mt-header fixed title="homeland">
      <mt-button icon="back" slot="left" @click.prevent="goBack">返回</mt-button>
      <mt-button v-if="currentMember && currentMember.isAdmin" icon="more" slot="right" @click.prevent="showAdminMenu"></mt-button>
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
      currentMember: null,
      actions: [{
        name: '添加新设备',
        method: () => {
          this.$router.push({name: 'admin_add_device'})
        }
      }]
    }
  },
  created(){
    this.currentMember = JSON.parse(window.signInMember)
  },
  methods: {
    showAdminMenu(){
      this.menuShow = true
    },
    goBack(){
      this.$router.go(-1)
    }
  }
}
</script>
