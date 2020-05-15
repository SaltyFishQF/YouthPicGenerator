<template>
  <div class="container">
    <group>
      <x-input title="姓名" v-model="form_data.input_name" text-align="right"></x-input>
      <popup-picker ref="picker" title="期号" :data="termList" v-model="this.form_data.termName"
                    @on-change='onChoosedTerm'></popup-picker>
      <x-input title="组织" v-model="form_data.input_organization" text-align="right"></x-input>
    </group>
    <x-button type="primary" @click.native="nextStep()">填好了，去截图！</x-button>

    <footer class="footer">
      <div>
        <p>
          @ SaltyFishQF 2020
        </p>
        <p>
          本项目GitHub地址：<a href="https://github.com/SaltyFishQF/YouthPicGenerator">YouthPicGenerator</a>,
          若在使用过程中发现任何问题，请开issue说明。更新说明将写在在Github的Readme公告中。
        </p>
      </div>
    </footer>
  </div>


</template>

<script>
  import {XInput, Group, PopupPicker, XButton} from 'vux'
  import storage from 'good-storage'


  export default {
    components: {
      Group,
      XInput,
      PopupPicker,
      XButton
    },
    data() {
      return {
        form_data: {
          input_name: '',
          termUrl: [],
          termName: '',
          pickedName: '',
          input_organization: '',
        },
        termList: null,
      }
    },

    created() {
      this.getLessonList()
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
        console.log(res)
        this.form_data.termUrl = res
        this.form_data.termName = [this.$refs.picker.getNameValues()]
      },

      getLocalCacheData() {
        this.form_data = storage.get("data", {
          input_name: '',
          termUrl: ['http://h5.cyol.com/special/daxuexi/9azsoxlpx2/'],
          termName: ["“青年大学习”第九季第一期"],
          input_organization: '',
        })
      },

      setCache() {
        storage.set("data", this.form_data)
      },

      getLessonList() {
        this.$axios.get('http://127.0.0.1:8000/api/getLessonList').then(res => {
          this.termList = [res.data]
        })
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

  .footer {
    height: 70vh;
    background-color: #f5f5f5;
    border-top: 1px solid #ddd;
    padding-top: 20px;
    color: #636363;
    font-size: 16px;
  }
</style>
