<template>
  <div id="pic" style="height: 100%;width: 100%;background-size: 100% 100%;background-repeat: no-repeat"  v-bind:style="{backgroundImage:'url(' + bg + ')'}"></div>
</template>

<script>
import {setMetaTitle} from '../js/util.js'
export default {
  
  data () {
    return {
      bg: ''
    }
  },

  created() {
    let convert = ['0', 'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p', 'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l'];

    let t = this.$route.params.term

    this.bg = "http://h5.cyol.com/special/daxuexi/daxuexi" + t[0] + convert[t[1]] + t[1] + "/images/end.jpg"
    setMetaTitle('“青年大学习” ' + this.$route.params.course)
  },

  methods: {
   
  }
}

</script>

<style>

</style>
