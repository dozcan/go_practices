**********************************************************************************************
algoritma şu şekilde çalışmaktadır
1)dosya aramada optimizasyon onemli olduğu için aranan cümle uzunluğu kadar memory alana chunk veri alınır
2)eğer chunk veri aranan ile uyumlu ise sonuç başarılıdır sonuç bulunmuştur
3)chunk veri stream olduğu için ve devamlı geldiği için bunu bir temp alanda biriktirmek gerekir
4)algoritma ise, aranını bulamama durumunda chunk içindeki ilk kelimeyi eler ve sonrası iteratif şekilde
  ikinci elemandan sonraki chunk verinin içindeki belli elemanı alarak aranan cümle uzunluğunda cümle oluşturur
  arama bu şekilde devam ettirilir
5) dosya bitirildiğinde aranan cümlenin hepsi var ise sonuç döndürülür ancak aranan bütünüyle yok ise
   cümle dosyada hangi oranda doğru şekilde mevcut ise onun oranı döndürülür.
6) diyelim ki dosyada şu şekilde bir cümle mevcut
   => "collection of well-known  quotations         that are associated with "
   aranan cümle ise aşağıdaki gibi olsun
   =>"well-known   quotations   that associated"
   başarı oranı %75 olacaktır
7) eğer daha büyük bir oran var ise o sonuç olarak yansıyacaktır
   => "collection of well-known  quotations         that are associated with "
   => "collection of well-known  quotations that associated with "
   aranan cümle ise aşağıdaki gibi olsun
   =>"well-known   quotations   that associated"
   başarı oranı %100 olacaktır


const fs = require('fs');
const crypto = require('crypto');
const stream = require('stream');

const folderPath = process.argv[2]
const searchPattern = process.argv[3]

fs.readdir(folderPath,(err, files) => {
  if(err){
    console.log("we got error",err)
    return
  } 
  files.forEach(file => {
     startSearchEngine(folderPath,file,searchPattern)
  });
});

var startSearchEngine = (folderPath,file,searchPattern)  => {
  let hashOfPattern = getSearchPatternHash(searchPattern)
  let lengthOfPattern = Buffer.from(searchPattern).length
  console.log(hashOfPattern)
 
  let previousChunk = ""
  const readabaleContents = fs.createReadStream(folderPath+"/"+file,{ highWaterMark: lengthOfPattern*2});
  readabaleContents.on('data', (chunk)=> {

    if(chunk.toString().split(' ').includes(searchPattern.split(" ")[0])){
      chunk.toString().indexOf(searchPattern.split(" ")[0])


    }

    let hashChunk = getSearchPatternHash(chunk.toString())

    if(hashOfPattern === hashChunk){
      console.log("we find it")
      return
    }
    else{
      previousChunk = chunk

    }
  });
}

var getSearchPatternHash = searchPattern => {
  return sha256Hash(searchPattern,"pattern with length" + searchPattern.split(' ').length)
}


var sha256Hash = function (str,hashKey){  
  var hashedStr = crypto.createHmac('sha256',str).update(hashKey).digest('hex');
  return hashedStr;
}





var s = "Barry Manilow may claim to write the songs, but it was "    
var s1 ="William Shakespeare who coined the phrases - he contributed more "
var s2 ="phrases and sayings  to the language than any other "  
var s3 ="individual, and  sayings to  the English language in daily use. Here's a " 
var s4 ="collection of well-known quotations that are associated with " 
var s5 ="Shakespeare. Most of these were the Bard's own work but he wasn't "
var s6 ="the well-known quotations is that associated for insane "


s = s+s1+s2+s3+s4+s5+s6
var pattern = "well-known quotations are associated"
let chunk=""
let temp= ""
let we_find_it = false
let moduloResultArr = []  

findIncremental = arg => {
  let incrementalCount = 0
  let firstIndex = arg[0]
  let max =[]
  
    for(let i=1;i<arg.length;i++){
       if(arg[i] - firstIndex === 1){
         incrementalCount++     
         firstIndex = arg[i]
       }
       else{
         max.push(incrementalCount)
         incementalCount=0
         firstIndex = arg[i]
       }
  }
  
  return Math.max(...max)
}

parseMultiSpace = str => {
  let s = str.trim().split(' ')
  return s.reduce((prev,next) => {
      if(next === "")
         return prev
      else return prev.concat(" ",next)
  })
}

isFilter = (source,destination)=> {
   let _destination = destination.split(' ')
   let _source = source.split(' ')
   let result = _destination.map(element => _source.indexOf(element))
   let incementalCount = findIncremental(result)
   return (incementalCount+1) / result.length
}

isFilterResult = (source,destination)=> {
   let exist = 0
   let _destination = destination.split(' ')
   let _source = source.split(' ')
   let result = _destination.map(element => _source.indexOf(element))
   if(result.includes(-1)) return false
   else{
    let len = _destination.length
     for(let i=0,j=1;i<result.length-1,j<result.length;i++,j++){
        exist = Math.abs(result[j]-result[i]) + exist
    }
    if(exist === result.length-1) return true
    else return false
   }
}
let a=0
for(let i=0;i<s.length;i=i+pattern.length){
   chunk = s.substr(i,pattern.length)
   if(chunk !== pattern){
     temp = temp.concat(chunk)
     if(temp.length != pattern.length){
       let tempPattern = temp.split(' ').slice(a)
       moduloResultArr= []
       for(j=1;j<tempPattern.length;j++){
         let newIteratifTemp = tempPattern.slice(j).join(' ')
         if(newIteratifTemp.length < pattern.length)break
         if(isFilterResult(newIteratifTemp,pattern)){
            console.log("find it with %100") 
            we_find_it =true
            break
         }
         else{
           moduloResultArr.push(isFilter(newIteratifTemp,pattern))
         }
       }
       a++
       if(we_find_it) break
     }
   }
   else {
      console.log("we find it with %100",pattern) 
      we_find_it =true
      break
   }
}
if(!we_find_it){
  console.log("we find it with %",100*Math.max(...moduloResultArr))
}
