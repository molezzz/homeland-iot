<template>
  <div class="device-show">
    <i class="iconfont icon-thermograph"></i>
    <p v-if="currentRecord != null">
    当前温度: {{currentRecord.temperature}}°C <br>
    当前湿度: {{currentRecord.humidity}}%     <br>
    pm2.5: {{currentRecord.pm25}}
    </p>
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
</style>


<script>
export default {
  data(){
    return {
      uuid: null,
      device: null,
      records: []
    }
  },
  created(){
    console.log(this.$route.params);
    this.uuid = this.$route.params.id;
    this.loadDevice();
  },
  computed: {
    currentRecord(){
      if(this.records.length > 0) { 
        return this.records[this.records.length - 1]
      }
      return null
    }
  },
  methods: {
    loadDevice(){
      if(!this.uuid) { return }
      this.$http.get('/api/equipments/' + this.uuid)
        .then((r) => {
          this.device = r.data
        }, (e) => {
          alert(e.response.data.message)
        })
    },
  },
  watch: {
    device(to){
      this.$http.get('/api/records', {params: {device_id: to.id}})
        .then((r) => {
          this.records = r.data
        }, (e) => {
          alert(e.response.data.message)
        })
    }
  }
}
</script>
