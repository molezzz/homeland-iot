<template>
<div>
  <div class="device-show" v-if="device">
    <i class="iconfont icon-thermograph"></i>
    {{device.name}}
  </div>
  <mt-button type="default" size="large" v-if="device && device.member && device.member.id == currentMember.id" @click.prevent="showDevice">您已绑定该设备，直接查看</mt-button>
  <mt-button type="danger" size="large" v-if="device && device.member && device.member.id != currentMember.id">设备已被别人绑定了</mt-button>
  <mt-button type="primary" size="large" v-if="device && device.member == null" @click.prevent="bind">绑定设备</mt-button>
</div>
</template>

<style lang="sass" scoped>
.device-show
    display: flex;
    flex-direction: column;
    text-align: center;
    color: #666;
    .iconfont
      margin: 20% auto 1.5rem;
      font-size: 64px;
.mint-button
  margin-top: 1.5rem;
</style>


<script>
export default {
  data(){
    return {
      currentMember: null,
      uuid: null,
      device: null,
      records: []
    }
  },
  created(){
    console.log(this.$route.params)
    this.uuid = this.$route.params.id
    this.loadDevice()
    this.currentMember = JSON.parse(window.signInMember)
    
  },
  methods: {
    loadDevice(){
      if(!this.uuid) { 
        alert('请提供设备UUID')
        return 
      }
      this.$http.get('/api/equipments/' + this.uuid)
        .then((r) => {
          this.device = r.data
        }, (e) => {
          alert(e.response.data.message)
        })
    },
    bind(){
      this.$http.put('/api/equipments', {
        uuid: this.uuid
      }).then((r) => {
        this.device.member = r.data.member;
      },(e) => {
        alert(e.response.data.message);
      })
    },
    showDevice(){
      this.$router.replace({name: 'device_show', params:{id: this.uuid}})
    }
  }
}
</script>
