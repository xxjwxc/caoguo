/* 用户信息相关 */
import Server from './def' 

const product = {}

product.GetAdInfo =async getaAdInfo=>{
	let resp = await Server.Post("/product.get_ad_info",{});
	return Server.GetReturn(resp);
};

product.GetBoutiqueGroup = async getBoutiqueGroup =>{
	let resp = await Server.Post("/product.get_boutique_group",{});
	return Server.GetReturn(resp);
};

product.GetProductDetails = async function(gpid){
	let resp = await Server.Post("/product.get_product_details",{gpid:gpid});
	return Server.GetReturn(resp);
};

product.ToFavorite =async function(gpid,isFavorite){
	let resp = await Server.Post("/product.favorite",{gpid:gpid,isFavorite:isFavorite});
	return Server.GetReturn(resp);
};

// 加入购物车
product.AddToCart =async function(gpid,number,sku){
	let resp = await Server.Post("/product.add_to_cart",{gpid:gpid,number:number,skuList:sku});
	return Server.GetReturn(resp);
};

product.GetCartList = async getCartList =>{
	let resp = await Server.Post("/product.get_cart_list",{});
	return Server.GetReturn(resp);
};

// 修改数量或者删除
product.ChangeCart = async function(id,number){
	let resp = await Server.Post("/product.change_cart",{id:id,number:number});
	return Server.GetReturn(resp);
};

// 直接购买时2添加到虚拟购物车
product.AddToBuyTmpCart = async function(gpid,number,sku){
	let resp = await Server.Post("/product.add_to_buy_tmp_cart",{gpid:gpid,number:number,skuList:sku});
	return Server.GetReturn(resp);
}

// 获取地址
product.GetAddress = async function(isDefault){
	let resp = await Server.Post("/product.get_address",{isDefault:isDefault});
	return Server.GetReturn(resp);
}

// 添加&修改&删除地址
product.AddAddress = async function(addr,op){
	let resp = await Server.Post("/product.add_address",{addr:addr,op:op});
	return Server.GetReturn(resp);
}

// 通过类型获取商品列表
product.GetProductByType = async function(typeId,pageNumber){
	let resp = await Server.Post("/product.get_product_by_type",{typeId:typeId,pageNumber:pageNumber});
	return Server.GetReturn(resp);
}

// 获取收藏列表
product.GetFavorite = async function(){
	let resp = await Server.Post("/product.get_favorite",{});
	return Server.GetReturn(resp);
}



export default product