# Prompt

## 说明

一个可以输入内容并返回的prompt组件, 支持一个默认slot放入中间.

## 用法

**父组件**

template中
```html
<prompt :visible.sync="promptVisible" :placeholder="输入店号" :defaultValue="123" @confirm="clickPromptConfirm" mainColor="#e74a39">
  <!-- 这里放入slot内容-->
</prompt>
```

```js
import Prompt from '@/components/prompt/index.vue'

export default {
  data() {
    return {
      // 控制弹框输入框显示
      promptVisible: false,
    }
  },
  methods: {
    /**
     * 点击弹出输入框确定
     */
    clickPromptConfirm(val) {
      console.log(val)
    },
  }
}
```