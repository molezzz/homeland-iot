<template>
  <div class="grid">
    <a href="javascript:;" class="item">
      <i class="iconfont icon-scan"></i>
      <span>添加设备</span>
    </a>
    <router-link to="/devices" class="item">
      <i class="iconfont icon-device"></i>
      <span>我的设备</span>
    </router-link>
  </div>
</template>

<style lang="sass" scoped>
.grid
  margin-top: 2rem;
  display: flex;
  .item
    display: flex;
    flex-direction: column;
    flex-grow: 1;
    text-align: center;
    .iconfont
      font-size: 56px;
      margin-bottom: 0.5rem;
</style>


<script>
export default {
  data(){
    return {}
  }
}
</script>
