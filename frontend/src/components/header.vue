<template lang='pug'>
  .m-header.f-df
    .title.f-fs-20.f-fwb TITLE
    .menu
      el-menu.f-ib.f-vam(:default-active="$route.path", mode="horizontal", router=true)
        el-menu-item.f-fs-16(index="/home") HOME
        el-menu-item.f-fs-16(index="/about") ABOUT
    .userInfo.f-csp
      img.avatar.f-vam.f-m-r-10(:src="user.avatar", :onerror='onerrorImg')
      span.f-fs-18.f-vam.f-m-r-10(v-text="user.email")
      img.f-vam(src="@/assets/images/icon_exit.png", @click='signout', title='登出')
</template>

<script lang="ts">
import Vue from 'vue'
import { mapActions, mapState } from 'vuex'
import { USER_SIGNOUT } from '@/store/mutation-types'
import { logout } from '@/service/user.ts'

export default Vue.extend({
  name: 'mheader',
  data () {
    return {
      avatar: require('@/assets/images/avator.png'),
      onerrorImg: "this.src='" + require('@/assets/images/avator.png') + "'"
    }
  },
  computed: mapState('user', { user: state => state }),
  methods: {
    ...mapActions('user', [USER_SIGNOUT]),
    async signout () {
      try {
        await logout()
        this.USER_SIGNOUT()
        this.$router.push({ name: 'login' })
      } catch (error) {
        console.log("error")
      }
    }
  }
})
</script>

<style lang='scss' scope>
.m-header {
  height: 64px;
  line-height: 64px;
  width: 100%;
  color: #999;
  border-bottom: 1px solid #ddd;
  background-color: #fff;
  .title{
    padding: 0 80px;
    flex: none;
  }
  .userInfo{
    flex: none;
    padding: 0 40px;
    .avatar{
      width: 40px;
    }
  }
}
</style>
