<template lang='pug'>
  div
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  created: function () { // 页面加载时跳转
    if (process.env.NODE_ENV === 'development') {
      location.href = `http://127.0.0.1:8000/auth/honglian`
    } else {
      location.href = `/auth/honglian`
    }
  }
})
</script>
