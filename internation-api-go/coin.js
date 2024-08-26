//从coinmarketcap获取数据
const getInfo=(temp1)=>{
    let all = []
    Array.from(temp1.children).map((item)=>{
        let code = item.children[2].querySelector("p.iqdbQL").textContent;
        let name = item.children[2].querySelector("p.kKpPOn").textContent;
        let logo = item.children[2].querySelector("img").src;
        all.push({
            code,
            name,
            logo
        })
    })
    console.log(all)
}