<template lang="pug">
  #app.g-content
    .m-header
      m-header
    .g-bd.f-beautify-scrollbar
      .g-warpper
        router-view
      .g-ft
        .f-fc-ccc.f-fs-14.f-tac.f-m-b-10 © 2013-{{ new Date().getFullYear() }} 上海弘连网络科技有限公司
</template>

<script>
import header from './components/header'
export default {
  name: 'App',
  components: {
    'm-header': header
  }
}
</script>

<style lang='scss'>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  width: 100vw;
  height: 100vh;
  display: flex;
  flex-direction: column;
}
.m-header{
  height: 64px;
  flex: none;
}
.g-bd{
  flex: auto;
  overflow: auto;
  padding: 0 20px;
  height: 100%;
}
.g-warpper{
  min-height: calc(100% - 70px);
}
</style>
