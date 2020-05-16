## 目录
- [公告](https://github.com/SaltyFishQF/YouthPicGenerator#公告)  
- [简介](https://github.com/SaltyFishQF/YouthPicGenerator#简介)  
- [使用方法](https://github.com/SaltyFishQF/YouthPicGenerator#使用方法)

# 公告
- 2020-5-15：发布v1.2版本。适配第九季。新增了一个Golang后端来爬取服务器信息。
- 2020-5-14：目前第九季修改了图片规则，不过解决方案已经制定好了，近期更新。由于大学习可能更新修改算法，建议watch本仓库，所有通知今日以后会及时发布。
- Earlyer  ：使用vue重构，jQuery版本停止维护，代码移动至jQuery分支。使用Good-Storage管理本地缓存，所有自动保存的信息都在本地，不会上传服务器。

# 简介
本项目旨在帮助仍然被强制要求收截图的同学快速生成截图。通过本项目生成截图，不会在后台留下观看记录，未来也不会实现能留下观看记录的方法，这违背了本项目的初衷。本项目采用VUE编写，考虑到移动用户访问，使用了[VUX组件](https://github.com/airyland/vux)。在1.2版本以后，由于大学习修改了相关算法，添加了一个轻量化的web服务，用于定期爬虫维护截图链接。  
本项目截图页第一页为江苏共青团青年大学习的页面，不同省份可能不一样，欢迎有能力的同学Fork本仓库提供不同省份的解决方案。同时，如果学校不再要求提交截图，本项目可能会停止维护。

# 使用方法
> 注意：若版本更新后，网页未及时更新，需要点击右上角“...”图标，并选择刷新。
1. 建议在微信浏览器中打开，可以达到以假乱真的效果。[Demo部署地址](youth.logan-qiu.cn)
2. 填写姓名、学校信息，并选择要生成截图的期号。(个人信息只需填写一次，会保存在浏览器缓存中。如果出现输入框不能输入的问题，可能是VUX组件有bug，双击输入框就能输入了)
3. 点击确定，截图!
<div style="display: block;">
  <img width="30%" src="http://youth.logan-qiu.cn/preview/step1.jpg">
  <img width="30%" src="http://youth.logan-qiu.cn/preview/step2.jpg">
  <img width="30%" src="http://youth.logan-qiu.cn/preview/step3.jpg">
</div>
