<template>
	<logistics :wlInfo="wlInfo"></logistics>
</template>

<script>
	import logistics from '@/components/xinyu-logistics/xinyu-logistics.vue'
	export default {
		data() {
			return {
				billId:"",
				wlInfo: {}
			}
		},
		onLoad(options){
			this.billId = options.billId;
			this.onInit();
		},
		components: {
			logistics
		},
		methods: {
			async onInit(){
				let resp = await this.$Order.GetOrdertrackInfo(this.billId)
				if (typeof(resp) != "undefined" && resp != null){
					this.wlInfo = resp;
				}else{ // 未找到
					uni.showToast({
						title: "暂无快递信息",
						icon: 'none'
					})	
					setTimeout(async ()=>{
						uni.navigateBack()
					}, 1500);	    
					
				}
			}
		}
	}
</script>

<style>

</style>
