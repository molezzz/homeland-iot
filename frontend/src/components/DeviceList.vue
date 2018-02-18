<template>
  <div class="device-list">
    <mt-cell :title="item.name" :label="item.uuid" is-link :to="'/devices/' + item.id" v-for="item in items"></mt-cell>
  </div>
</template>

<style lang="sass" scoped>
.device-list
  .iconfont
    font-size: 16px;
    color: #666;
</style>


<script>
export default {
  data(){
    return {
        items: []
    }
  },
  created(){
    this.loadData();
  },
  methods: {
    loadData(){
      this.$http.get('/api/equipments')
        .then((r) => {
          this.items = r.data;
        })
    }
  }
}
</script>
