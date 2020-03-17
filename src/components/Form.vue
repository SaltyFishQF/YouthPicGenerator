<template>
  <div class="app">
    <group>
       <x-input title="姓名" v-model="form_data.input_name" text-align="right"></x-input>
       <popup-picker title="期号" :data="termList" v-model="form_data.termName" @on-change='onChoosedTerm'></popup-picker>
       <x-input title="组织" v-model="form_data.input_organization" text-align="right"></x-input>
    </group>
    <x-button type="primary" @click.native="nextStep()">填好了，去截图！</x-button>
  </div>
</template>

<script>
import { XInput, Group, PopupPicker, XButton} from 'vux'
import storage from 'good-storage' 


export default {
  components: {
    Group,
    XInput,
    PopupPicker,
    XButton 
  },
  data () {
    return {
      form_data: null,
      termList: null,
    }
  },

  created() {
    this.termList = this.generateTermList()
    this.getLocalCacheData()
  },

  methods: {
    nextStep() {
      storage.set("data", this.form_data)
      this.$router.push({
        name: "info",
        params: {
          form_data: this.form_data
        }
      })
    },

    onChoosedTerm(res) {
        let term = 1
        let edition = 1

        term = parseInt(res[0].charAt(1))
        edition = parseInt(res[1].charAt(1))

        this.form_data.term = [term, edition]
        console.log(term)
        console.log(edition)
    },

    getLocalCacheData() {
      this.form_data = storage.get("data", {
        input_name:'',
        term: [1,1],
        termName: ["第1季","第1期"],
        input_organization: '',
      })
    },

    setCache() {
      storage.set("data", this.form_data)
    },

    generateTermList() {
      let listTerm = []
      let listEdition = []
      for (var i = 1; i <= 12; i++) {
        let tmpTerm = {}  
        let tmpEdition = {} 

        tmpTerm.name = '第' + i + '季'
        tmpTerm.value = '第' + i + '季'
        tmpTerm.key = i
        
        tmpEdition.name = '第' + i + '期'
        tmpEdition.value = '第' + i + '期'
        tmpEdition.key = i

        listTerm[i-1] = tmpTerm
        listEdition[i-1] = tmpEdition
      }
      return [listTerm, listEdition]
    },

  }
}

</script>

<style>
.vux-demo {
  text-align: center;
}
.logo {
  width: 100px;
  height: 100px
}
</style>
