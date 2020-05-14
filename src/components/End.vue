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
    let data = this.$route.params.form_data
    this.bg = data.termUrl + "/images/end.jpg"
    setMetaTitle(data.termName)
  },

  methods: {

  }
}

</script>

<style>

</style>
