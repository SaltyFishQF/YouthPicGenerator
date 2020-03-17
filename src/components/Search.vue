<template>
  <div class="app">
     <search
        auto-scroll-to-top
        @result-click="resultClick"
        @on-change="getResult"
        :results="results"
        v-model="value"
        @on-cancel="onCancel"
        @on-submit="onSubmit"
        ref="search"
    ></search>
  </div>
</template>

<script>
import {Search} from 'vux'

export default {
  components: {

    Search
  },
  data () {
    return {
      input_name:'',
      input_edition:'',
      input_term:'',
      termList: null,
      results:   [{title: 'hello'}, {title: 'hel1lo'}, {title: 'he2llo'}] ,
      value: 'test',
    }
  },

  created() {
    this.termList =  this.generateTermList()
  },

  methods: {
    onChoosedTerm() {

    },

    generateTermList() {
      let listTerm = []
      let listEdition = []
      for (var i = 1; i <= 12; i++) {
        let tmpTerm = {}  
        let tmpEdition = {} 
        tmpTerm.name = '第' + i + '季'
        tmpTerm.value = i
        tmpEdition.name = '第' + i + '期'
        tmpEdition.value = i
        listTerm[i-1] = tmpTerm
        listEdition[i-1] = tmpEdition
      }
      return [listTerm, listEdition]
    },

    setFocus () {
      this.$refs.search.setFocus()
    },
    resultClick (item) {
      window.alert('you click the result item: ' + JSON.stringify(item))
    },
    getResult (val) {
      console.log('on-change', val)
      // this.results = val ? getResult(this.value) : []
    },
    onSubmit () {
      this.$refs.search.setBlur()
      this.$vux.toast.show({
        type: 'text',
        position: 'top',
        text: 'on submit'
      })
    },
    onFocus () {
      console.log('on focus')
    },
    onCancel () {
      console.log('on cancel')
    },

  }
}


function getResult (val) {
  let rs = []
  for (let i = 0; i < 20; i++) {
    rs.push({
      title: `${val} result: ${i + 1} `,
      other: i
    })
  }
  return rs
}
</script>s

<style>
.vux-demo {
  text-align: center;
}
.logo {
  width: 100px;
  height: 100px
}
</style>
