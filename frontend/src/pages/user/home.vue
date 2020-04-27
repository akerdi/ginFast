<template lang="pug">
  .m-user-home

</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  name: 'userHome'
})
</script>

<style lang='scss'>
.m-user-home {
  .sections {
    padding: 40px;
  }
  .createPartition {
    position: absolute;
    width: 50px;
    height: 50px;
    bottom: 48px;
    right: 48px;
  }
}

</style>
