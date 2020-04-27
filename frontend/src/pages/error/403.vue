<template lang='pug'>
  div
    | 403 访问被拒绝
    a(@click="back") 返回上一页
</template>

<script>
import Vue from 'vue'

export default Vue.extend({
  methods: {
    back () {
      this.$router.go(-1)
    }
  }
})
</script>
