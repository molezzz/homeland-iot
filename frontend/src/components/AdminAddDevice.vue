<template>
<div>
  <mt-field label="设备名" placeholder="请输入设备名" v-model="name"></mt-field>
  <mt-field label="UUID" placeholder="请输入设备UUID" type="email" v-model="uuid"></mt-field>
  <mt-button type="primary" :disabled="!name || !uuid" size="large" style="margin-top: 2rem" @click="addDevice">保存</mt-button>
</div>
</template>

<script>
export default {
  data(){

    return {
      name: null,
      uuid: null
    }
  },
  methods: {
    addDevice(){
      if(!this.name || !this.uuid) { return }
     
      this.$http.post('/api/equipments', {
        name: this.name,
        uuid: this.uuid
      }).then((r) => {
        alert('保存成功');
        this.$router.replace({name: 'frontpage'})
      }, (e) => {
        alert(e.response.data.message);
      })
    }
  }
}
</script>
