<template>
  <div id="app">
    <table></table><!-- 防止外边距溢出 -->
    <div class="icon"></div>
    <div class="term">
        <span id="term1">第八季第一期</span>
    </div>
    <h3></h3>
    <form action="/youth/lesson/enroll" method="post">
        <div class="tip">请确认您的个人信息</div>
        <div class="confirm-user-info">
            <p><span>当前课程：</span><span id="term2">{{ course }}</span></p>
            <p><span>您的姓名：</span><span id="name">{{ name }}</span></p>
            <p><span>所在单位：</span><span id="unit">{{ org }}</span></p>
        </div>

        <div class="buttons">
            <div class="prev">修改信息</div>
            <div class="next" @click="nextStep()">确定</div>
            <div class="clear"></div>
        </div>

        <div class="group"></div>
    </form>
</div>
</template>

<script>
import {formatSection, formatNum, numberTranToCN, setMetaTitle} from '../js/util.js'

export default {
  data () {
    return {
      name: '',
      org: '',
      course: ''
    }
  },

  created() {
    let data = this.$route.params.form_data
    this.name = data.input_name
    this.org = data.input_organization

    let name = data.termName[0]
    if (name.slice(0,5) == "青年大学习") {
      name = name.slice(5, name.length)
    } else if (name.slice(0,7) == "“青年大学习”"){
      name = name.slice(7, name.length)
    }

    this.course = name
    setMetaTitle('我正在参加江苏省“青年大学习”')
  },

  methods: {
    nextStep() {
      this.$router.push({
        name: 'end',
        params: {
          form_data: this.$route.params.form_data,
          course: this.course
        }
      })
    }
  }
}

</script>

<style>
html,
        body {
            width: 100%;
            height: 100%;
            user-select: none;
            -webkit-user-select: none;
            font-size: calc(100vw / 3.75)
        }

        #app {
            position: relative;
            height: 100%;
            background-image: url('https://service.jiangsugqt.org//youth/static/image/background.jpg');
            background-repeat: no-repeat;
            background-position: center;
            background-size: auto 100%;
            padding-bottom: 50px;
        }

        .icon {
            position: absolute;
            top: 0.2rem;
            left: 50%;
            width: 2.38rem;
            height: 1.32rem;
            transform: translateX(-50%);
            background-image: url('https://service.jiangsugqt.org//youth/static/image/icon.png');
            background-repeat: no-repeat;
            background-size: 100% 100%;
            background-position: center;
        }

        .term {
            position: absolute;
            top: 1.56rem;
            left: 50%;
            display: flex;
            align-items: flex-end;
            justify-content: center;
            width: 2.32rem;
            height: 0.78rem;
            transform: translateX(-50%);
            color: #FAEDAF;
            background-image: url('https://service.jiangsugqt.org//youth/static/image/term.png');
            background-repeat: no-repeat;
            background-size: 100% 100%;
        }

        .term span {
            font-size: 0.12rem;
            padding-bottom: 0.09rem;
            transform: scale(0.8);
        }

        h3 {
            margin-top: 2.54rem;
            text-align: center;
            font-size: 0.18rem;
            font-weight: 800;
            color: rgba(250, 237, 175, 1);
        }

        .tip {
            margin-top: 0.52rem;
            margin-left: 0.26rem;
            margin-bottom: 0;
            font-size: 0.15rem;
            font-weight: 500;
            color: rgba(255, 255, 255, 1);
        }

        .input-wrapper {
            position: relative;
            margin: 0.1rem auto 0.12rem;
            width: 3.25rem;
            height: 0.44rem;
            background: rgba(255, 255, 255, 1);
            border: 0.02rem solid rgba(255, 195, 93, 1);
            border-radius: 0.22rem;
            overflow: hidden;
        }

        .input-wrapper input {
            position: absolute;
            top: 0.08rem;
            left: 0;
            padding-left: 0.27rem;
            width: 100%;
            height: 0.24rem;
            line-height: 0.24rem;
            outline: none;
            border: none;
            font-size: 0.14rem;
        }

        .buttons {
            width: 90%;
            margin: auto;
        }

        .clear {
            clear: both;
        }

        .next, .prev {
            margin: 0 auto;
            width: 1rem;
            height: 0.44rem;
            line-height: 0.44rem;
            border: 0.02rem solid rgba(255, 195, 93, 1);
            border-radius: 0.22rem;
            font-size: 0.18rem;
            text-align: center;
            color: rgba(190, 39, 25, 1);
            background-color: #FFE165;
        }

        .prev {
            float: left;
        }

        .next {
            float: right;
        }

        .group {
            position: absolute;
            width: 100%;
            left: 50%;
            bottom: 0.12rem;
            width: 2.09rem;
            height: 0.15rem;
            transform: translateX(-50%);
            background-image: url('https://service.jiangsugqt.org//youth/static/image/group.png');
            background-repeat: no-repeat;
            background-size: 100% 100%;
            background-position: center;
        }

        .confirm-user-info {
            padding: 10px;
            background: #ffffff;
            font-size: 16px;
            line-height: 28px;
            color: #333333;
            border: 0.02rem solid rgba(255, 195, 93, 1);
            width: 90%;
            box-sizing: border-box;
            margin: 15px auto 15px;
            border-radius: 5px;
        }

        .confirm-user-info p {
            padding: 0;
        }
</style>
